module github.com/kliuchnikovv/engi-example

go 1.22

toolchain go1.23.1

require github.com/kliuchnikovv/engi v0.0.0-20240131115705-a839d2f14f4d

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	golang.org/x/text v0.14.0 // indirect
)

require (
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gorilla/mux v1.8.1
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.55.0 // indirect
	go.opentelemetry.io/otel v1.30.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.30.0 // indirect
	gorm.io/driver/sqlite v1.5.6
	gorm.io/gorm v1.25.12
)

replace github.com/kliuchnikovv/engi => ../engi
