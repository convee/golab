package main

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var ErrPoolClosed = errors.New("资源池已经关闭")

//一个安全的资源池，被管理的资源必须实现io.Close接口
type Pool struct {
	m       sync.Mutex                //互斥锁
	res     chan io.Closer            //有缓冲的通道，通道类型是io.Closer接口
	factory func() (io.Closer, error) //函数类型
	closed  bool                      //资源池是否被关闭
}

//创建一个资源池
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size的值太小")
	}
	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil
}

//从资源池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquire:共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:新生成资源")
		return p.factory()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	//关闭通道，不让写入
	close(p.res)
	//关闭通道里的资源
	for r := range p.res {
		r.Close()
	}
}

func (p *Pool) Release(r io.Closer) {
	//保证该操作和close方法的操作是安全的
	p.m.Lock()
	defer p.m.Unlock()
	//资源池关闭了，就剩这一个没有释放的资源，释放即可
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.res <- r:
		log.Println("资源释放到池子里了")
	default:
		log.Println("资源池满了，释放这个资源吧")
		r.Close()

	}
}

const (
	//模拟的最大goroutine
	maxGoroutine = 5
	//资源池的大小
	poolRes = 2
)

func main() {
	//等待任务完成
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)

	p, err := New(createConnection, poolRes)
	if err != nil {
		log.Println(err)
		return
	}
	//模拟好几个goroutine同时使用资源池查询数据
	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("开始关闭资源池")
	p.Close()
}

//模拟数据库查询
func dbQuery(query int, pool *Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer pool.Release(conn)

	//模拟查询
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("第%d个查询，使用的是ID为%d的数据库连接", query, conn.(*dbConnection).ID)
}

//数据库连接
type dbConnection struct {
	ID int32 //连接的标志
}

//实现io.Closer接口
func (db *dbConnection) Close() error {
	log.Println("关闭连接", db.ID)
	return nil
}

var idCounter int32

//生成数据库连接的方法，以供资源池使用
func createConnection() (io.Closer, error) {
	//并发安全，给数据库连接生成唯一标志
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}, nil
}
