ZhihuToNote
===============

将知乎收藏自动发送到Evernote/印象笔记中/为知笔记；可以是自己的知乎收藏也可以是其他人的收藏内容；
后续有时间会支持OneNote以及有道笔记


Python版本请移步到[zhihu_to_evernote](https://github.com/huaisha1224/zhihu_to_evernote)

####下载第三方库

- "github.com/opesun/goquery"
- "github.com/widuu/goini"


####使用说明
	
- 1. 修改`config.ini`配置文件的内容；
- 2. 填写发送邮件地址和host以及账号密码
- 3. 填写你的Evernote/印象笔记邮件账号地址
- 4. 填写你在Evernote/印象笔记中的笔记本
- 5. 将`ZhihuToNote.exe`和`config.ini`放在同一个目录下
- 6. 然后运行`ZhihuToNote.exe`


####`config.ini`配置文件
	
	[info]
	Url = http://www.zhihu.com/collection/20261977
	MailHost = smtp.126.com
	MailUser = user*****@126.com
	MailPassword = password
	[note]
	Note = wiz
	ReceiveMail = ******@mywiz.cn
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
	[note]
	;如果想保存到为知笔记就填写wiz即可；wiz/evetnote
	Note = wiz
	;接收内容的邮件地址;wiz和印象笔记提供
	ReceiveMail = ******@mywiz.cn
	;如果是为知笔记的话这个可以不填写
	Notebook = 知乎收藏文章


####注意事项

- 1. 如果你不想安装Go开发环境、可以直接下载`Build/ZhihuToNote.exe`和`config.ini`文件
- 2. `config.ini`文件必须是`utf-8`无BOM格式
