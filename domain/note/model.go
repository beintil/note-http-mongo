package note

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	UUID        primitive.ObjectID `bson:"uuid,omitempty" json:"uuid,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	CreateAt    time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdateAt    time.Time          `bson:"update_at,omitempty" json:"update_at,omitempty"`
}
