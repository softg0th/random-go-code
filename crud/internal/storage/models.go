package storage

import (
	bsonPrimitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type ObjectID = bsonPrimitive.ObjectID

type PostDocument struct {
	Name        string
	Description string
	Pic         string
}

type FieldSearch struct {
	Value string
	Field string
}

type FieldUpdate struct {
	ID        ObjectID
	NewValues map[string]string
}

type StrictFieldSearch struct {
	ID ObjectID
}
