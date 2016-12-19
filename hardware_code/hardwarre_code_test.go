package hardware_code

import (
	"testing"
	"fmt"
)

func TestGetCode(t *testing.T) {
	code:=GetCode()
	fmt.Println(code)
}