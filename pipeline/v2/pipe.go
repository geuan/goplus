package v2

import (
	"fmt"
	"sync"
	"time"
)

//使用channel改进管道模式

// 凡是支持管道模式的函数，其参数必须是channel，返回channel
type Slice []int

type Cmd func(list Slice) chan int // 值得借鉴，将相同的函数类型抽象为一个类型

type PipeCmd func(in chan int) chan int // 支持管道的函数

// 求偶数
func Events(list Slice) chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for _, value := range list {
			if value%2 == 0 { // 业务流程
				//time.Sleep(time.Second * 1)
				c <- value
			}
		}
	}()
	return c
}

// 乘以 10
func M10(in chan int) chan int { // 这个函数是支持管道的
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range in {
			time.Sleep(time.Second*1)
			out <- value * 10
		}
	}()
	return out
}

//func Pipe(args Slice, c1 Cmd, cs ...PipeCmd) chan int {
//	ret := c1(args)
//	if len(cs) == 0 {
//		return ret
//	}
//	retList := make([]chan int, 0)
//	for index, c := range cs {
//		if index == 0 {
//			retList = append(retList, c(ret))
//		} else {
//			getChan := retList[len(retList)-1]
//			retList = append(retList, c(getChan))
//		}
//	}
//	return retList[len(retList)-1]
//}

// 多路复用
func Pipe2(args Slice, c1 Cmd, cs ...PipeCmd) chan int {
	ret := c1(args) // 找偶数
	out := make(chan  int)
	wg := sync.WaitGroup{}
	for _, c := range cs {
		getChan := c(ret)
		wg.Add(1)
		go func(input chan int) {
			defer wg.Done()
			for v := range input {
				out <- v
			}
		}(getChan)
	}
	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out
}

func Test(nums []int) {
	ret := Pipe2(nums, Events, M10,M10,M10,M10)
	for r := range ret {
		fmt.Printf("%d ", r)
	}
}
