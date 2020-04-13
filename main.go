package main

import "gin-todo/router"

func main() {
	r:=router.NewRouter()
	r.Run(":8080")
}
