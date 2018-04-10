package main

import "trickyunits/dirry"
import "trickyunits/qff"
import "trickyunits/qstr"
import "trickyunits/ansistring"
import jcr6 "trickyunits/jcr6/jcr6main"
import _ "trickyunits/jcr6/jcr6lzma"
import "os"
import "fmt"
import "strings"

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
	wo+=fmt.Sprintf("TOTALSCORE %d\n",user.totalscore)
	wo+=fmt.Sprintf("SESSIONS %d\n",user.sessions)
	wo+="ANSI "+b2s(user.ansi)+"\n"
	// Session data
	// Writeout         
	e:=os.MkdirAll(userdir,0777)
	if e!=nil {
		panic (e)
	}
	j:=jcr6.JCR_Create(user.file,"BRUTE")
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

func LoadUser(f,un,pw string) bool{
	if jcr6.Recognize(f)=="NONE" {
		wred("ERROR! ")
		wyel("User "+un+" not recognized")
	}
	j:=jcr6.Dir(f)
	b:=jcr6.JCR_B(j,"User")
	d:=string(b)
	l:=strings.Split(d,"\n")
	user=tuser{}
	for li,ln:=range l {
		if ln!="" {
			p:=strings.Index(ln," ")
			if p<0 { panic(fmt.Sprintf("Invalid line in userdata in line %d",li)) }
			c:=qstr.MyTrim(ln[:p])
			a:=qstr.MyTrim(ln[p+1:])
			switch c{
				case "PW": 
					user.password=a
				case "INSESSION":
					user.insession=a=="true"
				case "ANSI":
					user.ansi=a=="true"
				default:
						wred("ERROR! ")
						wyel(fmt.Sprintf("I don't understand line %d -- %s",li,ln))
						fmt.Println("")
			}
		}
	}
	if user.password!=pw {
		wred("ERROR! ") 
		wyel("Incorrect password!\n")
	}
	return user.password==pw
}

func Login(){
	for {
		username:= ai("UserName: ")
		password:= ai("Password: ")
		file:=userdir+username
		mn:=!qff.Exists(file)
		if  mn { CreateUser(file,username,password) }
		if LoadUser(file,username,password) { 
			return
		}
	}
}
