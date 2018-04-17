package main

import(
	"fmt"
	"math/rand"
	"time"
	"trickyunits/qstr"
)

var CSDBG = true // When true debugging information is shown during creation of the session.

func CSChat(k string){
	if CSDBG { fmt.Println("DEBUG> ",k) }
}
func chr(bte byte)string{
	a:=[]byte{bte}
	b:=string(a)
	return b
}


func RandomName() string{
	rni:=rand.Intn(len(Names))
	return Names[rni]
}

func NHave(name string) bool{
	_,ok:=user.ses.files[name]
	return ok
}

type tHint struct { h func(virus,name string) string }
var Hints = []tHint{
	{	func(virus,name string) string{
			st:=qstr.Left(virus,1)
			for st==qstr.Left(virus,1){
				st=chr(byte(65+rand.Intn(26)))
			}
			return "The virus starts with the letter: "+st
		},
	},
}

func CreateSession(){
	// Init memory
	user.ses = session{}
	s:=&user.ses
	s.files = map[string] string{}
	// Define the virus
	virus:=RandomName()
	s.files[virus]="*VIRUS*" 
	CSChat(virus+" is the virus")
	for i:=0;i<99;i++{
		rn:=""
		for rn=="" || (NHave(rn)){  // Very ugly approach, but since Go has no support for do{}while or repeat+until, this was all I can do.... :-/
			rn=RandomName()
		}
		CSChat(fmt.Sprintf("For record %d name %s was chosen",i,rn))
		ft:="" 
		for ft=="" {ft=Hints[rand.Intn(len(Hints))].h(virus,rn)}
		CSChat("= Text: "+ft)
	}
}


func init(){
	seed:=time.Now().UTC().UnixNano()
	rand.Seed( seed )
	CSChat(fmt.Sprintf("Seed set to %d",seed))
}
