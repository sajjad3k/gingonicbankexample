package main

import "github.com/sajjad3k/ginbankapiex/routes"

func main() {

	r := routes.Setroutes()

	r.Run()

}
