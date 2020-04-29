package main

import (
	blogInit "blog/init"
	"blog/routers"
)

func main() {
	blogInit.Init()

	r := routers.NewRouter()
	r.Run(":8000")
}
