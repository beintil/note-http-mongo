package note

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateNoteDTD struct {
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	CreateAt    time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdateAt    time.Time `bson:"update_at,omitempty" json:"update_at,omitempty"`
}

type AllNoteDTD struct {
	UUID        primitive.ObjectID `bson:"uuid,omitempty" json:"uuid,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	CreateAt    time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdateAt    time.Time          `bson:"update_at,omitempty" json:"update_at,omitempty"`
}

type UpdateNoteDTD struct {
	UUID        primitive.ObjectID `bson:"uuid" json:"uuid"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	UpdateAt    time.Time          `bson:"update_at,omitempty" json:"update_at,omitempty"`
}

type DeleteNoteDTD struct {
	UUID        primitive.ObjectID `bson:"uuid,omitempty" json:"uuid,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	CreateAt    time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdateAt    time.Time          `bson:"update_at,omitempty" json:"update_at,omitempty"`
}
