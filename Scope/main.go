package main

import f "fmt"

const done = false // pakage Scope

func main(){
	var  task = " running" // Block  Scope
	f.Println(done,task )
	const  done =  true
	f .Printf("done in the main %v\n",done)
}

// ---------------------------------------------------------------------------------------------------------------------
// scopes
//-----------------------------
//  Type scope
//-----------------
// 1.  block scope
//2. package  Scope
// 3. file Scope  
//-------------------------------------------------------------------------------