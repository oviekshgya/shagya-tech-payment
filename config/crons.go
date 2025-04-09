package config

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"shagya-tech-payment/api/routes"
)

func initCrons() {
	c := cron.New()

	c.AddFunc("0 0 1 * *", func() {
		fmt.Println("cron job start add game product master")
		res := routes.CronsController.GetDataJson()
		fmt.Println("Add Game Product Master: ", res)
	})

	c.Start()
}
