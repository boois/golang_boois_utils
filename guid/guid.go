package guid

import (
	"os"
	"fmt"
)

func  Guid() string{
	f, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x", b)
	return uuid
}
