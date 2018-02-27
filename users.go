package main

import "trickyunits/dirry"
import "trickyunits/qff"

type tuser struct{
	name string
	file string
	password string
	ses session
	totalscore int
	sessions int
	ansi bool
}

var user tuser
var userdir = dirry.Dirry("$AppSupport$/$LinuxDot$Phantasar Productions/Who Is The Virus/")

func CreateUser(file,username,password string){
}

func Login(){
	username:= ai("UserName: ")
	password:= ai("Password: ")
	file:=userdir+username
	if !qff.Exists(file) { CreateUser(file,username,password) }
}
