package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	if err := run(ctx); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func run(ctx context.Context) error {

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	articles := Scrape(ctx, logger)

	for article := range articles {
		err := os.WriteFile(
			fmt.Sprintf("./fireboltdocs/%s.md", article.ID),
			[]byte(article.Content),
			0644,
		)
		if err != nil {
			logger.Error("Failed to write article", "id", article.ID, "error", err)
		}
	}
	return nil
}
