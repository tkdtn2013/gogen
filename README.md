# gogen

CLI tool to generate Go REST API projects using Gin.

Features:
- Automatically sets up Swagger documentation.
- Configures a MariaDB database using Gorm.

## Usage

```bash
gogen create myapp
```

The generated project will include Swagger and database setup. Provide the MariaDB connection string via the `DB_DSN` environment variable. If not set, a default DSN connecting to `localhost:3306` will be used.

