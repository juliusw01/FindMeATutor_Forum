package MongoDB

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Thread struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Titel     string             `json:"Titel" bson:"Titel"`
	Date      primitive.DateTime `json:"Date" bson:"Date"`
	User      string             `json:"User" bson:"User"`
	Frage     string             `json:"Frage" bson:"Frage"`
	Antworten []Antwort          `json:"Antworten" bson:"Antworten"`
}

type Antwort struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Date   primitive.DateTime `json:"Date" bson:"Date"`
	User   string             `json:"User" bson:"User"`
	Inhalt string             `json:"Inhalt" bson:"Inhalt"`
}
