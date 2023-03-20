package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var myChan chan int
	myChan = make(chan int, 10)
	wite := sync.WaitGroup{}
	wite.Add(2)
	go func(myChan chan int, wite *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			fmt.Println("写入：", i)
			myChan <- i
		}
		close(myChan)
		(*wite).Done()
	}(myChan, &wite)

	go func(myChan chan int, wite *sync.WaitGroup) {
		for {
			data, ok := <-myChan
			fmt.Println("读取：", data)
			if !ok {
				fmt.Println("返回")
				break
			}
			time.Sleep(time.Second)
		}
		wite.Done()
	}(myChan, &wite)
	wite.Wait()

}
