package main

import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/keppy/pour"
	"github.com/urfave/cli/v2"
	"github.com/valyala/fasthttp"
)

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))

}

func main() {
	var uri string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "uri",
				Value:       "http://localhost:8969",
				Aliases:     []string{"u"},
				Usage:       "URI for the POST",
				Destination: &uri,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "json",
				Aliases: []string{"j"},
				Usage:   "post json",
				Action: func(c *cli.Context) error {
					fmt.Println("json post: ", c.Args().First())

					req := fasthttp.AcquireRequest()

					pour.JSON(c.Args().First(), req.BodyWriter())

					req.Header.SetMethod("POST")
					req.Header.SetContentType("application/json")
					req.SetRequestURI(uri)

					res := fasthttp.AcquireResponse()

					if err := fasthttp.Do(req, res); err != nil {
						fmt.Println("err: ", err)
						panic("handle error")
					}
					fasthttp.ReleaseRequest(req)

					body := res.Body()
					fmt.Println("response body: ", b2s(body))

					fasthttp.ReleaseResponse(res) // Only when you are done with body!
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
