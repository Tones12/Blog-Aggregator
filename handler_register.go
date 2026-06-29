package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
	"context"

	"github.com/tones12/blog-aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt:	time.Now().UTC(),
		Name:		name,
	})
	if err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}
	
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}

	fmt.Println("User successfully registered.")

	fmt.Printf("User %v Created:\nID: %v\nCreatedAt: %v\nUpdatedAt: %v\nName: %v\n", user.Name, user.ID, user.CreatedAt, user.UpdatedAt, user.Name)

	return nil
}
