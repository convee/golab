最近在看nsq的源码时候，发现它处理message的时候，都会采用字节序进行数据包的处理，于是我觉得有必要深入了解下TCP协议中 字节序的知识

字节序（Byte Order）
我们一般把字节（byte）看作是数据的最小单位。当然，其实一个字节中还包含8个bit (bit = binary digit)。 在一个32位的CPU中“字长”为32个bit，也就是4个byte。在这样的CPU中，总是以4字节对齐的方式来读取或写入内存， 那么同样这4个字节的数据是以什么顺序保存在内存中的呢？我们下面详细探讨一下。

字节序包括：大端序和小端序，为什么要这么麻烦还要分门别类呢？举个例子，255用二进制表达就是1111 1111，再加1就是1 0000 0000，多了一个1出来，显然我们需要再用额外的一个字节来存放这个1，但是这个1要存放在第一个字节还是第二个字节呢？这时候因为人们选择的不同，就出现了大端序和小端序的差异。

而所谓大字节序（big endian），便是指其“最高有效位（most significant byte）”落在低地址上的存储方式。例如像地址a写入0x0A0B0C0D之后，在内存中的数据便是：

big-end

而对于小字节序（little endian）来说就正好相反了，它把“最低有效位（least significant byte）”放在低地址上。例如：

little-end

对于我们常用的CPU架构，如Intel，AMD的CPU使用的都是小字节序，而例如Mac OS以前所使用的Power PC使用的便是大字节序（不过现在Mac OS也使用Intel的CPU了）。

Go 处理字节序
Go中处理大小端序的代码位于 encoding/binary ,包中的全局变量BigEndian用于操作大端序数据，LittleEndian用于操作小端序数据，这两个变量所对应的数据类型都实行了ByteOrder接口：

type ByteOrder interface {
    Uint16([]byte) uint16
    Uint32([]byte) uint32
    Uint64([]byte) uint64
    PutUint16([]byte, uint16)
    PutUint32([]byte, uint32)
    PutUint64([]byte, uint64)
    String() string
}
其中，前三个方法用于读取数据，后三个方法用于写入数据。

大家可能会注意到，上面的方法操作的都是无符号整型，如果我们要操作有符号整型的时候怎么办呢？很简单，强制转换就可以了，比如这样：

func PutInt32(b []byte, v int32) {
        binary.BigEndian.PutUint32(b, uint32(v))
}
为了深入了解它们，我们先写一个程序观察下go处理大端序和小端序的方式：

package main

import (
    "encoding/binary"
    "fmt"
    "unsafe"
)

const INT_SIZE int = int(unsafe.Sizeof(0))

//判断我们系统中的字节序类型
func systemEdian() {
    var i int = 0x1
    bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
    if bs[0] == 0 {
        fmt.Println("system edian is little endian")
    } else {
        fmt.Println("system edian is big endian")
    }
}

func testBigEndian() {

    // 0000 0000 0000 0000   0000 0001 1111 1111
    var testInt int32 = 256
    fmt.Printf("%d use big endian: \n", testInt)
    var testBytes []byte = make([]byte, 4)
    binary.BigEndian.PutUint32(testBytes, uint32(testInt))
    fmt.Println("int32 to bytes:", testBytes)

    convInt := binary.BigEndian.Uint32(testBytes)
    fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func testLittleEndian() {

    // 0000 0000 0000 0000   0000 0001 1111 1111
    var testInt int32 = 256
    fmt.Printf("%d use little endian: \n", testInt)
    var testBytes []byte = make([]byte, 4)
    binary.LittleEndian.PutUint32(testBytes, uint32(testInt))
    fmt.Println("int32 to bytes:", testBytes)

    convInt := binary.LittleEndian.Uint32(testBytes)
    fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func main() {
    systemEdian()
    fmt.Println("")
    testBigEndian()
    testLittleEndian()
}

执行的结果：

system edian is big endian

256 use big endian:
int32 to bytes: [0 0 1 0]
bytes to int32: 256

256 use little endian:
int32 to bytes: [0 1 0 0]
bytes to int32: 256
总结
为了程序的兼容，我们在开发跨服务器的TCP服务时，每次发送和接受数据都要进行转换，这样做的目的是保证代码在任何计算机上执行时都能达到预期的效果。