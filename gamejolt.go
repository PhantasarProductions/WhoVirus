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
