package main

import (
	"fmt"
	"go++/pipeline/getdata"
	"time"
)

/*
//认识管道模式，最初级的代码

// 从一个切片中找到偶数，并乘以10
type Slice []int

type Cmd func(list Slice) (ret Slice)  // 值得借鉴，将相同的函数类型抽象为一个类型

func Events(list Slice) (ret Slice) {
	for _,value := range list {
		if value % 2 == 0 {
			ret = append(ret,value)
		}
	}
	return ret
}

func Multipy(list Slice) (ret Slice)  {
	for _,value := range list {
		ret = append(ret,value*10)
	}
	return ret
}

func P(args Slice,c1 Cmd,c2 Cmd)  {
	ret := c1(args)

	fmt.Println(c2(ret))
}

func main()  {
	nums := []int{1,2,3,4,5,6,7,8}
	P(nums,Events,Multipy)
}
*/



/*
// 使用channel改进管道模式

// 凡是支持管道模式的函数，其参数必须是channel，返回channel
type Slice []int

type Cmd func(list Slice) chan int // 值得借鉴，将相同的函数类型抽象为一个类型

type PipeCmd func(in chan int) chan int  // 支持管道的函数

// 求偶数
func Events(list Slice) chan  int {
	c := make(chan  int,5)
	go func() {
		defer close(c)
		for _,value := range list {
			if value % 2 == 0 {   // 业务流程
				//fmt.Println(value)
				c <- value
			}
		}
	}()
	return c
}

// 乘以 10
func M10(in chan int) chan int {  // 这个函数是支持管道的
	out := make(chan  int)
	go func() {
		defer  close(out)
		for value := range in  {
			out <- value * 10
		}
	}()
	return out
}

func Pipe(args Slice,c1 Cmd,c2 PipeCmd)  chan int{
	ret := c1(args)

	return c2(ret)
}

func main()  {
	nums := []int{1,2,3,4,5,6,7,8}
	ret := Pipe(nums,Events,M10)
	for value := range ret {
		fmt.Println(value)
	}
}
*/



/*
// 管道模式之多路复用、提高性能
func test(v string)  {
	nums := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
	start:=time.Now().Unix()
	if v == "v1" {
		v1.Test(nums)

	} else {
		v2.Test(nums)
	}
	end:=time.Now().Unix()
	fmt.Printf("%s测试--用时:%d秒\r\n",v,end-start)
}

func main()  {
	//test("v1")
	test("v2")

}
*/

func testData()  {
	start := time.Now().Unix()
	getdata.Test()
	end := time.Now().Unix()
	fmt.Printf("测试用时：%d秒\r\n",end-start)

}
func main()  {
	testData()
}





















