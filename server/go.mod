module server

go 1.16

require (
	github.com/fsnotify/fsnotify v1.5.4
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.8.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/google/uuid v1.1.2
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/songzhibin97/gkit v1.2.4
	github.com/spf13/viper v1.12.0
	go.uber.org/zap v1.17.0
	golang.org/x/sync v0.0.0-20220513210516-0976fa681c29
	gorm.io/driver/mysql v1.3.4
	gorm.io/gorm v1.23.5
)

require (
	github.com/jonboulle/clockwork v0.3.0 // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible
)
