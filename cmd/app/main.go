package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gwaysys/goapp/module/db"
	"github.com/gwaysys/goapp/version"
)

func main() {
	fmt.Println("git commit:", version.GitCommit)
	mdb := db.GetCache("master")
	defer db.CloseCache()
	_ = mdb

	fmt.Println("[ctrl+c to exit]")
	end := make(chan os.Signal, 2)
	signal.Notify(end, os.Interrupt, os.Kill)
	<-end
}
