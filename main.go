package main

import (
	"fmt"
	"log"
	"os"

	"github.com/keppy/pour"
	"github.com/urfave/cli/v2"
	"github.com/valyala/fasthttp"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "post",
				Aliases: []string{"p"},
				Usage:   "post form to endpoint",
				Action: func(c *cli.Context) error {
					fmt.Println("posting data: ", c.Args().First())
					fmt.Println("to endpoint: ", c.Args().Tail())
					return nil
				},
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "get a http endpoint",
				Action: func(c *cli.Context) error {
					fmt.Println("get: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "json",
				Aliases: []string{"j"},
				Usage:   "options for json",
				Subcommands: []*cli.Command{
					{
						Name:  "post",
						Usage: "POST JSON to endpoint",
						Action: func(c *cli.Context) error {
							fmt.Println("json post: ", c.Args().First())

							req := fasthttp.AcquireRequest()

							pour.JSON(c.Args().First(), req.BodyWriter())

							req.Header.SetMethod("POST")
							req.Header.SetContentType("application/json")
							req.SetRequestURI("http://localhost:3000")

							res := fasthttp.AcquireResponse()

							if err := fasthttp.Do(req, res); err != nil {
								panic("handle error")

							}
							fasthttp.ReleaseRequest(req)

							body := res.Body()
							fmt.Println("response body: ", body)

							fasthttp.ReleaseResponse(res) // Only when you are done with body!
							return nil
						},
					},
					{
						Name:  "get",
						Usage: "GET JSON from endpoint",
						Action: func(c *cli.Context) error {
							fmt.Println("json get: ", c.Args().First())
							return nil
						},
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
