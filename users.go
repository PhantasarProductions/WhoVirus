package main

import "trickyunits/dirry"
import "trickyunits/qff"
import "trickyunits/ansistring"
import jcr6 "trickyunits/jcr6/jcr6main"
import _ "trickyunits/jcr6/jcr6lzma"
import "os"
import "fmt"

type tuser struct{
	name string
	file string
	password string
	ses session
	insession bool
	totalscore int
	sessions int
	ansi bool
}

var user tuser
var userdir = dirry.Dirry("$AppSupport$/$LinuxDot$Phantasar Productions/Who Is The Virus/")

func SaveUser(){
	doing("Saving user: ",user.name)
	// global user data
	wo:="PW "+user.password+"\n"
	wo+="INSESSION "+b2s(user.insession)+"\n"
	wo+=fmt.Sprintf("TOTALSCORE %s\n",user.totalscore)
	wo+=fmt.Sprintf("SESSIONS %s\n",user.sessions)
	wo+="ANSI "+b2s(user.ansi)+"\n"
	// Session data
	// Writeout         
	e:=os.MkdirAll(userdir,0777)
	j:=jcr6.JCR_Create(user.file,"BRUTE")
	if e!=nil {
		panic (e)
	}
	j.AddString(wo,"User","BRUTE",0777,0,"Mr. Virus","I ruin you all")
	j.Close()
}

func NewUser(f,uname,pass string){
	user = tuser{}
	doing("Creating user:",uname)
	user.name=uname
	user.file=f
	user.password=pass
	user.insession=false
	user.totalscore=0
	user.sessions=0
	user.ansi=ansistring.ANSI_Use
}


func CreateUser(file,username,password string){
	if !yes("Create user '"+username+"'") { os.Exit(0) }
	NewUser(file,username,password)
	SaveUser()
}

func Login(){
	username:= ai("UserName: ")
	password:= ai("Password: ")
	file:=userdir+username
	if !qff.Exists(file) { CreateUser(file,username,password) }
}
