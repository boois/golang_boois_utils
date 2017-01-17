package test

import (
	"testing"
	"io/ioutil"
	"fmt"
	"log"
	"os"
	"boois_ipc/client"
)

func BenchmarkName(b *testing.B) {
	cli, err := client.Connect(NET_TYPE, ADDR);
	if err != nil {
		log.Println(":", err)
	}
	for i := 0; i < b.N; i++ {

		bs, err := ioutil.ReadFile("2016-12-27-18-05-56.json")
		fmt.Println("发送的单条消息长度", len(bs))
		if err != nil {
			log.Println("读取文件失败")
			os.Exit(1)
		}
		res, err := client.Send(cli, bs)
		fmt.Println("结果:", string(res))
	}
}
func TestCli(t *testing.T) {
	cli, err := client.Connect(NET_TYPE, ADDR);
	if err != nil {
		log.Println(":", err)
	}
	bs, err := ioutil.ReadFile("2016-12-27-18-05-56.json")
	fmt.Println("发送的单条消息长度", len(bs))
	if err != nil {
		log.Println("读取文件失败")
		os.Exit(1)
	}
	res, err := client.Send(cli, bs)
	fmt.Println("结果:", string(res))

}
