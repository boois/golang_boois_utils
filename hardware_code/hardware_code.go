package hardware_code

import (
	"strings"
	"fmt"
	"runtime"
	"net"
	"os/exec"
	"crypto/hmac"
	"crypto/md5"
	"io/ioutil"
)
var boois_soft_sec="福州市鼓楼区步易天下网络科技有限公司"

func get_mac() (string, error) {
	interfaces, err := net.Interfaces()
	macs_arr := []string{}
	if err != nil {
		return "", err
	}
	for _, inter := range interfaces {
		//fmt.Println("each",inter.Name,inter.HardwareAddr)
		if len(inter.HardwareAddr) == 0 {
			continue
		}
		if strings.HasPrefix(inter.Name, "eth") || strings.HasPrefix(inter.Name, "en") || strings.Contains(inter.Name,"网") {
			macs_arr = append(macs_arr, fmt.Sprintf("%s", inter.HardwareAddr))
		}
	}
	return strings.Join(macs_arr, ","), nil
}
func get_hard_disk_linux()(string,error) {
	out, err := exec.Command("ls /dev/disk/by-uuid/").Output()
	if err != nil {
		return "",err
	}
	return string(out),err
}
func hmac_md5(string_to_sign, secret_access_key string) string {
	h := hmac.New(md5.New, []byte(secret_access_key))
	h.Write([]byte(string_to_sign))
	return fmt.Sprintf("%x",h.Sum(nil))
}
func dxdiag_windows() (string,error) {
	_, err := exec.Command("cmd","dxdiag /t dxdiag.txt").Output()
	//fmt.Println("out",string(out))
	if err != nil {
		//fmt.Println(err)
		return "",err
	}

	bs,err:=ioutil.ReadFile("dxdiag.txt")
	//fmt.Println("bs",string(bs))
	if err != nil {
		return "",err
	}

	for _,line:=range strings.Split(string(bs),"\n"){
		//fmt.Println("line>",line)
		if strings.Contains(line,"Machine Id"){
			fmt.Println("MID",line)
			arr:=strings.Split(line,":")
			if len(arr)!=2{
				return "",nil
			}
			mid:=strings.Replace(arr[1],"{","",-1)
			mid=strings.Replace(mid,"}","",-1)
			return mid,nil
		}
	}
	return "",nil
}
func GetCode() string {

	txt_arr:=[]string{}
	//网卡 mac
	mac, _ := get_mac()
	txt_arr=append(txt_arr,"mac:"+mac)
	//fmt.Println("mac",mac, err)
	//硬盘
	hard_disk:=""
	//fmt.Println("当前系统",runtime.GOOS)
	switch runtime.GOOS {
	case "darwin":
	case "linux":
		hard_disk,_=get_hard_disk_linux()
		txt_arr=append(txt_arr,"hard_disk:"+hard_disk)

	case "windows":
		//速度太慢 不使用
		//mid,_:=dxdiag_windows()
		//txt_arr=append(txt_arr,"mid:"+mid)
	}
	_=hard_disk

	return hmac_md5(strings.Join(txt_arr,"\n"),boois_soft_sec)
}
