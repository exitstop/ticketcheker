package sendemail

import (
	"log"
	"net/smtp"
)

func main() {
  Send("hello world" ,  "exitstop@list.ru", "gbe643412@gmail.com", "fgjkriJDdjrjhfhIF73hfd" )
}

func Send(body string, sendTo string, user string,pass string) {

	msg := "From: " + user + "\n" +
		"To: " + sendTo + "\n" +
		"Subject: Ticket Sender\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", user , pass, "smtp.gmail.com"),
		user , []string{sendTo}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

  log.Print("sent: " + body)
}
