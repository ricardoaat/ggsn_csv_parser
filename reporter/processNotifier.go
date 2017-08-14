package reporter

import (
	"fmt"

	"github.com/ricardoaat/ggsn_csv_parser/config"

	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func notifyTeam(message string) {
	log.Info(fmt.Sprintf("Notifying Team: %s", message))

	for _, r := range config.Conf.Notif.Recipients {
		s := `GGSN CSV's parsed REPORTER`
		sendMsjToEmail(message, r, s)
	}
}

func sendMsjToEmail(b, r, s string) {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Conf.Notif.FromAddrs)
	m.SetHeader("To", r)
	m.SetHeader("Subject", s)
	m.SetBody("text/html", b)

	d := gomail.Dialer{Host: "127.0.0.1", Port: 25}

	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}
