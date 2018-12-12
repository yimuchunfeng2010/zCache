package tool

import (
	"ZCache/global"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strings"
	"time"
)

//代码源自https://www.jianshu.com/p/45c05c7c2111
func main() {
	// 邮箱地址
	UserEmail := global.Config.UserEmail
	// 端口号，:25也行
	Mail_Smtp_Port := global.Config.MailPort
	//邮箱的授权码，去邮箱自己获取
	Mail_Password := global.Config.MailAuthCode
	// 此处填写SMTP服务器
	Mail_Smtp_Host := global.Config.MailSmtpHost
	auth := smtp.PlainAuth("", UserEmail, Mail_Password, Mail_Smtp_Host)
	to := []string{global.Config.ToMail}
	nickname := "发送人名称"
	user := UserEmail

	subject := "标题"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := "邮件内容."
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(Mail_Smtp_Host+":"+Mail_Smtp_Port, auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}

type SendMail struct {
	user     string
	password string
	host     string
	port     string
	auth     smtp.Auth
}

type Attachment struct {
	name        string
	file        []byte
	contentType string
}

type Message struct {
	from        string
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	contentType string
	attachment  []Attachment
}

func (mail *SendMail) Auth() {
	mail.auth = smtp.PlainAuth("", mail.user, mail.password, mail.host)
}

func (mail SendMail) Send(message Message) error {
	mail.Auth()
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	Header := make(map[string]string)
	Header["From"] = message.from
	Header["To"] = strings.Join(message.to, ";")
	Header["Cc"] = strings.Join(message.cc, ";")
	Header["Bcc"] = strings.Join(message.bcc, ";")
	Header["Subject"] = message.subject
	Header["Content-Type"] = "multipart/related;boundary=" + boundary
	Header["Date"] = time.Now().String()
	mail.writeHeader(buffer, Header)

	var imgsrc string

	//多图片发送
	for _, v := range message.attachment {
		attachment := "\r\n--" + boundary + "\r\n"
		attachment += "Content-Transfer-Encoding:base64\r\n"
		attachment += "Content-Type:" + v.contentType + ";name=\"" + v.name + "\"\r\n"
		attachment += "Content-ID: <" + v.name + "> \r\n"
		buffer.WriteString(attachment)

		//拼接成html
		//imgsrc += "<p><img src=\"cid:" + graphname + "\" height=200 width=300></p><br>\r\n\t\t\t"

		defer func() {
			if err := recover(); err != nil {
				fmt.Printf(err.(string))
			}
		}()
		//mail.writeFile(buffer, v.file)
		mail.writeFileByte(buffer, v.file)

	}

	//需要在正文中显示的html格式
	var template = `
    <html>
        <body>
            <p>text:%s</p><br>
            %s
        </body>
    </html>
    `
	var content = fmt.Sprintf(template, message.body, imgsrc)
	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type: text/html; charset=UTF-8 \r\n"
	body += content
	buffer.WriteString(body)
	buffer.WriteString("\r\n--" + boundary + "--")
	fmt.Println(buffer.String())
	smtp.SendMail(mail.host+":"+mail.port, mail.auth, message.from, message.to, buffer.Bytes())
	return nil
}

func (mail SendMail) writeHeader(buffer *bytes.Buffer, Header map[string]string) string {
	header := ""
	for key, value := range Header {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return header
}

func (mail SendMail) writeFile(buffer *bytes.Buffer, fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(payload, file)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}

}

func (mail SendMail) writeFileByte(buffer *bytes.Buffer, data []byte) {
	fmt.Println("len(data)", len(data))
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(payload, data)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
	buffer.WriteString("\r\n")
}
