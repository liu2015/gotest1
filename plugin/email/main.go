package email

import "ginserver/global"

type emailPlugin struct{}

func CreateEmailPlug(To, From, Host, Secret, Nickname, string, Port int, IsSSL bool) *emailPlugin {
	global.GlobalConfig.To = To

}
