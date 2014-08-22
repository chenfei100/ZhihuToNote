package main

import (
	"fmt"
	"github.com/opesun/goquery"
	"github.com/widuu/goini"
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
	pTitle := p.Find("title").Text()
	pHtml := p.Html()
	fmt.Print(pTitle)
	//subject := pTitle
	body := pHtml
	subject := "This is the email body."
	GetConf(subject, body)
}

//获取到config.ini里面的配置文件
func GetConf(subject, body string) {
	conf := goini.SetConfig("./config.ini")
	mailHost := conf.GetValue("info", "MailHost") + ":25"
	mailUser := conf.GetValue("info", "MailUser")
	mailPassword := conf.GetValue("info", "MailPassword")
	evernoteMail := conf.GetValue("info", "EvernoteMail")
	//notebook := conf.GetValue("info", "Notebook")

	fmt.Println(mailHost)
	fmt.Println(mailUser)
	fmt.Println(mailPassword)
	fmt.Println(evernoteMail)
	//fmt.Println(notebook)
	fmt.Println(subject)
	//fmt.Println(body)
	//subject := "This is the email body."

	fmt.Println("send email")
	err := SendToEvernote(mailUser, mailPassword, mailHost, evernoteMail, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}

}

//发送邮件到Evernote
func SendToEvernote(user, password, host, to, subject, body, mailtype string) error {
	hostPort := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hostPort[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " +
		subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	conf := goini.SetConfig("./config.ini")
	url := conf.GetValue("info", "Url")
	GetZhihuQuestionList(url + "?page=")
	fmt.Print(len(urlList))

	//subject := "This is the email body."
	//body := `
	//   <html>
	//   <body>
	//   <h3>
	//   "Test send email by golang"
	//   </h3>
	//   </body>
	//   </html>
	//   `
	//fmt.Println("send email")
	//err := SendToEvernote(mailUser, mailPassword, mailHost, evernoteMail, subject, body, "html")
	//if err != nil {
	//	fmt.Println("send mail error!")
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("send mail success!")
	//}

}
