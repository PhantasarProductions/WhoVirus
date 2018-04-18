package main


import "trickyunits/ansistring"
import "trickyunits/qstr"
import "fmt"
import "strings"

func mag(s string) string{ return ansistring.SCol(s,ansistring.A_Magenta,0) }
func yel(s string) string{ return ansistring.SCol(s,ansistring.A_Yellow,0) }
func cya(s string) string{ return ansistring.SCol(s,ansistring.A_Cyan,0) }
func red(s string) string{ return ansistring.SCol(s,ansistring.A_Red,0) }


func wmag(s string) { fmt.Print(mag(s)) }
func wcya(s string) { fmt.Print(cya(s)) }
func wyel(s string) { fmt.Print(yel(s)) }
func wred(s string) { fmt.Print(red(s)) }



func ai(question string) string{
	wmag(question)
	//fmt.Print(ansistring.SCol(question,ansistring.A_Magenta,0)+" ")
	fmt.Print(ansistring.ISCol(ansistring.A_Cyan,0))
	ret:=qstr.RawInput("")
	fmt.Print(ansistring.ANUL())
	return ret
}


func yes(question string) bool{
	wmag(question)
	fmt.Print(ansistring.SCol(" ? ",ansistring.A_Cyan,ansistring.A_Blink))
	answer:=ai("(Y/N) ")
	answer =qstr.MyTrim(answer)
	answer =strings.ToUpper(answer)
	return qstr.Left(answer,1)=="Y"
}

func doing(a,b string){
	wyel(a+" ")
	wcya(b)
	fmt.Print("\n")
}
