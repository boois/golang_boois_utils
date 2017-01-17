//服务端解包过程
package server

import (
	"net"
	"boois_ipc/boois_protocol"
)

func Send(c net.Conn, bs []byte) (int, error) {
	return c.Write(boois_protocol.Packet(bs));
}
func OnMsg(net_type string, addr string, each_fn func(net.Conn, []byte, error)) error {
	netListen, err := net.Listen(net_type, addr)
	//defer netListen.Close()
	if err != nil {
		return err
	}

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		go func() {
			//声明一个临时缓冲区，用来存储被截断的数据
			tmpBuffer := make([]byte, 0)

			buffer := make([]byte, 1024)

			//进行循环读取
			for {
				n, err := conn.Read(buffer)//返回读取到的长度
				if err != nil && err.Error() !="EOF"  {
					each_fn(conn, []byte{}, err)
				}
				//把未解包的片段加上本次读取到的数据
				//:n 只取前面有读取到的数据  后面都是空的
				tmpBuffer = boois_protocol.Unpack(append(tmpBuffer, buffer[:n]...), func(data []byte) {
					each_fn(conn, data, nil)
				})

			}
		}()

	}
}







