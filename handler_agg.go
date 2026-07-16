package main

import (
	"context"
	"fmt"
	"time"
	"github.com/google/uuid"
	"database/sql"
	"github.com/tones12/blog-aggregator/internal/database"
	"github.com/lib/pq"
	"errors"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time between posts>", cmd.Name)
	}

	time_between_reqs := cmd.Args[0]
	timeBetweenRequests, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("error in duration between requests: %w", err)
	}
	ticker := time.NewTicker(timeBetweenRequests)
	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)
	for ; ; <-ticker.C {
		fmt.Println("New feed request in progress")
		err = scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("error scraping feed %w", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed fetched: %w", err)
	}

	feedRSS, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error collecting RSS information from feed: %w", err)
	}
	
	var pqErr *pq.Error

	for _, item := range feedRSS.Channel.Item {
		
		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			t, err = time.Parse(time.RFC1123, item.PubDate)
		}

		publishedAt := sql.NullTime{}
		if err == nil {
			publishedAt = sql.NullTime{
			Time:  t,
			Valid: true,
			}
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:				uuid.New(),
			CreatedAt:		time.Now().UTC(),
			UpdatedAt:		time.Now().UTC(),
			Title:			sql.NullString{
				String:		item.Title,
				Valid:		item.Title != "",
			},
			Url:			item.Link,
			Description:	sql.NullString{
				String:		item.Description,
				Valid:		item.Description != "",
			},
			PublishedAt:	publishedAt,
			FeedID:			feed.ID,
		})
		if err != nil {
			if errors.As(err, &pqErr) {
				if pqErr.Code == "23505" {
					continue
				}
			}
			return fmt.Errorf("error creating post: %w", err)
		}
	}

	return nil
}