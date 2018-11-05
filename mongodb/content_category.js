db.auth('root', 'mysqlpassword')

db.getSiblingDB('go_db').createUser({user:'root', pwd:'mongodbpassword', roles:[{role:'readWrite', db:'go_db'}]})

db.content_category.insertMany(
  [
	  {
	    "name": "Rust",
	    "slug": "rust",
	    "short_description": "Resource for learning Rust language"
	  },
	  {
	    "name": "Go",
	    "slug": "go",
	    "short_description": "Resource for learning Go language"
	  },
	  {
	    "name": "PHP",
	    "slug": "php",
	    "short_description": "Resource for learning PHP language"
	  },
	  {
	    "name": "Python",
	    "slug": "phyton",
	    "short_description": "Resource for learning Phyton language"
	  },
	  {
	    "name": "Node.js",
	    "slug": "nodejs",
	    "short_description": "Resource for learning Node.js"
	  },
	  {
	    "name": "React",
	    "slug": "react",
	    "short_description": "Resource for learning React"
	  },
	  {
	    "name": "Erlang",
	    "slug": "erlang",
	    "short_description": "Resource for learning Erlang language"
	  },
	  {
	    "name": ".NET Core",
	    "slug": "dotnetcore",
	    "short_description": "Resource for learning .NET Core"
	  },
	  {
	    "name": "Ruby",
	    "slug": "ruby",
	    "short_description": "Resource for learning Ruby language"
	  },
	  {
	    "name": "Haskell",
	    "slug": "haskell",
	    "short_description": "Resource for learning Haskell language"
	  },
	  {
	    "name": "Scala",
	    "slug": "scala",
	    "short_description": "Resource for learning Scala language"
	  },
	  {
	    "name": "Symfony",
	    "slug": "symfony",
	    "short_description": "Resource for learning Symfony framework"
	  }
  ]
)
