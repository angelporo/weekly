package main

import (
	"errors"
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
	app.Usage = "example: weekly <*starttime 2018-12-23> <*endtime 2019-04-12> <*content 周报主要内容> <*nextContent 下周纪要>"

	app.Action = func(c *cli.Context) error {
		start := c.Args().Get(0)
		end := c.Args().Get(1)
		content := c.Args().Get(2)
		nextContent := c.Args().Get(3)
		if start == "" || end == "" || content == "" || nextContent == "" {
			return errors.New("你难道不会看看命令行帮助吗?")
		}
		fmt.Println("发送中...")
		config, getconfigErr := mail.GetConfig()
		if getconfigErr != nil {
			return getconfigErr
		}
		Exc := subjoin.Excel{
			Start:       c.Args().Get(0),
			End:         c.Args().Get(1),
			Content:     c.Args().Get(2),
			NextContent: c.Args().Get(3),
			Auth:        config.Auth,
		}
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
