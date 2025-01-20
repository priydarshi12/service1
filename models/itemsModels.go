package models
import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name string `json:"name"`	
	Type string `json:"type"`
	User string `json:"email"`	
}

