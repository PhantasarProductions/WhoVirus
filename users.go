package main

import "trickyunits/dirry"
import "trickyunits/qff"

type tuser{
	name string
	file string
	password string
	ses session
	totalscore int
	sessions int
	ansi bool
}

var user tuser
var userdir = dirry("$AppSupport$/$LinuxDot$Phantasar Productions/Who Is The Virus/")

func CreateUser(file,username,password){
}

func Login(){
	username:= ai("UserName: ")
	password:= ai("Password: ")
	file:=userdir+username
	if !qff.isfile(file) { CreateUser(file,username,password) }
}
