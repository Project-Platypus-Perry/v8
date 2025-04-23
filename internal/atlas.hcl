data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/model",
    "--dialect", "postgres",
  ]
}

data "composite_schema" "app" {
  # Load enum types first.
  schema "public" {
    url = "file://internal/schema.sql"
  }
  # Then, load the GORM models.
  schema "public" {
    url = data.external_schema.gorm.url
  }
}

env "gorm" {
  src = data.composite_schema.app.url
  url = "postgres://postgres:postgres@localhost:5434/postgres?sslmode=disable"
  dev = "postgres://postgres:postgres@localhost:5434/postgres_dev?sslmode=disable"
  migration {
    dir = "file://internal/db/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
