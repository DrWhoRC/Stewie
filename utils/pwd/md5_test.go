package utils

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	pwd1 := MakePassword("Wpywatsendw0517", "wangpuyan")
	fmt.Println(pwd1)
	pwd2 := MakePassword("Wpywatsendw0517", "wangpuyang")
	fmt.Println(pwd2)
}
