package email

import (
	"admin_api/global"
	"admin_api/internal/data"
	"admin_api/internal/response"
	"admin_api/internal/service/store/mysql/account"
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"html/template"
	"net/http"
	"net/smtp"
	"os"
	"time"
)

type EmailTemplateData struct {
	Name       string
	Id         string
	ActivityId string
	ChangeType string
	Points     int
	Reason     string
}

var scheduler *gocron.Scheduler

type EmailCampaignApi struct {
}

type Email struct {
	Email string `json:"email" binding:"required,email"`
}

type TimeEmail struct {
	Time  int    `json:"time" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (e EmailCampaignApi) Crete(ctx *gin.Context) {
	var id *Email
	if err := ctx.ShouldBind(&id); err != nil {
		ctx.JSON(http.StatusOK, response.Error(1, err.Error()))
		//
		fmt.Println("11111", id.Email, err.Error())
		return
	}

	if id.Email == `` {
		ctx.JSON(http.StatusOK, response.Error(1, fmt.Errorf("the email id should not nil").Error()))
		fmt.Println("the email id should not nil")
	}
	//初始化数据库
	em := account.NewEmailService(data.NewEmailRepo(global.DB))

	if ok := em.Create(id.Email); ok {
		ctx.JSON(http.StatusOK, response.Error(1, errors.New("记录创建成功，已发送").Error()))
	}
	to := id.Email
	subject := "积分变动通知"
	data := EmailTemplateData{
		Name:       "John Doe",
		Id:         "123456",
		ActivityId: "789",
		ChangeType: "-100",
		Points:     100,
		Reason:     "参加活动",
	}

	err := sendEmail(to, subject, data)
	if err != nil {
		fmt.Println("Failed to send email:", err)
	} else {
		fmt.Println("Email sent successfully")
	}

	return
}

func sendEmail(to, subject string, data EmailTemplateData) error {

	var from = global.ServerConfig.EmailConfig.EmailID
	var password = global.ServerConfig.EmailConfig.Password
	var smtpHost = global.ServerConfig.EmailConfig.SmtpHost
	var smtpPort = global.ServerConfig.EmailConfig.SmtpPort

	auth := smtp.PlainAuth("", from, password, smtpHost)
	fmt.Println("00000000000000", global.ServerConfig.EmailConfig, "99999999")
	// Working Directory  os.Getwd() 可用于返回项目的根工作目录，然后与模板文件的内部路径连接
	wd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("目录：", wd)
	// Template
	// 使用相对路径
	tmpl, err := template.ParseFiles(wd + "./internal/api/email/template/email.html")
	if err != nil {
		fmt.Printf("Failed to parse template: %v", err)
		return err
	}

	// Render the email template with the data
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		fmt.Println(err)
		return err
	}

	// Generate the email message
	msg := bytes.NewBuffer(nil)
	msg.WriteString(fmt.Sprintf("From: %s\r\n", from))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", to))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: text/html; charset=utf-8\r\n")
	msg.WriteString("\r\n")
	msg.WriteString(buf.String())

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg.Bytes())
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Email sent successfully")
	return nil
}

func (e EmailCampaignApi) TimeSend(ctx *gin.Context) {
	var Time *TimeEmail
	if err := ctx.ShouldBind(&Time); err != nil {
		ctx.JSON(http.StatusOK, response.Error(1, err.Error()))
		//
		fmt.Println("11111", Time, err.Error())
		return
	}

	timezone, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(timezone)

	data := EmailTemplateData{
		Name:       "John Doe",
		Id:         "123456",
		ActivityId: "789",
		ChangeType: "-100",
		Points:     100,
		Reason:     "参加活动",
	}

	s.Every(Time.Time).Minutes().Do(func() {
		go sendEmail(Time.Email, "subject参数", data)
	})

	// 启动定时任务
	scheduler = s
	s.StartAsync()
	return
}

func (e EmailCampaignApi) CancelTimeSend(ctx *gin.Context) {
	// 获取前端参数
	cancel := ctx.Query("cancel")

	// 停止定时任务
	if cancel == "true" && scheduler != nil {
		scheduler.Clear()
		scheduler = nil
	}

	// 返回成功信息
	ctx.JSON(http.StatusOK, "取消成功")
	return
}
