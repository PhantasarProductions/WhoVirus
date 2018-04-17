package main

import(
	"fmt"
	"math/rand"
	"time"
)

var CSDBG = true // When true debugging information is shown during creation of the session.

func CSChat(k string){
	if CSDBG { fmt.Println("DEBUG> ",k) }
}

func CreateSession(){
	// Init memory
	user.ses = session{}
	s:=&user.ses
	s.files = map[string] string{}
	// Define the virus
	
}


func init(){
	seed:=time.Now().UTC().UnixNano()
	rand.Seed( seed )
	CSChat(fmt.Sprintf("Seed set to %d",seed))
}
