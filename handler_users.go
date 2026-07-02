package main

import (
	"fmt"
	"context"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	
	users, err := s.db.GetUsers(context.Background())
		if err != nil {
		return fmt.Errorf("error user does not exist: %w", err)
	}
	fmt.Println("Users:")
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}