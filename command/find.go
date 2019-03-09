package command

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/urfave/cli"
	"gopkg.in/h2non/gentleman.v2"
	"strings"
)

var (
	Sites = map[string]string {
		"JavBus": "https://www.javbus.com",
	}
)

var (
	FindCommand = cli.Command{
		Name:  "find",
		Usage: "查詢",
		ArgsUsage: "<番號>",
		Action: func(c *cli.Context) error {
			sn := c.Args().Get(0)

			client := gentleman.New()
			client.URL(Sites["JavBus"])

			req := client.Request()
			req.Path("/" + sn)

			res, err := req.Send()

			if err != nil {
				return err
			}

			if !res.Ok {
				fmt.Printf("Invalid server response: %d\n", res.StatusCode)
				return nil
			}

			htmlNode, err := htmlquery.Parse(strings.NewReader(res.String()))

			if err != nil {
				return err
			}

			list := htmlquery.Find(htmlNode, "//span[@class='header']/text()")

			for _, node := range list {
				fmt.Println(node.Data)
			}


			return nil
		},
	}
)
