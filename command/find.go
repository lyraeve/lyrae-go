package command

import (
	"fmt"
	"github.com/urfave/cli"
	"gopkg.in/h2non/gentleman.v2"
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

			print(Sites["JavBus"])
			print(sn)
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

			// fmt.Printf("Body: %s", res.String())

			return nil
		},
	}
)
