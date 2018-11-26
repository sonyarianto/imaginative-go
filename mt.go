package main

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
	"log"
)

type Tag struct {
	Tag string `bson:"tag"`
}

type Content struct {
	Id               string `bson:"id"`
	Title            string `bson:"title"`
	Slug             string `bson:"slug"`
	ShortDescription string `bson:"short_description"`
	ContentFile      string `bson:"content_file"`
	//Tags bsonx.Arr `json:"tags"`
	Tags []Tag `bson:"tags"`
}

func main() {
	// Prepare database.
	client, err := mongo.NewClient("mongodb://root:mongodbpassword@localhost:32771")
	if err != nil {
		log.Fatal(err)
	}

	// Connect to database.
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Select a database.
	db := client.Database("go_db")

	// Do the query to a collection on database.
	c, err := db.Collection("sample_content").Find(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close(context.TODO())

	rowsData := make([]Content, 0)

	doc := bsonx.Doc{}
	// Start looping on the query result.
	for c.Next(context.TODO()) {
		doc = doc[:0]
		err := c.Decode(&doc)

		if err != nil {
			log.Fatal(err)
		}

		//title, err := doc.LookupErr("title")
		id := doc.Lookup("_id").ObjectID().Hex()
		title := doc.Lookup("title").StringValue()
		shortDescription := doc.Lookup("short_description").StringValue()
		slug := doc.Lookup("slug").StringValue()
		contentFile := doc.Lookup("content_file")
		
		//tags := doc.Lookup("tags").Array()
		tags := doc.Lookup("tags")

		log.Println(tags)

		arr := tags.Array()

		for _, val := range arr {
			//require.Equal(t, bson.TypeEmbeddedDocument, val.Type())
				subdoc := val.Document()

				//require.Equal(t, 1, len(subdoc))
				tag := subdoc.Lookup("tag")

				log.Println(tag)
//require.NoError(t, err)
		}

		//haha := [...]Tag{Tag: "beginner"}

		queryResult := Content{
			Id:               id,
			Title:            title,
			ShortDescription: shortDescription,
			Slug:             slug,
			ContentFile:      contentFile,
			//Tags:             tags,
			Tags:             []Tag{Tag{Tag: "beginner"}},
		}
		// elem := bson.Doc{}

		// if err = c.Decode(elem); err != nil {
		// 	log.Fatal(err)
		// }

		// queryResult := Content{
		// 	ID:          elem.Lookup("_id").ObjectID().Hex(),
		// 	Title:   elem.Lookup("title").StringValue(),
		// 	Slug:  elem.Lookup("slug").StringValue(),
		// 	ShortDescription:        elem.Lookup("short_description").StringValue(),
		// 	ContentFile: elem.Lookup("content_file").StringValue(),
		// 	//Tags: elem.Lookup("tags"),
		// }

		log.Println(queryResult)

		rowsData = append(rowsData, queryResult)
	}
}
