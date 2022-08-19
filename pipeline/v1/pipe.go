package v1

import (
	"fmt"
	"time"
)

type Slice []int

type Cmd func(list Slice) (ret Slice)  // 值得借鉴，将相同的函数类型抽象为一个类型

func Events(list Slice) (ret Slice) {
	for _,value := range list {
		if value % 2 == 0 {
			time.Sleep(time.Second*1)
			ret = append(ret,value)
		}
	}
	return ret
}

func Multipy(list Slice) (ret Slice)  {
	for _,value := range list {
		time.Sleep(time.Millisecond*300)
		ret = append(ret,value*10)
	}
	return ret
}

func P(args Slice,c1 Cmd,c2 Cmd) []int {
	ret := c1(args)

	//fmt.Println(c2(ret))
	return  c2(ret)

}

func Test(nums []int)  {
	 ret := P(nums,Events,Multipy)
	for r:=range ret{
		fmt.Printf("%d ",r)
	}
}

