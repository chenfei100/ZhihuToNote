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
- 3. 填写你的Evernote/印象笔记/wiz邮件账号地址
- 4. 填写你在Evernote/印象笔记中的笔记本；
- 5. 将`ZhihuToNote.exe`和`config.ini`放在同一个目录下
- 6. 然后运行`ZhihuToNote.exe`


####OneNote用户需要进行如下设置
	在OneNote上添加电子邮件地址；也就是确保你填写在config.ini里面MailUser的地址出现在OneNote的邮件列表中
![OneNote](http://github001.qiniudn.com/note/OneNote.png)


####有道笔记用户设置如下
	登录你的有道笔记网页版本在设置界面中；确保你填写在config.ini里面MailUser的地址出现在有道的邮件列表中
![有道笔记](http://github001.qiniudn.com/note/youdao.png)


####`config.ini`配置文件
	
	[info]
	Url = http://www.zhihu.com/collection/20261977
	MailHost = smtp.126.com
	MailUser = user*****@126.com
	MailPassword = password
	
	[note]
	Note = evernote
	
	[evernote]
	ReceiveMail = evernote*****@m.yinxiang.com
	Notebook = 知乎收藏文章
	
	[wiz]
	ReceiveMail = wiz*****@mywiz.cn
	
	[onenote]
	ReceiveMail = me@onenote.com 
	
	[youdao]
	ReceiveMail = save@note.youdao.com


####`config.ini`配置文件说明
	
	[info]
	Url = http://www.zhihu.com/collection/20261977
	;知乎收藏地址
	MailHost = smtp.126.com
	;你的邮件的smtp地址；从你邮件服务商那里得到
	MailUser = huser*****@126.com
	;发送邮件账号
	MailPassword = password
	;笔记类型；evernote/onenote/youdao/wiz全部小写

	[note]
	Note = evernote

	;根据上面Note里面填写的笔记类型填写下面对应的内容
	;evernote支持发送到指定笔记本
	[evernote]
	ReceiveMail = evernote*****@m.yinxiang.com
	Notebook = 知乎收藏文章
	
	[wiz]
	ReceiveMail = wiz*****@mywiz.cn
	;onenote接收邮件地址是固定的不需要修改
	[onenote]
	ReceiveMail = me@onenote.com 
	;youdao接收邮件地址是固定的不需要修改
	[youdao]
	ReceiveMail = save@note.youdao.com


####注意事项1

- 1. 如果你不想安装Go开发环境、可以直接下载`Build/ZhihuToNote.exe`和`config.ini`文件
- 2. `config.ini`文件必须是`utf-8`无BOM格式

