package console

import (
	"fmt"

	"github.com/robfig/cron"
)

// Schedule Initialization
func Schedule() {
	Job := cron.New()
	Job.AddFunc("@daily", sendMails)
	Job.Start()
}

func sendMails() {
	fmt.Printf("Sending Emails \n")
}
