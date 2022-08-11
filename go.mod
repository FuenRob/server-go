module github.com/fuenrob/server-go

go 1.19

require (
	github.com/labstack/echo/v4 v4.7.2
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20220809184613-07c6da5e1ced // indirect
	golang.org/x/sys v0.0.0-20220808155132-1c4a2a72c664 // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace github.com/fuenrob/server-go/utils => ./utils
replace github.com/fuenrob/server-go/config => ./config
replace github.com/fuenrob/server-go/model => ./model
