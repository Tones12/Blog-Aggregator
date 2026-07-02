package main

import (
	"fmt"
	"context"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	
	err := s.db.DeleteUsers(context.Background())
		if err != nil {
		return fmt.Errorf("error, could not delete users: %w", err)
	}

	fmt.Println("Users successfully deleted!")
	return nil
}