package books

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Authors     []string           `bson:"authors" json:"authors"`
	PublishDate string             `bson:"publish_date" json:"publish_date"`
}
