package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "teslapid"
	app.Usage = "Program to manage automatic uploads from your Tesla Dashcam to the cloud"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "storage-dir",
			Value:  "TeslaUSB",
			EnvVar: "TESLAPI_STORAGE",
			Usage:  "Directory where recordings are stored",
		},
		cli.StringFlag{
			Name:   "username",
			Value:  "teslapi",
			EnvVar: "TESLAPI_USERNAME",
			Usage:  "Username used for logging into the web interface",
		},
		cli.StringFlag{
			Name:   "password",
			Value:  "Password1!",
			EnvVar: "TESLAPI_PASSWORD",
			Usage:  "Password used for logging into the web interface",
		},
		cli.StringFlag{
			Name:   "s3-key",
			Value:  "",
			EnvVar: "TESLAPI_S3_KEY",
			Usage:  "The AWS key for accessing the bucket",
		},
		cli.StringFlag{
			Name:   "s3-secret",
			Value:  "",
			EnvVar: "TESLAPI_S3_SECRET",
			Usage:  "The AWS secret for accessing the bucket",
		},
		cli.StringFlag{
			Name:   "s3-region",
			Value:  "us-east-1",
			EnvVar: "TESLAPI_S3_REGION",
			Usage:  "The region that you created your bucket in",
		},
		cli.StringFlag{
			Name:   "s3-endpoint",
			Value:  "",
			EnvVar: "TESLAPI_S3_ENDPOINT",
			Usage:  "If you are using a local S3 compatabile project like Minio or using DigitalOcean Spaces, you need to override the endpoint",
		},
	}

	// this is the default action to run
	app.Action = func(c *cli.Context) error {
		fmt.Println(os.Getenv("TESLAPI_PASSWORD"))
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "run the web server",
			Action: func(c *cli.Context) error {
				fmt.Println("TODO build the web server")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
