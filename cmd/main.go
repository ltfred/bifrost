package main

import (
	"bifrost/commands"
	"bifrost/internal/app"
)

func init()  {
	app.InitDatabase()
}


func main()  {
	commands.Execute()
}
