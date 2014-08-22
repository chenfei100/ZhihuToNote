package main

import (
	"encoding/base64"
	"fmt"
	"github.com/opesun/goquery"
	"github.com/widuu/goini"
	"net/mail"
	"net/smtp"
	"strconv"
	"strings"
)

var urlList = []string{}

//var subject string
//var body string

//通过传入的url、分析提取url里面的问题列表
func GetZhihuQuestionList(url string) {
	urlHeader := "http://www.zhihu.com"
	for i := 1; i < 100; i++ {
		url := url + strconv.Itoa(i)
		fmt.Println(url)
		r, err := goquery.ParseUrl(url)
		if err != nil {
			panic(err)
		} else {
			text := r.Find(".zm-item-title a") //查找所有问题列表
			if text.Length() > 0 {
				//取到text里面的所有"href"属性的数据
				for i := 0; i < text.Length(); i++ {
					singleUrl := urlHeader + text.Eq(i).Attr("href")
					urlList = append(urlList, singleUrl)
					GetSubjectBody(singleUrl) //调用函数提取单个问题页面的title和内容
				}
			} else { //如果text的长度小于0表示没有找到
				fmt.Print("NO\n")
				break
			}
		}
	}

	//fmt.Println(urlList)
}

func GetSubjectBody(url string) {
	//var url = "http://www.zhihu.com/question/24859069"
	p, error := goquery.ParseUrl(url)
	if error != nil {
		panic(error)
	}
	subject := p.Find("title").Text()
	body := p.Html()
	fmt.Print(subject)
	//subject := "This is the email body."
	GetConf(subject, body)
}

//获取到config.ini里面的配置文件
func GetConf(subject, body string) {
	conf := goini.SetConfig("./config.ini")
	mailHost := conf.GetValue("info", "MailHost") // + ":25"
	mailUser := conf.GetValue("info", "MailUser")
	mailPassword := conf.GetValue("info", "MailPassword")
	//evernoteMail := conf.GetValue("info", "EvernoteMail")
	//notebook := conf.GetValue("info", "Notebook")
	evernoteMail := "279478776@qq.com"

	fmt.Println(mailHost)
	fmt.Println(mailUser)
	fmt.Println(mailPassword)
	fmt.Println(evernoteMail)
	fmt.Println(subject)

	fmt.Println("send email")
	//err := SendToEvernote(mailUser, mailPassword, mailHost, evernoteMail, subject, body, "html")
	err := SendToEvernote(mailUser, mailPassword, mailHost, evernoteMail, subject, body)
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}

}

//发送邮件到Evernote

func SendToEvernote(user, password, host, to, subject, body string) error {
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	from := mail.Address{user, user}
	toMail := mail.Address{to, to}
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = toMail.String()
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(subject)))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(body))
	auth := smtp.PlainAuth("", user, password, host)
	err := smtp.SendMail(
		host+":25",
		auth, user,
		[]string{toMail.Address},
		[]byte(message),
	)

	if err != nil {
		panic(err)
	}
	return err
}

func main() {
	conf := goini.SetConfig("./config.ini")
	url := conf.GetValue("info", "Url")
	GetZhihuQuestionList(url + "?page=")
	fmt.Print(len(urlList))

}
