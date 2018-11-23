db.auth('root', 'mongodbpassword')

db.getSiblingDB('go_db')

db.sample_content.insertMany(
  [
	  {
	    title: "Hello World",
	    slug: "hello-world",
	    short_description: "Hello World is the standard ritual for us when learning new programming language. It's good for you mind and soul hahaha!",
	    tags: ["beginner", "hello world"],
	    content_file: "hello-world.md"
	  }
  ]
)
