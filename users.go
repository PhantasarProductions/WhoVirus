package main

import "trickyunits/dirry"
import "trickyunits/qff"
import "os"

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
	
}

func NewUser(f,uname,pass)
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
