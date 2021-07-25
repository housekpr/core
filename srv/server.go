package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/housekpr/core/srv/devices/gpioctrl"
	"github.com/housekpr/core/srv/devices/pump"
	"github.com/housekpr/core/srv/graph"
	"github.com/housekpr/core/srv/graph/generated"
)

func main() {

	// Get the os parameters
	port := "8080"
	if osval := os.Getenv("HKPR_PORT"); osval != "" {
		port = osval
	}
	env := "DEV"
	if osval := os.Getenv("HKPR_ENV"); osval != "" {
		env = osval
	}
	log.Printf("Starting in '%s' environment.", env)

	//
	initDevices()
	// todo: move to https://pkg.go.dev/os/signal. No defer on ctrl+c...
	defer gpioctrl.ReleaseGPIO()

	// Serve the webapp
	http.Handle("/", http.FileServer(http.Dir("./webapp")))

	// Start GraphQL Server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	// Enable GraphQL Playground only in development scenarios
	if env == "DEV" {
		http.Handle("/play", playground.Handler("GraphQL playground", "/query"))
		log.Printf("GraphQL playground started at http://localhost:%s/play for GraphQL playground", port)
	}
	// Serve GraphQL at /query endpoint
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Make sure devices are in defined state before proceeding
func initDevices() {
	log.Println("Initialising the devices in order to start with defined state.")
	pump.Stop()
}
