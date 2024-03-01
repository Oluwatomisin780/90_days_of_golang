package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main(){
	var newFile*os.File
// A pointer is variable that stores the memory address of other variable
	fmt.Printf("%T\n", newFile)

	var error error 
	newFile,error = os.Create("a.txt")

	if error !=nil{
		// fmt.Println(error)
		// os.Exit(1)
		log.Fatal(error)
	}
	error= os.Truncate("a.text" ,0)

newFile.Close()

file,err:= os.OpenFile("a.txt",os.O_CREATE| os.O_APPEND,0644)
if err!=nil{
		// fmt.Println(error)
		// os.Exit(1)
		log.Fatal(error)
	}
	file.Close()
	//getting file info
	var fileInfo os.FileInfo 
	fileInfo,err= os.Stat("a.txt")

	p:= fmt.Println

	p("FileName",fileInfo.Name())
	p("size in bytes",fileInfo.Size())
	p("last modified",fileInfo.ModTime())
	p("Is Directory",fileInfo.IsDir())
	p("permission",fileInfo.Mode())

	fileInfo,err= os.Stat("b.txt")

	if err != nil{
		if os.IsNotExist(err){
			fmt.Println("file does not exist")
		}
	}


	oldPath:="a.txt"
	newPath:= "b.txt"

	err=os.Rename(oldPath,newPath)
	if err!=nil{
		 fmt.Println(error)
		// os.Exit(1)
		// log.Fatal(error)
	}

	//writing bytes to file oswrite 
	file2,err:= os.OpenFile(
		"d.txt",os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err!=nil{ 
		log.Fatal(err)
	}
	defer file2 .Close()

	byteSlice:= []byte ("i learnt golang!")
	byteWritten,err:=  file2.Write(byteSlice)
	if err!=nil{ 
		log.Fatal(err)
	}
	log.Printf("byteWritten: %d\n",byteWritten)
//
bs := []byte ("Go programming is cool")

err = ioutil.WriteFile("c.txt",bs,0644)
if err!=nil{ 
		log.Fatal(err)
	}

cd := [] byte ("golang da gan")
 err =os.WriteFile("f.txt",cd,0644)
 if err!=nil{ 
		log.Fatal(err)
	}

// write to file using buffered writer 
file3 ,err := os.OpenFile("myFile.txt",os.O_WRONLY|os.O_CREATE,0644)

if err!= nil{
	log.Fatal(err)
}
defer file3.Close()
// creating a buffer writer 
bufferWriter := bufio.NewWriter(file3)
bs1 := []byte {97,98,99}
byteWrittens,err:= bufferWriter.Write(bs1)
if err!= nil{
	log.Fatal(err)
}
log.Printf("byte wriiten to buffer (not file ) %d\n",byteWrittens)
byteAvailable := bufferWriter.Available()
log.Printf("bytes is available in buffer %d\n",byteAvailable) 

byteWrittens,err =   bufferWriter.WriteString("\n just a random string")
if err != nil{
	log.Fatal(err)
}

unflushedBuffered := bufferWriter.Buffered()
log.Printf("bytes buffered: %d\n",unflushedBuffered)
bufferWriter.Flush()

//ReadingFile using bufferedbytes

file4,err:= os.Open("myFile.txt")

if err !=nil{
	log.Fatal(err)
}
defer file4.Close()
byteSlice2 := make([]byte,8)

numberByteRead,err := io.ReadFull(file4,byteSlice2)
if err != nil {
log.Fatal(err)
}
log.Printf("number of byte read: %d\n",numberByteRead)
log.Printf("Data read: %s\n",byteSlice2)


file5,err :=  os.Open("main.go")
 if err != nil {
	log.Fatal(err)
 }
 defer file5.Close()
 data,err := ioutil.ReadAll(file5)

 if err != nil{
	 log.Fatal(err)
 }

 fmt.Printf("Data is a string: %s\n",data)

 // reading file using scanner line 

 file6,err := os.Open("c.txt")
 if  err != nil{
	log.Fatal(err) 
 }
 defer file6.Close()
 scanner := bufio.NewScanner(file6) 

 success:= scanner.Scan()

 if success ==false{
	scanner.Err()
	if  err == nil{
		log.Println("Scan was completed and it reach Eof")
	}else{
		log.Fatal(err)
	}

 }
 fmt.Println("first line found:",scanner.Text())  
 for scanner.Scan(){
	fmt.Println(scanner.Text())
 }
 if err :=scanner.Err();err !=nil{
	log.Fatal()
 }

}