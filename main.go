package main

import (
	"fmt"
	"log"
	"os"

	"github.com/liyuan/weekly/mail"
	"github.com/liyuan/weekly/subjoin"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "weekly"
	app.Usage = "请修改 weekly.json 中发送周报详情"
	app.Action = func(c *cli.Context) error {
		config, getconfigErr := mail.GetConfig()
		if getconfigErr != nil {
			return getconfigErr
		}
		Exc := subjoin.Excel{
			Start:       config.StartTime,
			End:         config.EndTime,
			Content:     config.Content,
			NextContent: config.NextWeeklyContent,
			Auth:        config.Auth,
			TeamTitle:   config.TeamTitle,
		}
		fmt.Println("发送中请稍候...")
		if err := Exc.NewExcel(); err != nil {
			return err
		}
		config.FileNameDir = Exc.GetFileDir()
		config.Title = Exc.GetFileName()
		config.Send()

		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
