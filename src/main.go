package main

import (
	"cloudflare-dns-record-updater/app"
	"cloudflare-dns-record-updater/data"
	"fmt"
)

func main() {
	data, filled, err := data.NewData()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !filled {
		fmt.Println("Please fill data. Path:", data.Path)
		return
	}
	err = app.Execute(data.Data)
	if err != nil {
		fmt.Println(err.Error())
	}
}
