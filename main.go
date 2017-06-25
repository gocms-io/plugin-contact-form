package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocms-io/gocms/controllers"
	"github.com/gocms-io/plugin-contact-form/context"
	"github.com/gocms-io/plugin-contact-form/database"
	"github.com/gocms-io/plugin-contact-form/repositories"
	"github.com/gocms-io/plugin-contact-form/services"
	_ "github.com/joho/godotenv/autoload"
)

//go:generate apidoc -c ./ -i ./models -i ./controllers/ -o ./docs/ -f ".*\\.go$" -f ".*\\.js$"
func main() {

	port := flag.Int("port", 30001, "port to run on.")
	flag.Parse()

	// setup context.
	context.Init()

	// setup database
	db := database.Default()

	// migrate cms db
	db.MigrateSql()

	// start gin with defaults
	r := gin.Default()
	// setup repositories

	rg := repositories.DefaultRepositoriesGroup(db)

	// setup services
	sg := services.DefaultServicesGroup(rg)

	// setup controllers
	controllers.DefaultControllerGroup(r, sg)

	fmt.Printf("Listening on port: %v\n", *port)

	r.Run(fmt.Sprintf("localhost:%v", *port))

}
