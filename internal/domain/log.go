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
	UserID     string    `json:"user_id" form:"user_id"`
	ActionType string    `json:"action_type" form:"action_type"`
	UserRole   string    `json:"user_role" form:"user_role"`
	From       string    `json:"from" form:"from"`
	To         string    `json:"to" form:"to"`
	FromDate   time.Time `json:"from_date" swaggerignore:"true"`
	ToDate     time.Time `json:"to_date" swaggerignore:"true"`
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
