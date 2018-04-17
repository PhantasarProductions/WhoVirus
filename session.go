package main

import(
	"fmt"
)

var CSDBG = true // When true debugging information is shown during creation of the session.

func CSChat(k string){
	if CSDBG { fmt.Println("DEBUG> ",k) }
}

func CreateSession(){
	user.ses = session{}
	s:=&user.ses
	s.files = map[string] string{}
}
