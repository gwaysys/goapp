package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gwaysys/goapp/cmd"
	"github.com/gwaysys/goapp/cmd/web/model/log"
	_ "github.com/gwaysys/goapp/cmd/web/route"
	"github.com/urfave/cli/v2"
)

func init() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
}

type FilterHandler struct {
}

func (h *FilterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func(start time.Time) {
		log.Printf("%s %s\n", log.ColorForMethod(r.Method), r.URL.String())
	}(time.Now())

	http.DefaultServeMux.ServeHTTP(w, r)
}

var server *cmd.App

func init() {
	server = &cmd.App{}
	server.Register("server", &cli.Command{
		Name:  "run",
		Usage: "run the server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "addr",
				Value: ":8081",
				Usage: "listen address",
			},
		},
		Action: func(ctx *cli.Context) error {
			// 过虑器
			filter := &FilterHandler{}
			addr := ctx.String("addr")
			log.Printf("Listen: %s\n", addr)
			log.Fatal(http.ListenAndServe(addr, filter))
			return nil
		},
	})
	server.Register("client", &cli.Command{
		Name:  "check",
		Usage: "check the server url",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "url",
				Value: "http://127.0.0.1:8081/hacheck",
				Usage: "the server alive address",
			},
			&cli.StringFlag{
				Name:  "method",
				Value: "HEAD",
				Usage: "http method",
			},
			&cli.StringFlag{
				Name:  "values",
				Value: "",
				Usage: "form encode values",
			},
		},
		Action: func(ctx *cli.Context) error {
			req, err := http.NewRequest(ctx.String("method"), ctx.String("url"), bytes.NewBuffer([]byte(ctx.String("values"))))
			if err != nil {
				return err
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			fmt.Printf("status:%d,body:%s\n", resp.StatusCode, string(body))
			return nil
		},
	})
}
func main() {
	server.Setup()
	server.RunAndExitOnError()
}
