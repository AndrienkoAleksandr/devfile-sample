# Results API Server

## Variables

| Environment Variable     | Description                                                                                                                       | Example                                      |
|--------------------------|-----------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------|
| DB_USER                  | Postgres Database user                                                                                                            | user                                         |
| DB_PASSWORD              | Postgres Database Password                                                                                                        | hunter2                                      |
| DB_HOST                  | Postgres Database host                                                                                                            | /cloudsql/my-project:us-east1:tekton-results |
| DB_NAME                  | Postgres Database name                                                                                                            | tekton_results                               |
| DB_SSLMODE               | Database SSL mode                                                                                                                 | verify-full                                  |
| DB_ENABLE_AUTO_MIGRATION | Auto-migrate the database on startup (create/update schemas). For further details, refer to <https://gorm.io/docs/migration.html> | true (default)                               |
| GRPC_PORT                | gRPC Server Port                                                                                                                  | 50051 (default)                              |
| REST_PORT                | REST proxy Port                                                                                                                   | 8080  (default)                              |
| PROMETHEUS_PORT          | Prometheus Port                                                                                                                   | 9090  (default)                              |
| TLS_HOSTNAME_OVERRIDE    | Override the hostname used to serve TLS. This should not be set (or set to the empty string) in production environments.          | results.tekton.dev                           |
| TLS_PATH                 | Path to TLS files                                                                                                                 | /etc/tls                                     |
| NO_AUTH                  | Disable RBAC check for resources                                                                                                  | false (default)                              |
| LOG_LEVEL                | Log level for api server                                                                                                          | info (default)                               |
| LOGS_API                 | Enable logs storage service                                                                                                       | false (default)                              |
| LOGS_TYPE                | Determine Logs storage backend type                                                                                               | File (default)                               |
| LOGS_BUFFER_SIZE         | Buffer for streaming logs                                                                                                         | 32768 (default)                              |
| LOGS_PATH                | Logs storage path                                                                                                                 | logs (default)                               |
| S3_BUCKET_NAME           | S3 Bucket name                                                                                                                    | <S3 Bucket Name>                             |
| S3_ENDPOINT              | S3 Endpoint                                                                                                                       | https://s3.ap-south-1.amazonaws.com          |
| S3_HOSTNAME_IMMUTABLE    | S3 Hostname immutable                                                                                                             | false (default)                              |
| S3_REGION                | S3 Region                                                                                                                         | ap-south-1                                   |
| S3_ACCESS_KEY_ID         | S3 Access Key ID                                                                                                                  | <S3 Acces Key>                               |
| S3_SECRET_ACCESS_KEY     | S3 Secret Access Key                                                                                                              | <S3 Access Secret>                           |
| S3_MULTI_PART_SIZE       | S3 Multi part size                                                                                                                | 5242880 (default)                            |

These values can also be set in the config file located in the `config/env/config` directory.

Values derived from Postgres DSN

If you use the default postgres database we provide, the `DB_HOST` can be set as `tekton-results-postgres-service.tekton-pipelines`.