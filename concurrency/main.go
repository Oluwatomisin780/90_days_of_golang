//concurrency means loading more goroutine at a time.
//parallelism means muiltple gorountine at the same time

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func doSomething(){
	fmt.Println("Do something here")
}
func f1 (wg *sync.WaitGroup)  {
	fmt.Println("f1(goroutine started)" )
	for i:= 0;i<3 ;i++{
		fmt.Println("F1 I:",i)
		time.Sleep(time.Second)
	}
	fmt.Println("f1  execution finished") 
	wg.Done()
}
func f2 ()  {
	fmt.Println("f2(goroutine started)" )
	for i:= 0;i<3 ;i++{
		fmt.Println("F2 I:",i)
	}
	fmt.Println("f2  execution finished") 
	
}
func main(){

//normal function calling
	//doSomething()

//spawning a gouroutines ----- a go routine is a lightweight  thread  of excution----
	//go doSomething()


	//time.Sleep(time.Second*2)


	// fmt.Println("main execution started!!")
	// fmt.Println("no of cpu",runtime.NumCPU())
	// fmt.Println("no of Gorutine",runtime.NumGoroutine())
	// fmt.Println("os", runtime.GOOS)
	// fmt.Println("Arch",runtime.GOARCH)
	// fmt.Println("GOMAxProcs",runtime.GOMAXPROCS(0))
 

	var wg sync.WaitGroup

	wg.Add(1)
	go f1(&wg)
	fmt.Println("no of Gorutine",runtime.NumGoroutine())

	f2()

//time.Sleep(time.Second*3)
 wg.Wait()
	fmt.Println("main execution stopped")
}