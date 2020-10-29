package database

const (
	MongoConnectURI     = "mongodb://localhost:27017"
	MongoDatabase       = "test"
	MongoConnectTimeOut = 20
	MongoMaxPoolSize    = 20
)

const (
	RedisConnectAddr = "localhost:6379"
	RedisPassword    = ""
	RedisDatabase    = 0
)

const (
	MySQLConnectURI   = "root:@tcp(127.0.0.1:3306)/test"
	MySQLMaxOpenConns = 20
	MySQLMaxIdleConns = 20
)
