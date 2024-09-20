package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type cred struct {
	Env       map[string]string `json:"env"`
	ExpiresAt *time.Time        `json:"expiresAt"`
}

func main() {
	c := cred{
		Env: map[string]string{
			"key": "value",
		},
		ExpiresAt: ptr(time.Now().Add(5 * time.Second)),
	}
	cJSON, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("failed to marshal credential: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(cJSON))
}

func ptr[T any](t T) *T {
	return &t
}
