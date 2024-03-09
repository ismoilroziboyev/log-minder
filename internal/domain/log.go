package domain

import "time"

type LogStructure struct {
	ID        string                 `json:"id"`
	User      map[string]interface{} `json:"user"`
	Action    map[string]interface{} `json:"action"`
	CreatedAt time.Time              `json:"created_at"`
}

type InsertLogPayload struct {
	User      map[string]interface{} `json:"user" bson:"user"`
	Action    map[string]interface{} `json:"action" bson:"action"`
	CreatedAt time.Time              `json:"created_at" bson:"created_at"`
}
