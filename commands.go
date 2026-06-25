package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if c.commandMap[cmd.name] == nil {
		return fmt.Errorf("The command, %v, does not exist.", cmd.name)
	}
	return c.commandMap[cmd.name](s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Error, login command requires username")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}
	
	fmt.Printf("User, %v, has been set.", s.cfg.UserName)

	return nil
}