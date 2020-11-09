package main

import (
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
)

const kVersion = "0.0.1"

func PrintVersion() {
	fmt.Printf("Version %s\n", kVersion)
}

func main() {
	var serverAddr string

	var taskYaml string

	app := &cli.App{
		Name:  "kbctl",
		Usage: "make an explosive entrance",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "s",
				Value:       "137.189.97.26:76107",
				Usage:       "The address of the server",
				Destination: &serverAddr,
			},
			&cli.StringFlag{
				Name:        "f",
				Value:       "./task.yaml",
				Usage:       "The task config file",
				Destination: &taskYaml,
			},
		},
		Action: func(c *cli.Context) error {
			PrintVersion()
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Check the version of tbctl",
				Action: func(c *cli.Context) error {
					PrintVersion()
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Submit a task to Server",
				Action: func(c *cli.Context) error {
					err := NewTaskRequest(serverAddr, taskYaml)
					return err
				},
			},
			{
				Name:    "check",
				Aliases: []string{"c"},
				Usage:   "Check the task execution status",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "Show the node status in testbed network",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
