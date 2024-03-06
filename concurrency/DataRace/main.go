package main

import (
	"fmt"
	"sync"
	"time"
)


func main(){
	const gr = 100

	var wg  sync.WaitGroup
	wg.Add(gr*2)
	var n int =0 

	for i:=0;i<gr;i++{
		go func (){
			time.Sleep(time.Second/10)
			n++
			wg.Done()
		}()// function invocation

		go func (){
			time.Sleep(time.Second/10)
			n--
			wg.Done()
		}()// function invocation
	}
	wg.Wait()
	fmt.Println(n)
}