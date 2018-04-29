package main

import "github.com/TrickyGameJolt/GoGameJolt"

type tGJAuth struct{
	gameid string
	privatekey string
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
