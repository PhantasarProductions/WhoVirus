package main

import(
	"fmt"
	"math/rand"
	"time"
	"trickyunits/qstr"
	"strings"
)

var CSDBG = false // When true debugging information is shown during creation of the session.

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

func got(r string) bool{
	for _,s := range user.ses.files{
		//CSChat("= checking: "+n+" contains \""+s+"\" and may not be \""+r+"\"")
		if s==r { return true }
	}
	return false
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
	{	func(virus,name string) string{
			st:=qstr.Left(virus,1)
			r:= "The virus doesn't start with the letter: "+st
			if got(r) { return "" }
			return r
		},
	},
	{	func(virus,name string) string{
			l:=len(virus)
			for l==len(virus){
				l=rand.Intn(7)+3
			}
			r:= fmt.Sprintf("The virus has %d letters",l)
			if got(r) { r="" }
			return r
		},
	},
	{	func(virus,name string) string{
			l:=len(virus)
			r:= fmt.Sprintf("The virus does not have %d letters",l)
			return r
		},
	},
	{	func(virus,name string) string{
			st:=strings.ToUpper(qstr.Right(virus,1))
			for st==strings.ToUpper(qstr.Right(virus,1)){
				st=strings.ToUpper(chr(byte(65+rand.Intn(26))))
			}
			return "The virus ends with the letter: "+st
		},
	},
	{	func(virus,name string) string{
			st:=qstr.Right(virus,1)
			r:= "The virus doesn't end with the letter: "+strings.ToUpper(st)
			if got(r) { return "" }
			return r
		},
	},
	{	func(virus,name string) string{
			n:=rand.Intn(len(virus))+1
			l:=qstr.Mid(virus,n,1)
			return "The virus does not contain the letter: "+strings.ToUpper(l)
		},
	},
	{	func(virus,name string) string{
			n:=RandomName()
			if !NHave(n) { return "" }
			if user.ses.files[n]=="*VIRUS*" { return "" }
			return n+" is the virus!"
		},
	},
	{
		func(virus,name string) string{
			if Sex[virus]=="F" {
				return "The virus has a boy's name"
			} else {
				return "The virus has a girl's name"
			}
		},
	},
	{
		func(virus,name string) string{
			if Sex[virus]=="F" {
				return "The virus does not have a girl's name"
			} else {
				return "The virus does not have a boy's name"
			}
		},
	},
	{
		func(virus, name string) string{
			return "Everything stated about the virus is true"
		},
	},
	{
		func(virus, name string) string{
			ret:=""
			if strings.ToUpper(virus)!=strings.ToUpper(user.name) { ret="The virus has your name" }
			return ret
		},
	},
	{
		func(virus, name string) string{
			a:=rand.Intn(10)+5
			if len(virus)>a { return fmt.Sprintf("The virus has less than %d letters",a) } else {return ""}
		},
	},
	{
		func(virus, name string) string{
			a:=rand.Intn(10)+5
			if len(virus)<a { return fmt.Sprintf("The virus has more than %d letters",a) } else {return ""}
		},
	},
	{
		func(virus, name string) string{
			l:=qstr.Left(virus,1)
			if l=="Y" {return ""}
			if l=="A" || l=="E" || l=="I" || l=="O" || l=="U" { return "The virus starts with a consonant" }
			return "The virus starts with a vowel"
		},
	},
	{
		func(virus, name string) string{
			l:=strings.ToUpper(qstr.Right(virus,1))
			if l=="Y" {return ""}
			if l=="A" || l=="E" || l=="I" || l=="O" || l=="U" { return "The virus ends with a consonant" }
			return "The virus ends with a vowel"
		},
	},
	{
		func(virus,name string) string{
			if len(virus)==len(name) {
				return "I don't have as many letters in my name as the virus does"
			} else {
				return "I have as many letters in my name as the virus does"
			}
		},
	},
	{
		func(virus,name string) string{
			e:=strings.ToUpper(qstr.Left(virus,1))
			l:=strings.ToUpper(qstr.Right(virus,1))
			if e==l {
				return "The first letter of the virus does not match its last letter"
			} else {
				return "The first letter of the virus matches the last letter"
			}
		},
	},
	{
		func(virus,name string) string{
			e:=fmt.Sprintf("%d",len(virus))
			r:=""
			switch qstr.Right(e,1){
				case "0","2","4","6","8": r="odd"
				case "1","3","5","7","9": r="even"
			}
			return "The number of letters in the virus is "+r
		},
	},
	{
		func(virus,name string) string{
			if virus!="Kenny"{
				return "Oh my god! They killed the virus! You bastard!"
			} else {
				return ""
			}
		},
	},
	{
		func(virus, name string) string{
			l:=strings.ToUpper(qstr.Mid(virus,2,1))
			if l=="Y" {return ""}
			if l=="A" || l=="E" || l=="I" || l=="O" || l=="U" { return "A consonant is the second letter of the virus" }
			return "A vowel is the second letter of the virus"
		},
	},
	
}

func CreateSession(){
	doing("Creating","new session")
	// Init memory
	user.ses = session{}
	s:=&user.ses
	s.files = map[string] string{}
	s.revealed = map[string] bool{}
	// Define the virus
	virus:=RandomName()
	s.files[virus]="*VIRUS*" 
	CSChat(virus+" is the virus")
	// Define the 99 other files
	for i:=0;i<99;i++{
		rn:=""
		for rn=="" || (NHave(rn)){  // Very ugly approach, but since Go has no support for do{}while or repeat+until, this was all I can do.... :-/
			rn=RandomName()
		}
		CSChat(fmt.Sprintf("For record %d name %s was chosen",i,rn))
		ft:="" 
		for ft=="" || got(ft) {ft=Hints[rand.Intn(len(Hints))].h(virus,rn)}
		CSChat("= Text: "+ft)
		s.files[rn]=ft
		s.revealed[rn]=false
	}
}


func init(){
	seed:=time.Now().UTC().UnixNano()
	rand.Seed( seed )
	CSChat(fmt.Sprintf("Seed set to %d",seed))
}
