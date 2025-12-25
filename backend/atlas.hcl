data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/models",
    "--dialect", "postgres"
  ]
}

env "local" {
  src = data.external_schema.gorm.url
  url = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
  dev = "docker://postgres/15/dev?search_path=public"
  migration {
    dir = "file://migrations"
    format = atlas
  }
}

env "dev" {
  src = data.external_schema.gorm.url
  url = getenv("DB_BASE_URL")
  dev = getenv("DB_BASE_URL")
  migration {
    dir = "file://migrations"
    format = atlas
  }
}