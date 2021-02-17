package main

import (
	"BackendGo/router"
	"BackendGo/server"
	"log"
)

func main() {
	app, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	if err = router.ApplyRoutes(app); err != nil {
		log.Fatal(err)
	}
	if err = app.Router.Run(); err != nil {
		log.Fatal(err)
	}
}
