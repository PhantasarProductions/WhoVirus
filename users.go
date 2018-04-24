package main

import "trickyunits/dirry"
import "trickyunits/qff"
import "trickyunits/qstr"
import "trickyunits/ansistring"
import jcr6 "trickyunits/jcr6/jcr6main"
import _ "trickyunits/jcr6/jcr6lzma"
import "os"
import "fmt"
import conv "strconv"
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
	if user.insession {
		fls:=""
		rev:=""
		for n,s := range user.ses.files {
			fls+=Sex[n]+" "+n+" = "+s+"\n"
			if user.ses.revealed[n] { rev+=n+"\n" }
		}
		j.AddString(fls,"Session/Files","BRUTE",0777,0,"Mr. Virus","Fear me!")
		j.AddString(rev,"Session/Revelations","BRUTE",0777,0,"Mr. Virus","And you'll never find out who created me :P")
		j.AddString(fmt.Sprintf("DELETIONS %d\nRUNS %d\n",user.ses.deletions,user.ses.runs),"Session/Data","BRUTE",0777,0,"Mr. Virus","I will cause a lot of damage!")
	}
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
	user.file=f
	user.name=un
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
				case "TOTALSCORE":
					i,e:=conv.ParseInt(a,10,32)
					if e!=nil { panic(e); }
					user.totalscore=int(i)
				case "SESSIONS": 
					i,e:=conv.ParseInt(a,10,32)
					if e!=nil { panic(e); }
					user.sessions=int(i)
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
	if user.insession {
		doing("Continuing","session")
		user.ses          = session{}
		user.ses.files    = map[string]string{}
		user.ses.revealed = map[string]bool{}
		Sex               = map[string]string{}
		fls:=string(jcr6.JCR_B(j,"Session/Files"))
		flslines:=strings.Split(fls,"\n")
		for li,ln:=range flslines{
			if ln!="" {
				dfs:=strings.Split(ln," = ")
				if len(dfs)!=2 {
					wred("ERROR!")
					wyel(fmt.Sprintf("Session files syntax error type 1 in line %d -- %s\n",li,ln))
					wcya("I'll start a new session in stead")
					user.insession=false
					return user.password==pw
				}
				vr:=strings.Split(dfs[0]," ")
				if len(vr)!=2 {
					wred("ERROR!")
					wyel(fmt.Sprintf("Session files syntax error type 2 in line %d -- %s\n",li,ln))
					wcya("I'll start a new session in stead")
					user.insession=false
					return user.password==pw
				}
				Sex[vr[1]] = vr[0]
				user.ses.files[vr[1]]=dfs[1]
				user.ses.revealed[vr[1]]=false
			}
		}
		rev:=string(jcr6.JCR_B(j,"Session/Revelations"))
		for _,n := range strings.Split(rev,"\n"){ user.ses.revealed[n]=true }
		b:=jcr6.JCR_B(j,"Session/Data")
		d:=string(b)
		l:=strings.Split(d,"\n")	
		for li,ln:=range l {
			if ln!="" {
				p:=strings.Index(ln," ")
				if p<0 { panic(fmt.Sprintf("Invalid line in session data in line %d",li)) }
				c:=qstr.MyTrim(ln[:p])
				a:=qstr.MyTrim(ln[p+1:])
				switch c{
				case "DELETIONS":
					i,e:=conv.ParseInt(a,10,32)
					if e!=nil { panic(e); }
					user.ses.deletions=int(i)
				case "RUNS":
					i,e:=conv.ParseInt(a,10,32)
					if e!=nil { panic(e); }
					user.ses.runs=int(i)
				default:
						wred("ERROR! ")
						wyel(fmt.Sprintf("I don't understand line %d -- %s",li,ln))
						fmt.Println("")
				}
			}
		}
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
			if !user.insession{
				CreateSession()
			}
			return
		}
	}
}
