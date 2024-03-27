package main

import (
	"context"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/takumi616/english-example-sentence-api/config"
)

func run(ctx context.Context) error {
	//Get environment variables
	cfg, err := config.New()
	if err != nil {
		log.Println("Failed to get config.")
		return err
	}

	//Create http listener with port received from config file
	listener, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Printf("Failed to create http listener with port %s", cfg.Port)
		return err
	}

	url := fmt.Sprintf("http://%s", listener.Addr().String())
	log.Printf("URL: %v", url)

	//Get routing info
	mux := setUpRouting(ctx, cfg)
	if err != nil {
		return err
	}

	//Start http server
	return runServer(ctx, listener, mux)
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("HTTP server did not work correctly: %v", err)
	}
	// config, err := config.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// dataSourceName := "host=" + config.DBHost + " port=" + config.DBPort + " user=" + config.DBUser + " password=" + config.DBPassword + " dbname=" + config.DBName + " sslmode=" + config.DBSSLMODE
	// _, err = sql.Open("postgres", dataSourceName)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connected to postgresql successfully.")
}
