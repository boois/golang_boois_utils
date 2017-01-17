//客户端发送封包
package client

import (
	"net"
	"boois_ipc/boois_protocol"
)

func Send(conn net.Conn,bs []byte) ([]byte,error) {
	_,err:=conn.Write(boois_protocol.Packet(bs))
	if err!=nil{
		return []byte{},err
	}
	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0)

	buffer := make([]byte, 1024)
	res:=[]byte{}
	//进行循环读取
	for {
		//有收到一个结果就停止
		if len(res)>0{
			break;
		}
		n, err := conn.Read(buffer)//返回读取到的长度
		if err != nil {
			return res, err
		}
		//把未解包的片段加上本次读取到的数据
		//:n 只取前面有读取到的数据  后面都是空的
		tmpBuffer = boois_protocol.Unpack(append(tmpBuffer, buffer[:n]...), func(data []byte) {
			res=data
		})
	}
	return res,nil
}

func Connect(net_type string,addr string)(net.Conn,error) {
	conn, err :=net.Dial(net_type,addr)
	if err != nil {
		return conn,err
	}
	//defer conn.Close()
	return conn,err
}
