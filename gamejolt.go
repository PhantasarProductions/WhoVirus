/*
	Who is the Virus?
	Game Jolt functions
	
	
	
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

import "github.com/TrickyGameJolt/GoGameJolt"
import "fmt"

type tGJAuth struct{
	gameid string
	privatekey string
	scoretable string
	ach map[int] string
	fail string
}

var GJAuth *tGJAuth


var vgjuser *gj.GJUser

func GJLogin(usname,token string) bool{
	if GJAuth==nil {
		wred("Error! ")
		wyel("This version does not support Game Jolt")
		wcya("\n\n")
		return false
	}
	doing("Logging in to Game Jolt as ",usname)
	vgjuser = gj.Login(GJAuth.gameid,GJAuth.privatekey,usname,token)
	return vgjuser.LoggedIn
}


// Victory submission
func GJSubmit(score int) {
	if GJAuth==nil { return }
	if user.gjuser=="" { // guest submission
		gj.SubmitGuestScore(user.name,GJAuth.gameid,GJAuth.privatekey,fmt.Sprintf("%d points",score),fmt.Sprintf("%d",score),GJAuth.scoretable)
	} else {
		// score submission for user
		vgjuser.SubmitScore(fmt.Sprintf("%d points",score),fmt.Sprintf("%d",score),GJAuth.scoretable)
		// Trophy submission
		for rq,aid := range GJAuth.ach{
			if score<rq { vgjuser.AwardTrophy(aid) }
		}
	}
}

func GJFail(){
	if GJAuth==nil { return }
	if user.gjuser=="" { return }
	vgjuser.AwardTrophy(GJAuth.fail)
}
