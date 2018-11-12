package TEST

import (
	"fmt"
	"testing"
)

/**
* @program: src
*
* @description:进行单元测试
*
* @author: sunlei
*
* @create: 2018-11-12 15:33
**/

func Testadd(t *testing.T) {
	r := add(1, 2)
	if r != 3 {
		fmt.Println("测试错误,返回值为", r)
		t.Fatal("测试错误,返回值为", r)
	}
	fmt.Println("ok!")
	t.Log("ok!")
}
