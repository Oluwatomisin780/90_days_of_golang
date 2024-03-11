package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"sync"
)


func checkAndSaveBody(url string, wg *sync.WaitGroup,c chan string){
resp,err:=http.Get(url)
 if err !=nil{
	//fmt.Println(err)
	 s:= fmt.Sprintf("%s is Down\n",url)
	 s+= fmt .Sprintf("Error: %\n ",err)
	 c <- s // sending to channel
 }else{
	 defer resp.Body.Close()
	s:= fmt.Sprintf("%s -> status code: %d \n ",url,resp.StatusCode)
	if resp.StatusCode==200{
		bodyBytes,err:= ioutil.ReadAll(resp.Body)
		file:= strings.Split(url, "//")[1]
		file += ".txt"
		s+=fmt.Sprintf("Writing response body to %s\n",file)
		

		err =ioutil.WriteFile(file,bodyBytes,0664)
		if err !=nil{
			//log.Fatal(err)
			s+="Error Writing into file "
			c  <-s 
		}
		s+= fmt.Sprintf("%s is up\n ",url)
	}
 }
 wg.Done()
}



func main(){
	// Building a concurrency projrct
	var wg sync.WaitGroup

	urls:= []string{"https://golang.org","https://www.google.com", "https://www.medium,com"}
	wg.Add(len(urls))
//declare a new channel
c:= make(chan string)
	for _,url := range urls{
		go  checkAndSaveBody(url,&wg,c)
		fmt.Println(strings.Repeat("#",20))
	}


	fmt.Println("Number of Gorutines",runtime.NumGoroutine())

	for i:=0 ; i<len(urls); i++{
		fmt.Println(<-c)
	}
	wg.Wait()
}