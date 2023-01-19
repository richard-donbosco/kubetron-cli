package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "kubetron-cli",
		Usage:       "Kubetron-cli is wrapper for kubectl",
		Description: "Kubectron-cli is wrapper for kubectl and it used to manage multiple kubernetes clusters",
		Authors: []*cli.Author{
			{
				Name:  "Richard",
				Email: "ricdon41@gmail.com",
			},
		},

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
				Value:   "~/.kube/config",
			},
		},

		Commands: []*cli.Command{
			{
				Name:        "get",
				Aliases:     []string{"g"},
				Usage:       "Get one or many resources",
				Description: "Get one or many resources of a specific type",
				Subcommands: []*cli.Command{
					{
						Name:    "cluster",
						Aliases: []string{"clusters", "c"},
						Usage:   "Get all the cluster",
						Action: func(c *cli.Context) error {
							fmt.Printf("Get all the clusters")
							return nil
						},
					},
				},
			},
		},
	}

	_ = app.Run(os.Args)

}
