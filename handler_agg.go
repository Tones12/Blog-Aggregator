package main

import (
	"fmt"
	"context"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	ctx := context.TODO()
	
	tempURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(ctx, tempURL)
	if err != nil {
		return err
	}
	
	fmt.Println(feed)
	return nil
}