/*
	Who is the Virus?
	Ansi functions
	
	
	
	(c) Jeroen P. Broks, 2018, All rights reserved
	
		This program is free software: you can redistribute it and/or modify
		it under the terms of the GNU General Public License as published by
		the Free Software Foundation, either version 3 of the License, or
		(at your option) any later version.
		
		This program is distributed in the hope that it will be useful,
		but WITHOUT ANY WARRANTY; without even the implied warranty of
		MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
		GNU General Public License for more details.
		You should have received a copy of the GNU General Public License
		along with this program.  If not, see <http://www.gnu.org/licenses/>.
		
	Exceptions to the standard GNU license are available with Jeroen's written permission given prior 
	to the project the exceptions are needed for.
Version: 18.05.08
*/
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
