package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)


func checkAndSaveBody(url string, wg *sync.WaitGroup){
resp,err:=http.Get(url)
 if err !=nil{
	fmt.Println(err)
	 fmt.Printf("%s is Down\n",url)

 }else{
	 defer resp.Body.Close()
	fmt.Printf("%s -> status code: %d \n ",url,resp.StatusCode)
	if resp.StatusCode==200{
		bodyBytes,err:= ioutil.ReadAll(resp.Body)
		file:= strings.Split(url, "//")[1]
		file += ".txt"
		fmt.Printf("Writing response body to %s\n",file)

		err =ioutil.WriteFile(file,bodyBytes,0664)
		if err !=nil{
			log.Fatal(err)
		}
	}
 }
 wg.Done()
}



func main(){
	// Building a concurrency projrct
	var wg sync.WaitGroup

	urls:= []string{"https://golang.org","https://www.google.com", "https://www.medium,com"}
	wg.Add(len(urls))

	for _,url := range urls{
		go  checkAndSaveBody(url,&wg)
		fmt.Println(strings.Repeat("#",20))
	}

	fmt.Println("Number of Gorutines",runtime.NumGoroutine())
	wg.Wait()
}