package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/ukasyah99/user-api/app"
	_ "gitlab.com/ukasyah99/user-api/docs"
)

// @title User API
// @version 1.0
// @description A simple user API example.

// @contact.name Ukasyah
// @contact.email hi.ukasyah@gmail.com

// @license.name The Unlicense
// @license.url http://unlicense.org

func main() {
	app := &app.App{}
	app.Run()
}
