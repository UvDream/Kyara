package utils

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"server/global"
	"server/model"
	"strconv"
)

// @title   SendEmail
// @description   发送邮件
// @auth  （2020/12/16  11:30）
// @param     path     string
// @return    err     error
func SendEmail(title string, body string) (msg string, err error) {
	var sysConfig model.SysConfig
	db := global.GVA_DB
	err = db.Find(&sysConfig).Error
	//接收邮箱设置
	mailTo := []string{
		sysConfig.ReceiveEmail,
	}
	fmt.Println(mailTo)
	//定义邮箱服务器信息
	mailConn := map[string]string{
		"user": sysConfig.EmailUser,
		"pass": sysConfig.EmailPassword,
		"host": sysConfig.EmailHost,
		"port": sysConfig.EmailPort,
	}
	port, _ := strconv.Atoi(mailConn["port"]) //port转换为int
	fmt.Println(mailConn, port)
	email := gomail.NewMessage()
	email.SetHeader("From", sysConfig.DomainName+"<"+mailConn["user"]+">")
	email.SetHeader("To", mailTo...)
	email.SetHeader("Subject", title) //设置邮件主题
	email.SetBody("text/html", body)  //设置邮件正文
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err = d.DialAndSend(email)
	if err != nil {
		return "发送邮件失败,请联系开源项目作者", err
	}
	return "发送成功", err
}
