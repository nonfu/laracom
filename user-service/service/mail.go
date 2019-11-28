package service

import (
    "errors"
    "log"
    "net/smtp"
    "os"
)

type MailService struct {
    host string
    port string
    tls bool
    user string
    password string
}

func (srv *MailService) newInstance() *MailService {
    host := os.Getenv("MAIL_HOST")
    port := os.Getenv("MAIL_PORT")
    tls := os.Getenv("MAIL_TLS") == "true"
    user := os.Getenv("MAIL_USER")
    password := os.Getenv("MAIL_PASSWORD")

    return &MailService{host, port, tls, user, password}
}

func (srv *MailService) SendMail(to []string, subject, message []byte) error {
    if to == nil {
        return errors.New("收件人不能为空")
    }
    auth := smtp.PlainAuth("", srv.user, srv.password, srv.host)
    err := smtp.SendMail(srv.host + ":" + srv.port, auth, srv.user, to, message)
    if err != nil {
        log.Fatalf("邮件发送失败: %v\n", err)
        return errors.New("邮件发送失败")
    }
    return nil
}
