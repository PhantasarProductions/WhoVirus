package main


import "trickyunits/ansistring"
import "trickyunits/qstr"
import "fmt"
import "strings"

func mag(s string) string{ return ansistring.SCol(s,ansistring.A_Magenta,0) }
func cya(s string) string{ return ansistring.SCol(s,ansistring.A_Cyan,0) }

func wmag(s string) { fmt.Print(mag(s)) }



func ai(question string) string{
	wmag(question)
	//fmt.Print(ansistring.SCol(question,ansistring.A_Magenta,0)+" ")
	ret:=qstr.RawInput(ansistring.ICol(ansistring.A_Cyan,0,0))
	fmt.Print(ansistring.ANUL())
	return ret
}


func yes(question string) bool{
	wmag(question)
	fmt.Print(ansistring.SCol(" ? ",ansistring.A_Cyan,ansistring.A_Blink))
	answer:=ai("(Y/N) ")
	answer=qstr.MyTrim(answer)
	answer=strings.ToUpper(answer)
	return qstr.Left(answer,1)=="Y"
}
