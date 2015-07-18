package models

type Command struct {
	Name string `json:"name"`
}

type Commands struct {
	Collection string    `json:"collection"`
	CommandList   []Command `json:"commands"`
}
