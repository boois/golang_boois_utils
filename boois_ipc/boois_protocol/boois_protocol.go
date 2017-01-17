//通讯协议处理，主要处理封包和解包的过程
package boois_protocol

import (
	"bytes"
	"encoding/binary"
)

const (
	ConstHeader         = "www.01happy.com" //包头
	ConstHeaderLength   = 15 //包头的长度  上面的那个字符串长度就是15
	ConstSaveDataLength = 4
)

//封包
func Packet(message []byte) []byte {
	//写入包头
	//写入message的长度
	//写入消息实体

	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

//解包
//完整的消息实体会被推入readerChannel
//返回不完整的信息  供下一次处理
//func Unpack(buffer []byte,each_fn func([]byte), readerChannel chan []byte) []byte {
func Unpack(buffer []byte,each_fn func([]byte)) []byte {
	length := len(buffer)

	var i int
	for i = 0; i < length; i = i + 1 {
		//长度小于包头的长度和存储消息长度的4个字节  说明包头不完整 或包头+数据长度不完整
		if length < i+ConstHeaderLength+ConstSaveDataLength {
			break
		}
		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {//这是包头
			//取出消息长度
			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstSaveDataLength])
			//如果buf的升度小于包头长度+4个字节+消息长度   说明消息实体未发完整
			if length < i+ConstHeaderLength+ConstSaveDataLength+messageLength {
				break
			}
			//取出消息实体
			data := buffer[i+ConstHeaderLength+ConstSaveDataLength : i+ConstHeaderLength+ConstSaveDataLength+messageLength]
			//推到channel中
			each_fn(data)
			//i往后偏移
			i += ConstHeaderLength + ConstSaveDataLength + messageLength - 1
		}
	}
	//位置移到了最后一样  说明这个buf读完了
	if i == length {
		//剩余没被读取的是空的
		return make([]byte, 0)
	}
	//返回没被读取的
	return buffer[i:]
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
