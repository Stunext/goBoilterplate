package console

import (
	"log"

	"github.com/robfig/cron"
)

// Schedule Initialization
func Schedule() {
	Job := cron.New()
	Job.AddFunc("@daily", sendMails)
	Job.Start()

	log.Printf("Scheduled jobs loaded")
}

func sendMails() {
	log.Printf("Sending Emails \n")
}
