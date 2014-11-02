package main

import (
	"encoding/base64"
	"fmt"
	"github.com/opesun/goquery"
	"github.com/widuu/goini"
	"net/mail"
	"net/smtp"
	"strconv"
	//"strings"
)

var urlList = []string{}

func GetZhihuQuestionList(url string) {
	/*
	 *  通过传入的url、分析提取url里面的问题列表
	 *  用for循环提交分页URL地址
	 *  并用goquery查找页面内容是否存在、用以判断是否还有分页
	 *  用提取到的短URL加上统一URL地址头得到某一个完整URL地址
	 */
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
					GetSubjectBody(singleUrl)
				}
			} else { //如果text的长度小于0表示没有找到
				fmt.Print(".........没有文章了.......\n")
				break
			}
		}
	}

}

func GetSubjectBody(url string) {
	/*
	 *  通过传入的单个问题url、分析提取url里面的问题标题和内容
	 *  用goquery查找页面提取里面title作为邮件的subject、用html()作为body
	 *
	 */
	//var url = "http://www.zhihu.com/question/24859069"
	p, error := goquery.ParseUrl(url)
	if error != nil {
		panic(error)
	}
	subject := p.Find("title").Text()
	body := p.Html()
	fmt.Print(subject)
	GetConf(subject, body)
}

func GetConf(subject, body string) {
	/*
	 *  用goini第三方库读取config.ini配置文件获取配置内容
	 *  用于发送邮件
	 *
	 */
	var receiveMail string
	conf := goini.SetConfig("./config.ini")
	mailHost := conf.GetValue("info", "MailHost")
	mailUser := conf.GetValue("info", "MailUser")
	mailPassword := conf.GetValue("info", "MailPassword")
	note := conf.GetValue("note", "Note")

	//判断用户的笔记软件类型
	if note == "evernote" {
		receiveMail = conf.GetValue("evernote", "ReceiveMail")
		notebook := conf.GetValue("evernote", "Notebook")
		subject += "@" + notebook
	} else if note == "wiz" {
		receiveMail = conf.GetValue("wiz", "ReceiveMail")
	} else if note == "onenote" {
		receiveMail = "me@onenote.com"
	} else if note == "youdao" {
		receiveMail = "save@note.youdao.com"
	}

	//调用发送邮件函数并传递参数
	fmt.Println("Save Note")
	err := SendToNote(
		mailUser,
		mailPassword,
		mailHost,
		receiveMail,
		subject,
		body,
	)
	if err != nil {
		fmt.Println("Save Note Error!")
		fmt.Println(err)
	} else {
		fmt.Println("Save Note Success!")
	}

}

func SendToNote(user, password, host, to, subject, body string) error {
	/*
	 *发送邮件到Note
	 */

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
	//发送邮件
	auth := smtp.PlainAuth("", user, password, host)
	err := smtp.SendMail(
		host+":25",
		auth,
		user,
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
