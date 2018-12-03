db.auth('root', 'mongodbpassword')

db.getSiblingDB('go_db')

db.sample_content.insertMany(
  [{
	"title": "Hello World",
	"slug": "hello-world",
	"short_description": "Hello World is the standard ritual for us when learning new programming language. It's good for your mind and soul hahaha!",
	"tags": [{"tag": "beginner"},{"tag": "hello world"}],
	"content_file": "hello-world.md"
   },
   {
	"title": "URL router a.k.a HTTP request multiplexer",
	"slug": "url-router-http-request-multiplexer",
	"short_description": "Create URL router for your web based application.",
	"tags": [{"tag": "beginner"},{"tag": "mux"},{"tag": "multiplexer"},{"tag": "http"}],
	"content_file": "url-router-http-request-multiplexer.md"
   },
   {
	"title": "Load a text file",
	"slug": "load-a-text-file",
	"short_description": "Load a text file is easy task but maybe useful for you.",
	"tags": [{"tag": "beginner"},{"tag": "io"},{"tag": "file"}],
	"content_file": "load-a-text-file.md"
   }]
)
