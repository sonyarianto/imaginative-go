db.auth('root', 'mysqlpassword')

db = db.getSiblingDB('go_db')

db.content_category.insertMany(
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
  }
)