package main

import "bifrost/internal/router"

func main()  {
	r := router.Routers()
	r.Run()
}
