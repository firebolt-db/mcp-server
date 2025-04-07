package main

import (
	"context"
	"log/slog"
	"strings"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/gocolly/colly/v2"
)

type DocumentationArticle struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func Scrape(ctx context.Context, logger *slog.Logger) <-chan DocumentationArticle {

	const maxWorkers = 5

	c := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(2),
	)
	c.Context = ctx

	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: maxWorkers,
	})
	if err != nil {
		logger.Error("Failed to set limit", "error", err)
		return nil
	}

	results := make(chan DocumentationArticle, maxWorkers)

	c.OnXML("//url/loc/text()", func(e *colly.XMLElement) {
		if err := e.Request.Visit(e.Text); err != nil {
			logger.Warn("Failed to visit URL", "url", e.Text, "error", err)
		}
	})

	c.OnHTML("div#main-content main", func(e *colly.HTMLElement) {

		clone := e.DOM.Clone()
		clone.Find("div.query-window").Remove()
		content, err := clone.Html()
		if err != nil {
			logger.Warn("Failed to find content", "error", err)
			return
		}
		content, err = htmltomarkdown.ConvertString(
			strings.TrimSpace(content),
		)
		if err != nil {
			logger.Warn("Failed to convert content to markdown", "error", err)
			return
		}

		results <- DocumentationArticle{
			ID:      generateID(e.Request.URL.String()),
			Content: content,
		}
	})

	if err := c.Visit("https://docs.firebolt.io/sitemap.xml"); err != nil {
		logger.Error("Failed to visit sitemap", "error", err)
	}

	go func() {
		c.Wait()
		close(results)
	}()

	c.OnError(func(response *colly.Response, err error) {
		logger.Warn("Request failed", "url", response.Request.URL.String(), "error", err)
	})

	return results
}

func generateID(url string) string {

	id := strings.TrimPrefix(url, "https://docs.firebolt.io/")
	id = strings.TrimSuffix(id, ".html")
	id = strings.TrimSuffix(id, "/")
	id = strings.ReplaceAll(id, "/", "_")
	id = strings.ReplaceAll(id, "-", "_")
	id = strings.ToLower(id)

	if id == "" {
		return "index"
	}

	return id
}
