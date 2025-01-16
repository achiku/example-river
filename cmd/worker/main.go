package main

import (
	"context"
	"fmt"
	"log"
	"os"

	rivertest "github.com/achiku/example-river"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
)

func main() {
	workers := river.NewWorkers()
	if err := river.AddWorkerSafely(workers, &rivertest.SortWorker{}); err != nil {
		fmt.Printf("Error adding worker: %v\n", err)
		return
	}

	ctx := context.Background()
	dbPool, err := pgxpool.New(ctx, os.Getenv("postgres://river@localhost:5432/river"))
	if err != nil {
		log.Fatal("pgxpool.New failed")
	}

	riverClient, err := river.NewClient(riverpgxv5.New(dbPool), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: 100},
		},
		Workers: workers,
	})
	if err != nil {
		log.Fatal("river.NewClient failed")
	}

	// Run the client inline. All executed jobs will inherit from ctx:
	if err := riverClient.Start(ctx); err != nil {
		// handle error
	}
	log.Println("started")
	<-ctx.Done()
}
