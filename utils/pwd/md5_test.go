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

//265a0bd24dd5f106bd9b4efd34a9a18e
