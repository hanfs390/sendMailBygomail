package main

import (
	"sendMailBygomail/tool"
	"fmt"
)

func main() {
	/* login qq mail */
	var em = &tool.Email{
		From:     "QQ number@qq.com",
		Host:     "smtp.qq.com",
		Port:     465, //使用SSL，端口号465或587
		UserName: "QQ nubmer",
		Password: "authorization code", //from qq mail - 设置 - 帐户 - POP3/IMAP/SMTP/Exchange/CardDAV/CalDAV服务 - open IMAP/SMTP服务
	}

	hb := "<h3>hello,世界</h3>" //支持html

	/* send to others mail */
	/**
	 * @subject: the subject of this email
	 * @htmlBody: the body of content
	 * @attachFile: the path of file
	 * @rename: rename the attachFile
	 * @to: the dst address. multi-addresses are permitted
	 */
	err := em.Send("title", hb, "./test.txt", "test1.txt", "test@qq.com") //可以多个接收方
	if err != nil {
		fmt.Println(err)
	}
}