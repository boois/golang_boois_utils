package test

import (
	"os"
	"net"
	"log"
	"fmt"
	"io/ioutil"
	"time"
	"testing"
	"boois_ipc"
)

func TestBooisUtilsServer(t *testing.T) {

	go func() {
		os.Remove("./test.sock");
		err := boois_ipc.ServerOnMsg(NET_TYPE,ADDR, func(conn net.Conn, bs []byte, err error) {
			if err != nil {
				log.Println("发生错误:", err)
				return
			}
			fmt.Println("收到消息", len(bs))
			ioutil.WriteFile("2016-12-27-18-05-56-2.json",bs,0644)
			//处理业务.....
			//返回结果
			boois_ipc.ServerSend(conn, []byte("处理完成了"));
		})
		fmt.Println(err)
	}()

	//服务端主进程
	fmt.Println("服务端主进程开始进行...")
	time.Sleep(50*time.Second)
	fmt.Println("服务端主进程结束")


}
