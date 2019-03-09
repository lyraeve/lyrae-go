package command

import (
	"fmt"
	"github.com/lyraeve/lyrae-go/finder"
	"github.com/urfave/cli"
)

var (
	FindCommand = cli.Command{
		Name:      "find",
		Usage:     "查詢",
		ArgsUsage: "<番號>",
		Action: func(c *cli.Context) error {
			sn := c.Args().Get(0)

			res, err := finder.FindByNumber(sn)

			if err != nil {
				return err
			}

			fmt.Println(res)

			return nil
		},
	}
)
