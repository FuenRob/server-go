module github.com/fuenrob/server-go

go 1.19

require (
	gorm.io/driver/mysql v1.4.1
	gorm.io/gorm v1.24.0
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20221012134737-56aed061732a // indirect
	golang.org/x/net v0.0.0-20221012135044-0b7e1fb9d458 // indirect
	golang.org/x/sys v0.0.0-20221010170243-090e33056c14 // indirect
	golang.org/x/text v0.3.8 // indirect
)

replace github.com/fuenrob/server-go/utils => ./utils

replace github.com/fuenrob/server-go/configs => ./configs

replace github.com/fuenrob/server-go/models => ./models

replace github.com/fuenrob/server-go/controllers => ./controllers
