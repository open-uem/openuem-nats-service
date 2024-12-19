module github.com/doncicuto/openuem-nats-service

go 1.23.4

replace github.com/doncicuto/openuem_utils => ./internal/utils

require (
	github.com/doncicuto/openuem_utils v0.0.0-00010101000000-000000000000
	github.com/nats-io/nats-server/v2 v2.10.23
	golang.org/x/sys v0.28.0
)

require (
	github.com/danieljoos/wincred v1.2.2 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/minio/highwayhash v1.0.3 // indirect
	github.com/nats-io/jwt/v2 v2.5.8 // indirect
	github.com/nats-io/nkeys v0.4.8 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.30.0 // indirect
	golang.org/x/time v0.8.0 // indirect
)

require (
	github.com/stretchr/testify v1.10.0 // indirect
	gopkg.in/ini.v1 v1.67.0
)
