module myproject

go 1.17

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/stretchr/testify v1.7.0
	github.com/yougov/go-sqlite3 v1.0.0
	github.com/lib/pq v1.10.0
	github.com/dgrijalva/jwt-go v4.0.0
	github.com/getsentry/sentry-go v0.12.0
	github.com/Shopify/sarama v0.0.0-20220330181343-8b9a8f9d3f3d
	github.com/Shopify/sarama/v2 v2.0.0-20220330181343-8b9a8f9d3f3d
)

replace (
	github.com/Shopify/sarama => github.com/Shopify/sarama/v2
)

replace (
	github.com/Shopify/sarama/v2 => github.com/Shopify/sarama/v2 v2.0.0-20220330181343-8b9a8f9d3f3d
)

go.mod is not a Go source file.