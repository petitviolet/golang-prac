package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Any interface{}
type Function func() Any

func asyncF(f Function) <-chan Any {
	ch := make(chan Any)
	go func() {
		ch <- f()
		close(ch)
	}()
	return ch
}

func main() {
	fmt.Println("vim-go")
	wg := new(sync.WaitGroup)
	f := func() Any {
		defer wg.Done()
		s := fmt.Sprintf("arguments: %+v", os.Args)
		time.Sleep(1000 * time.Millisecond)
		return s
	}
	g := func() Any {
		defer wg.Done()
		s := len(os.Args)
		time.Sleep(3000 * time.Millisecond)
		return s
	}

	fmt.Println(time.Now())
	var resF <-chan Any
	wg.Add(2)
	resF = asyncF(f)
	resG := asyncF(g)
	wg.Wait()
	fmt.Println(time.Now())
	fmt.Println(<-resF)
	fmt.Println(time.Now())
	fmt.Println(<-resG)
	fmt.Println(time.Now())
}
