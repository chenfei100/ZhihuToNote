ZhihuToEvernote
===============

将知乎收藏自动发送到Evernote/印象笔记中；可以是自己的知乎收藏也可以是其他人的收藏内容
Python版本请移步到[zhihu_to_evernote](https://github.com/huaisha1224/zhihu_to_evernote)

####下载第三方库

- "github.com/opesun/goquery"
- "github.com/widuu/goini"


####使用说明
	
- 1. 修改`config.ini`配置文件的内容；
- 2. 填写发送邮件地址和host以及账号密码
- 3. 填写你的Evernote/印象笔记邮件账号地址
- 4. 填写你在Evernote/印象笔记中的笔记本
- 5. 将`ZhihuToEvernote.exe`和`config.ini`放在同一个目录下
- 6. 然后运行`ZhihuToEvernote.exe`


####`config.ini`配置文件
	
	[info]
	Url = http://www.zhihu.com/collection/20261977
	MailHost = smtp.126.com
	MailUser = user*****@126.com
	MailPassword = password
	EvernoteMail = user***@m.yinxiang.com
	Notebook = 知乎收藏文章

####`config.ini`配置文件说明
	
	[info]
	Url = http://www.zhihu.com/collection/20261977
	;知乎收藏地址
	MailHost = smtp.126.com
	;你的邮件的smtp地址；从你邮件服务商那里得到
	MailUser = huser*****@126.com
	;发送邮件账号
	MailPassword = password
	;你的Evernote/印象笔记邮件账号
	EvernoteMail = user***@m.yinxiang.com
	;Evernote/印象笔记上的笔记本
	Notebook = 知乎收藏文章

