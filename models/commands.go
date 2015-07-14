package models

type Command struct {
	Name string `json:"name"`
}

type Commands struct {
	Collection string    `json:"collection"`
	Commands   []Command `json:"commands"`
}
