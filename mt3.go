package main

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

type Tag struct {
	Tag string `bson:"tag"`
}

type Content struct {
	Id               objectid.ObjectID `bson:"_id" json:"_id"`
	Title            string            `bson:"title" json:"title"`
	Slug             string            `bson:"slug" json:"slug"`
	ShortDescription string            `bson:"short_description" json:"short_description"`
	ContentFile      string            `bson:"content_file" json:"content_file"`
	Tags             []Tag             `bson:"tags" json:"tags"`
}

func main() {
	// Prepare database.
	client, err := mongo.NewClient("mongodb://root:mongodbpassword@localhost:32771")
	if err != nil {
		log.Fatal(err)
	}

	// Connect to database.
	err = client.Connect(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Select a database.
	db := client.Database("db-imaginative-go")

	// Do the query to a collection on database.
	c, err := db.Collection("sample_content").Find(nil, bson.D{{"slug", "hello-world"}})
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close(nil)

	var content []Content

	// Start looping on the query result.
	for c.Next(nil) {
		eachContent := Content{}

		err := c.Decode(&eachContent)
		if err != nil {
			log.Fatal(err)
		}

		content = append(content, eachContent)

		log.Println(eachContent)
	}
}
