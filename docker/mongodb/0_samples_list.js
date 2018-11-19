db.auth('root', 'mongodbpassword')

db.getSiblingDB('go_db')

db.code_samples_list.insertMany(
  [
	  {
	    "title": "The standard hello world ritual",
	    "slug": "hello-world",
	    "short_description": "This is a must ritual on any programming language.",
	    "result_available": true
	  }
  ]
)
