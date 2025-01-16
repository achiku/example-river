package main

import (
	"context"
	"fmt"
	"log"

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
	dbPool, err := pgxpool.New(ctx, "postgres://river@localhost:5432/river")
	if err != nil {
		log.Fatalf("pgxpool.New failed: %s", err)
	}

	riverClient, err := river.NewClient(riverpgxv5.New(dbPool), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: 100},
		},
		Workers: workers,
	})
	if err != nil {
		log.Fatalf("river.NewClient failed: %s", err)
	}

	// Run the client inline. All executed jobs will inherit from ctx:
	if err := riverClient.Start(ctx); err != nil {
		log.Fatalf("riverClient.Start failed: %s", err)
	}
	log.Println("started")
	<-ctx.Done()
}
