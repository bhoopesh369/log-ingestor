package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Level      string             `bson:"level"`
	Message    string             `bson:"message"`
	ResourceID string             `bson:"resourceId"`
	Timestamp  string             `bson:"timestamp"`
	TraceID    string             `bson:"traceId"`
	SpanID     string             `bson:"spanId"`
	Commit     string             `bson:"commit"`
	Metadata   Metadata           `bson:"metadata"`
}

// Metadata represents additional metadata for the log
type Metadata struct {
	ParentResourceID string `bson:"parentResourceId"`
}

func LogCollectionName() string {
	return "logs"
}
