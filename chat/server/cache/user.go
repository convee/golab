package cache

import (
	"encoding/json"
	"fmt"
	"golab/chat/common"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	UserTable = "users"
)

type UserManager struct {
	pool *redis.Pool
}

func NewUserManager() (mgr *UserManager) {
	mgr = &UserManager{
		pool: pool,
	}
	return
}

func (p *UserManager) getUser(id int) (user *common.User, err error) {
	res, err := redis.String(p.pool.Get().Do("HGet", UserTable, fmt.Sprintf("%d", id)))
	if err != nil {
		if err == redis.ErrNil {
			err = ErrUserNotExists
		}
		return
	}
	user = &common.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		return
	}
	return
}

func (p *UserManager) Register(user *common.User) (err error) {
	conn := p.pool.Get()
	defer conn.Close()
	if user == nil {
		fmt.Println("invalid user")
		err = ErrInvalidParams
		return
	}
	_, err = p.getUser(user.UserId)
	if err != nil {
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("HSet", UserTable, fmt.Sprintf("%d", user.UserId), string(data))
	if err != nil {
		return
	}
	return

}

func (p *UserManager) Login(id int, passwd string) (user *common.User, err error) {
	conn := p.pool.Get()
	defer conn.Close()
	user, err = p.getUser(id)
	if err != nil {
		return
	}
	if user.Passwd != passwd {
		err = ErrInvalidPasswd
		return
	}
	user.LastLogin = fmt.Sprintf("%v", time.Now())
	return

}
