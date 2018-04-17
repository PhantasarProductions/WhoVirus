package main

import(
	"fmt",
	"strings"
)



type tCommando struct {
	help string
	fun func(para []string)
}


var cmd = map[string] tCommando {}

func init(){
	cmd["HELP"] = tCommando{
		"Prints this help",
		func( para [] string){
			for k,v := range cmd {
				fmt.Println(k)
				fmt.Println("\t"+v.help)
			}
		},
	}
	cmd["RULES"] = tCommando{
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
}

func RunSession(){
	doing("Session has begun","\n Type either HELP or RULES for extra instructions")
	for{
		c:=ai("Ok > ")
		p:=strings.split(c," ")
		fmt.Print(p[0]) // I must close this session now, but I don't want parse errors. :)
	}
}
