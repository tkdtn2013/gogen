# gogen

CLI tool to generate Go REST API projects using Gin.

Features:
- Automatically sets up Swagger documentation.
- Configures a MariaDB database using Gorm.
- Supports HTTPS when `SSL_CERT_FILE` and `SSL_KEY_FILE` are provided.

## Usage

```bash
gogen create myapp
```

The generated project will include Swagger and database setup. Provide the MariaDB connection string via the `DB_DSN` environment variable. If not set, a default DSN connecting to `localhost:3306` will be used.

To enable HTTPS during development, set `SSL_CERT_FILE` and `SSL_KEY_FILE` to the paths of your certificate and private key.

