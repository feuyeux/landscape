package cli

import (
	"fmt"
	"github.com/feuyeux/landscape/src/common"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Run(redisClient common.RedisClient) {
	app := cli.NewApp()
	app.Name = "landscape"
	app.Usage = "A simple redis client cli"
	app.Version = "v0.0.1"
	app.Authors = []*cli.Author{{
		Name:  "Lu Han/LiuWeng",
		Email: "feuyeux@gmail.com",
	},
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "config, c",
			Value:    "",
			Usage:    "redis config file",
			FilePath: "/opt/landscape/conf",
		},
	}

	app.Before = func(c *cli.Context) error {
		conf := c.String("config")
		lines, err := common.StringToLines(conf)
		if err != nil {
			log.Fatal(err)
		}
		redisClient.Open(lines[0], lines[1], lines[2])
		return nil
	}
	app.After = func(c *cli.Context) error {
		redisClient.Close()
		return nil
	}

	app.Commands = []*cli.Command{
		{
			Name:    "save",
			Aliases: []string{"w"},
			Usage:   "write kv to redis",
			Action: func(c *cli.Context) error {
				key := c.Args().First()
				value := c.Args().Get(1)
				result, _ := redisClient.SaveString(key, value)
				fmt.Println(result)
				return nil
			},
		},
		{
			Name:    "read",
			Aliases: []string{"r"},
			Usage:   "read kv from redis",
			Action: func(c *cli.Context) error {
				key := c.Args().First()
				value := redisClient.ReadString(key)
				pretty := common.JsonPretty(value)
				fmt.Println(pretty)
				return nil
			},
		},
		{
			Name:    "queue",
			Aliases: []string{"q"},
			Usage:   "queue commands",
			Subcommands: []*cli.Command{
				{
					Name:  "push",
					Usage: "push kv to queue",
					Action: func(c *cli.Context) error {
						key := c.Args().First()
						value := c.Args().Get(1)
						result, _ := redisClient.PushToQueue(key, value)
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "pop",
					Usage: "pop kv from queue",
					Action: func(c *cli.Context) error {
						key := c.Args().First()
						result, _ := redisClient.PopFromQueue(key)
						pretty := common.JsonPretty(result)
						fmt.Println(pretty)
						return nil
					},
				},
				{
					Name:  "all",
					Usage: "get all kv from queue",
					Action: func(c *cli.Context) error {
						key := c.Args().First()
						result, _ := redisClient.GetAllFromQueue(key)
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "last",
					Usage: "get last one kv from queue",
					Action: func(c *cli.Context) error {
						key := c.Args().First()
						result, _ := redisClient.GetLastOne(key)
						pretty := common.JsonPretty(result)
						fmt.Println(pretty)
						return nil
					},
				},
			},
		},
		{
			Name:    "map",
			Aliases: []string{"q"},
			Usage:   "map commands",
			Subcommands: []*cli.Command{
				{
					Name:  "save",
					Usage: "save kv to map",
					Action: func(c *cli.Context) error {
						key := c.Args().First()
						field := c.Args().Get(1)
						value := c.Args().Get(2)
						result, _ := redisClient.SaveMapValue(key, field, value)
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "read",
					Usage: "get kv from map",
					Action: func(c *cli.Context) error {
						key := c.Args().First()
						field := c.Args().Get(1)
						result := redisClient.GetMapValue(key, field)
						pretty := common.JsonPretty(result)
						fmt.Println(pretty)
						return nil
					},
				},
				{
					Name:  "all",
					Usage: "get all kv from map",
					Action: func(c *cli.Context) error {
						key := c.Args().First()
						result, _ := redisClient.GetMap(key)
						fmt.Println(result)
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
