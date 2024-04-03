module example.com/config

go 1.21.2

require (
	example.com/models v0.0.0-00010101000000-000000000000
	gorm.io/driver/mysql v1.5.1
	gorm.io/gorm v1.25.4
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.14.0 // indirect
)

replace example.com/models => ../models
