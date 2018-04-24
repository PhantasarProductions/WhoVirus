package main

import(
	"fmt"
	"sort"
	"math"
	"strings"
	"trickyunits/qstr"
	conv "strconv"
)



type tCommando struct {
	help string
	fun func(para []string)
}


var cmd = map[string] *tCommando {}
var running bool
var goquit  bool

func init(){
	cmd["HELP"] = &tCommando{
		"Prints this help",
		func( para [] string){
			for k,v := range cmd {
				fmt.Println(cya(k))
				fmt.Println("\t"+yel(v.help))
			}
		},
	}
	cmd["SCORES"] = &tCommando{
		"Tells you the total scores of your previous sessions",
		func ( para[] string ) {
			fmt.Println(yel("Sessions:    "),cya(fmt.Sprintf("%d",user.sessions)))
			fmt.Println(yel("Total Score: "),cya(fmt.Sprintf("%d",user.totalscore)))
		},
	}	
	cmd["RULES"] = &tCommando{
		"Tells you the rules of the game",
		func ( para[] string) {
			fmt.Println("In this game your computer is infected with a virus")
			fmt.Println("One file is the virus, the others are 99 files contain a hint towards the virus")
			fmt.Println("Delete the file that is the virus and you win, but here's the rub!")
			fmt.Println("1.\tAll hints you are given are lies!")
			fmt.Println("2.\tRunning a file will display the hint, but give you one penalty point")
			fmt.Println("3.\tRunning the virus => Game over!")
			fmt.Println("4.\tDeleting a file that is not the virus, will get you 10 penalty points")
			fmt.Println("The trick is to delete the file that the virus with as little penalty points as possible")		
			fmt.Println("\nYou can type the name of the file by typing its name (case sensitive)")
			fmt.Println("You can use the commands in order to do extra stuff. The \"HELP\" command will tell you which commands you have")
			fmt.Println("\n\n\t\tGood luck")
		},
	}
	cmd["DIR"] = &tCommando{
		"Shows all files\n\t- When used without parameters all files are shown\n\tWhen a parameter is prefixed with a * you will see all files suffixed with that. You can also use * as a suffix for the other effect\n\t- ! can be used as the anti-*. Working is the same if the prefix/suffix are NOT what is asked\n\t- %<number> displays all files with that number of letters\n\t\n\tPlease note, of all files you've seen the content will be displayed as well when using this comment",
		func( para[] string ) {
			sorter:=[]string{}
			for n,_ := range user.ses.files{
				allow:=true // This variable will be used to allow more specific output
				if len(para)>0 {
					//allow=false
					for _,p:=range para{
						if        qstr.Suffixed(p,"*") { allow=allow && qstr.Prefixed(n,qstr.Left (p,len(p)-1)) 
						} else if qstr.Prefixed(p,"*") { allow=allow &&  qstr.Suffixed(n,qstr.Right(p,len(p)-1)) 
						} else if qstr.Suffixed(p,"!") { allow=allow && !qstr.Prefixed(n,qstr.Left (p,len(p)-1)) 
						} else if qstr.Prefixed(p,"!") { allow=allow && !qstr.Suffixed(n,qstr.Right(p,len(p)-1)) 
						} else if qstr.Prefixed(p,"%") {
							a:=qstr.Right(p,len(p)-1)
							v,e:=conv.ParseInt(a,10,32)
							if e!=nil { fmt.Println(red("ERROR! "),yel(e.Error())); return }
							allow = allow && int64(len(n))==v
						} else if p=="boys"  { allow = allow && Sex[n]=="M"
						} else if p=="girls" { allow = allow && Sex[n]=="F"
						} else { allow=allow && p==n}
					}
				}
				if allow {
					sorter=append(sorter,n)
				}
			}
			sort.Strings(sorter)
			for _,n:=range sorter{
				v:=user.ses.files[n]
				fmt.Print(mag(qstr.Right(fmt.Sprintf("     %d",len(n)),4))," ")
				if Sex[n]=="F"{
					fmt.Print(red(qstr.Left(n+"                    ",20)+" "))
				} else {
					fmt.Print(cya(qstr.Left(n+"                    ",20)+" "))
				}
				if user.ses.revealed[n] {
					fmt.Print(yel(v))
				}
				fmt.Println("")
			}
		},
	}
	cmd["LS"] = cmd["DIR"]
	cmd["DEL"] = &tCommando{
		"Deletes a file",
		func( para[] string ) {
			if len(para)<1 { 
				fmt.Println(red("ERROR! "),yel("Invalid input"))
				return
			}
			file:=para[0]
			if content,ok:=user.ses.files[file];ok{
				fmt.Println(cya(file)+" "+yel("has been deleted"))
				if content=="*VIRUS*"{
					fmt.Println(yel("CONGRATULATIONS! YOU KILLED THE VIRUS!"))
					total:=0
					fmt.Println(yel("Files run  x 1: ")+cya(fmt.Sprintf("%d",user.ses.runs))); total += user.ses.runs
					fmt.Println(yel("Deletions  x10: ")+cya(fmt.Sprintf("%d",user.ses.deletions))); total += (user.ses.deletions*10)
					fmt.Println();
					fmt.Println(yel("Total Score:    ")+cya(fmt.Sprintf("%d",total)))
					user.insession=false
					user.sessions++
					user.totalscore+=total
					fmt.Println("\n\n");
					fmt.Println(yel("Sessions:      ")+cya(fmt.Sprintf("%d",user.sessions)))
					fmt.Println(yel("Total:         ")+cya(fmt.Sprintf("%d",user.totalscore)))
					fmt.Println(yel("Average Score: ")+cya(fmt.Sprintf("%d",int(math.Round(float64(user.totalscore/user.sessions))))))
					fmt.Println();
					running=false
					user.insession=false
				} else {
					fmt.Println(mag("Unfortunately the file you deleted was not the virus"))
					user.ses.deletions++
				}
				delete ( user.ses.files,file )
			} else {
				fmt.Println(red("ERROR! "),yel("File not found"))
			}
		},
	}
	cmd["UNLINK"]=cmd["DEL"]
	cmd["RM"]=cmd["DEL"]
	cmd["SAVE"]=&tCommando{
		"Saves user data. (In case you fear a computer malfunction)",
		func( para[] string) { user.insession=true; SaveUser() },
	}
	cmd["QUIT"]=&tCommando{
		"Saves current session and quits",
		func(para[] string) { running=false; goquit=true },
	}
	cmd["BYE"]=cmd["QUIT"]
	cmd["EXIT"]=cmd["QUIT"]
}

func RunSession(){
	doing("Session has begun","\n Type either HELP or RULES for extra instructions")
	running=true
	user.insession=true
	for running {
		c:=ai("Ok > ")
		p:=strings.Split(c," ")
		para:=[]string{}
		for i:=1;i<len(p);i++ {para=append(para,p[i])}
		opdracht:=strings.ToUpper(p[0])
		if _,ok:=cmd[opdracht];ok {
			 cmd[opdracht].fun(para)
		} else if content,ok:=user.ses.files[p[0]]; ok {
			if content=="*VIRUS*" {
				fmt.Println(red("I AM THE VIRUS! I'VE DELETED YOUR ENTIRE SYSTEM!\n\nGAME OVER!!!"))
				running=false
				user.insession=false
			} else {
				user.ses.revealed[p[0]]=true
				user.ses.runs++
				fmt.Println(yel(content))
			}
		} else {
			fmt.Println(red("ERROR! "),yel("Unknown command or file name"))
		}
		//fmt.Print(p[0]) // I must close this session now, but I don't want parse errors. :)
		if !running && !goquit {
			running=yes("Start a new session")
			if running { CreateSession(); user.insession=true; SaveUser() }
		}
	}
	SaveUser()
}
