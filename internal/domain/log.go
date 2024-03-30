package domain

import "time"

type User struct {
	ID       string                 `json:"id" bson:"user_id,omitempty" binding:"required"`
	FullName string                 `json:"full_name" bson:"full_name,omitempty" binding:"required"`
	Role     string                 `json:"role" bson:"role,omitempty" binding:"required"`
	Details  map[string]interface{} `json:"details" bson:"details,omitempty"`
}

type Action struct {
	Type    string                 `json:"type" bson:"type,omitempty" binding:"required"`
	Model   string                 `json:"model" bson:"model,omitempty" binding:"required"`
	Details map[string]interface{} `json:"details" bson:"details,omitempty"`
}
type InsertLogPayload struct {
	User      User      `json:"user" bson:"user" binding:"required"`
	Action    Action    `json:"action" bson:"action" binding:"required"`
	Message   string    `json:"message" bson:"message,omitmepty" binding:"required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" swaggerignore:"true"`
}

type RetreiveLogsFilter struct {
	BasicFilter
	Query map[string]interface{} `json:"query"`
}

type Log struct {
	ID        string    `json:"id" bson:"_id"`
	User      User      `json:"user" bson:"user"`
	Action    Action    `json:"action" bson:"action"`
	Message   string    `json:"message" bson:"message,omitmepty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type RetreiveLogsResponse struct {
	Pagination
	Logs []Log `json:"logs"`
}
