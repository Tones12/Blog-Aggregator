package main

import (
	"context"
	"fmt"
	"github.com/tones12/blog-aggregator/internal/database"
	"strconv"

)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	limit := 2
	if len(cmd.Args) == 1 {
		i, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("Error: %w, optional limit argument must be a number in quotes", err)
		}
		limit = i
	}


	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID:	user.ID,
		Limit:	int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}
	fmt.Println("==================================================================")
	
	fmt.Printf("\nPrinting %d posts for %s:\n\n", limit, user.Name)
	fmt.Println("==================================================================")
	for _, post := range posts {
		fmt.Printf("\n%s\n", post.Url)
		fmt.Printf("\nTitle: %s\n", post.Title.String)
		fmt.Printf("\nDescription:\n%s\n\n", post.Description.String)
		fmt.Println("==================================================================")
	}
	return nil
}
