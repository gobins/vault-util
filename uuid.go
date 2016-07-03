package main

import "github.com/twinj/uuid"

func getUUID() string {
	uuid.Init()
	id := uuid.NewV1()

	return id.String()
}
