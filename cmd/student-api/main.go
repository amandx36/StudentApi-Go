package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/amandx36/studentCrudApiGo/internal/config"
	"github.com/amandx36/studentCrudApiGo/internal/http/handlers/student"

	"github.com/amandx36/studentCrudApiGo/internal/storage/sqlite"
)

func main() {

	// load the config
	cfg := config.MustLoad()
	// database set up
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Storage initilise dude ", slog.String("env", cfg.Env), slog.String("Version", "1.0.0"))
	// set up router
	router := http.NewServeMux()

// All api end points dude 



	// 1 
	router.HandleFunc("POST /api/v1/students", student.New(storage))
	// 2 
	router.HandleFunc("GET /api/v1/students/{id}", student.GetStudentById(storage))
	// 3 
	router.HandleFunc("GET /api/v1/getStudents",student.GetList(storage))
	// setup server dude
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("server is starting ")
	log.Println("Starting Student API...")
	log.Println("Server listening on", cfg.HTTPServer.Addr)
	slog.Info("Server started successfully ")

	// to shut down the server gracefully so that it complete the server gracefully run in another go routine dude

	// making a channel

	// to  storing and   closing signal into channel so to stop the server if it request dude

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Starting and handling the error
		error := server.ListenAndServe()

		if error != nil {
			log.Fatal("Failed to start the server dude")
		}

	}()

	<-done

	// now stoping the server
	slog.Info("shutting down the server ")

	// shutiing down the server in specific time
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("Failed to shut down the server ", slog.String("error", err.Error()))
	}
	slog.Info("Server shut down successfully dude ")
}
