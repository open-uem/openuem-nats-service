module github.com/doncicuto/openuem-nats-service

go 1.23.2

replace github.com/doncicuto/openuem_ent => ./internal/models/ent

replace github.com/doncicuto/openuem_utils => ./internal/utils

require (
	github.com/doncicuto/openuem_ent v0.0.0-00010101000000-000000000000
	github.com/doncicuto/openuem_utils v0.0.0-00010101000000-000000000000
	github.com/jackc/pgx/v5 v5.7.1
	golang.org/x/sys v0.26.0
)

require (
	github.com/danieljoos/wincred v1.2.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
)

require (
	ariga.io/atlas v0.19.1-0.20240203083654-5948b60a8e43 // indirect
	entgo.io/ent v0.14.1
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	gopkg.in/ini.v1 v1.67.0
)
