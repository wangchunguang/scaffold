package etcd

import (
	"fmt"
	"testing"
)

func Test_Demo(t *testing.T) {

	c := make(chan int)
	defer close(c)
	go func() {
		c <- 3 + 4
	}()

	for {
		select {
		case _, ok := <-c:
			fmt.Println(2222222, ok)
			return
		default:
			fmt.Println(1111111111)

		}
	}

}
