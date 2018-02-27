package main


import "trickyunits/ansistring"
import "trickyunits/qstr"
import "fmt"

func mag(s string) string{ return ansistring.SCol(s,ansistring.A_Magenta) }
func cya(s string) string{ return ansistring.SCol(s,ansistring.A_Cyan) }

func wmag(s string) { fmt.Print(wmag) }



func ai(question string) string{
	wmag(question)
	fmt.Print(ansistring.SCol(question,ansistring.A_Magenta)+" ")
	ret:=qstr.RawInput(ansistring.ICol(ansistring.A_Cyan,0,0))
	fmt.Print(ansistring.ANUL())
	return ret
}	
