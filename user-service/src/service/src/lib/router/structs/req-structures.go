package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type RequestFilter struct {
	Query *primitive.M `json:"query"`
	Skip int64 `form:"skip"`
	Limit int64 `form:"limit"`
	ID *string `json:"id"`
}
