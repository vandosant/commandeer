package models

import (
	"os/exec"
)

type commander interface {
	Command()
}

type Command struct {
	Name string `json:"name"`
}

func (c *Command) Command() {
	cmd := exec.Command("say", "hello")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

type Commands struct {
	Collection string    `json:"collection"`
	CommandList   []Command `json:"commands"`
}
