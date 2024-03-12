package main

import (
	"github.com/shenqingawei/framework/app"
)

func main() {
	err := app.Init("mysql", "auto_migrate_user")
	if err != nil {
		panic(err)
	}
}
