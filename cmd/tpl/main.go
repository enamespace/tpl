package main

import (
	"log"
	"os"
	"strings"

	"github.com/enamespace/tpl/pkg/app"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

func readConfig(ctx *cli.Context) error {
	err := readFile(ctx)

	if err != nil {
		log.Fatal(err)
	}

	readSetString(ctx)
	log.Printf("use configuration files is: %s\n", viper.ConfigFileUsed())
	return nil
}

func readFile(ctx *cli.Context) error {
	file := ctx.String("file")
	if file == "" {
		viper.SetConfigFile("./values.yaml")
	} else {
		viper.SetConfigFile(file)
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func readSetString(ctx *cli.Context) error {
	kvs := ctx.StringSlice("set-string")

	for _, kv := range kvs {
		kvp := strings.Split(kv, ",")
		if len(kvp) != 2 {
			log.Printf("%s is wrong format", kv)
			continue
		}
		viper.Set(kvp[0], kvp[1])
	}

	return nil
}

func main() {
	app := &cli.App{

		Name: "template",
		Authors: []*cli.Author{
			{
				Name:  "enamespace",
				Email: "enamespace@gmail.com",
			},
		},
		Usage: "template for other golang projects",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file, f",
				Usage: "load configuration from `FILE`",
			},
			&cli.StringSliceFlag{
				Name:  "set-string",
				Usage: "override configuration by `KEY1=VALUE1[,KEY2=VALUE2]`",
			},
		},
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:  "template",
				Usage: "show app configuration rendered result",
				Action: func(ctx *cli.Context) error {
					err := readConfig(ctx)
					if err != nil {
						log.Fatal(err)
					}

					// 获取所有配置信息
					allSettings := viper.AllSettings()

					// 打印所有配置信息
					log.Println("All Settings:")
					for key, value := range allSettings {
						log.Printf("%s: %v\n", key, value)
					}

					return nil
				},
			},
			{
				Name:  "run",
				Usage: "start one app instance",
				Action: func(ctx *cli.Context) error {
					err := readConfig(ctx)
					if err != nil {
						log.Fatal(err)
					}

					go func() {
						app.NewApp()
					}()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
