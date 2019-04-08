package console

import (
	"fmt"
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
	fmt.Printf("Sending Emails \n")
}
