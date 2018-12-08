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
	"tags": [{"tag": "beginner"},{"tag": "mux"},{"tag": "multiplexer"},{"tag": "http"},{"tag": "web"}],
	"content_file": "url-router-http-request-multiplexer.md"
   },
   {
	"title": "Simple template and passing data to template",
	"slug": "simple-template-and-passing-data-to-template",
	"short_description": "How to parse simple template and passing data to a template.",
	"tags": [{"tag": "beginner"},{"tag": "template"},{"tag": "web"}],
	"content_file": "template.md"
   },
   {
	"title": "Load a text file",
	"slug": "load-a-text-file",
	"short_description": "Load a text file is easy task but maybe useful for you.",
	"tags": [{"tag": "beginner"},{"tag": "io"},{"tag": "file"}],
	"content_file": "load-a-text-file.md"
   },
   {
	"title": "Get local IP address",
	"slug": "get-local-ip-address",
	"short_description": "Get local IP address.",
	"tags": [{"tag": "beginner"},{"tag": "network"},{"tag": "internet"}],
	"content_file": "get-local-ip-address.md"
   }]
)
