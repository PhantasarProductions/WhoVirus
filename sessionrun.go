package main

import(
	"fmt"
	"sort"
	"math"
	"strings"
	"trickyunits/qstr"
)



type tCommando struct {
	help string
	fun func(para []string)
}


var cmd = map[string] *tCommando {}
var running bool

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
		"Shows all files",
		func( para[] string ) {
			sorter:=[]string{}
			for n,_ := range user.ses.files{
				allow:=true // This variable will be used to allow more specific output
				if allow {
					sorter=append(sorter,n)
				}
			}
			sort.Strings(sorter)
			for _,n:=range sorter{
				v:=user.ses.files[n]
				fmt.Print(cya(qstr.Left(n+"                    ",20)+" "))
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
					user.totalscore=+total
					fmt.Println("\n\n");
					fmt.Println(yel("Sessions:      ")+cya(fmt.Sprintf("%d",user.sessions)))
					fmt.Println(yel("Total:         ")+cya(fmt.Sprintf("%d",user.totalscore)))
					fmt.Println(yel("Average Score: ")+cya(fmt.Sprintf("%d",int(math.Round(float64(user.totalscore/user.sessions))))))
					fmt.Println();
					running=false
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
}

func RunSession(){
	doing("Session has begun","\n Type either HELP or RULES for extra instructions")
	running=true
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
			} else {
				user.ses.revealed[p[0]]=true
				user.ses.runs++
				fmt.Println(yel(content))
			}
		} else {
			fmt.Println(red("ERROR! "),yel("Unknown command or file name"))
		}
		//fmt.Print(p[0]) // I must close this session now, but I don't want parse errors. :)
	}
}
