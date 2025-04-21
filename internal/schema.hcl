schema "public" {}

enum "language" {
  schema = schema.public
  values = ["en", "hi"]
}

enum "visibility" {
	schema = schema.public 
	values = ["private", "public"]
}

enum "content_type" {
	schema = "schema.public 
	values = ["notes", "dpp", "video"]
}