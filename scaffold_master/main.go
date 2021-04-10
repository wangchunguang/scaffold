package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type item struct {
	num int
}

func main() {
	i := &item{num: 5}
	data := i
	data.num = 10
	fmt.Println(i)

	//by := []byte{}
	//fmt.Println(cap(by))
	//user := &user{}
	//fmt.Println(cap(user.IDs))
	//user.IDs = append(user.IDs, "1")
	//fmt.Println(cap(user.IDs))
	//user.IDs = append(user.IDs, "2")
	//fmt.Println(cap(user.IDs))
	//user.IDs = append(user.IDs, "3")
	//// 1 2 3   4
	//fmt.Println(cap(user.IDs))
	//a := append(user.IDs, "4")
	//fmt.Println(a)
	//// 123 5
	//fmt.Println(cap(a))
	//b := append(user.IDs, "5")
	//fmt.Println(b)
	//// 123  6
	////c := append(user.IDs, "6")
	//// 123 7
	////d := append(user.IDs, "7")
	//// 123  8
	////e := append(user.IDs, "8")
	//
	////fmt.Println(a, b, c, d, e, )
	//user.IDs = append(user.IDs, "3")
	//user.IDs = append(user.IDs, "5")
	//user.IDs = append(user.IDs, "8")
	//fmt.Println(cap(user.IDs))
	//	 1024 1.25
}

func coordinateWithWaitGroup() {
	var wg sync.WaitGroup
	//两个goroutine
	wg.Add(2)
	num := int32(0)
	fmt.Printf("the number: %d [with sync.WaitGroup]\n", num)
	max := int32(10)

	go addNum(&num, 3, max, wg.Done)
	go addNum(&num, 4, max, wg.Done)
	wg.Wait()
}

// addNum 用于原子地址numP所指的变量的值
func addNum(numP *int32, id, max int32, deferFunc func()) {
	// 最后调用 wg.Done
	defer func() {
		deferFunc()
	}()

	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		if currNum >= max {
			break
		}
		newNum := currNum + 2
		time.Sleep(time.Microsecond * 20)
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			//fmt.Printf("the number :%d [%d-%d]\n", newNum, id, i)
			fmt.Println(1, id)
		} else {
			//fmt.Printf("the CAS operation failed [%d-%d]\n", id, i)
			fmt.Println(2, id)
		}
	}
}
