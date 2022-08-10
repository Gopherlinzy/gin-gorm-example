module github.com/Gopherlinzy/go-gin-example

go 1.18

require (
	github.com/astaxie/beego v1.12.3
	github.com/gin-gonic/gin v1.8.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.8.4
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.66.6
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.6 // indirect
	github.com/go-openapi/swag v0.22.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/goccy/go-json v0.9.10 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.2 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20220809184613-07c6da5e1ced // indirect
	golang.org/x/sys v0.0.0-20220808155132-1c4a2a72c664 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/Gopherlinzy/go-gin-example/conf => ./goWorkspace/go-gin-example/pkg/conf
	github.com/Gopherlinzy/go-gin-example/docs => ./goWorkspace/go-gin-example/pkg/docs
	github.com/Gopherlinzy/go-gin-example/middleware => ./goWorkspace/go-gin-example/middleware
	github.com/Gopherlinzy/go-gin-example/middleware/jwt => ./goWorkspace/go-gin-example/middleware/jwt
	github.com/Gopherlinzy/go-gin-example/models => ./goWorkspace/go-gin-example/models
	github.com/Gopherlinzy/go-gin-example/pkg/e => ./goWorkspace/go-gin-example/pkg/e
	github.com/Gopherlinzy/go-gin-example/pkg/logging => ./goWorkspace/go-gin-example/pkg/logging
	github.com/Gopherlinzy/go-gin-example/pkg/setting => ./goWorkspace/go-gin-example/pkg/setting
	github.com/Gopherlinzy/go-gin-example/pkg/util => ./goWorkspace/go-gin-example/pkg/util
	github.com/Gopherlinzy/go-gin-example/routers => ./goWorkspace/go-gin-example/routers
	github.com/Gopherlinzy/go-gin-example/routers/api => ./goWorkspace/go-gin-example/routers/api
	github.com/Gopherlinzy/go-gin-example/routers/api/vi => ./goWorkspace/go-gin-example/routers/api/vi
)
