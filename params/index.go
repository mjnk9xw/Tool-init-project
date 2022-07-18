package params

type TypeProject string
type ModelProject string
type DB string
type TypeConfig string

const (
	APIGIN  TypeProject = "api-gin"
	APIEcho TypeProject = "api-echo"
	APIMux  TypeProject = "api-mux"
	APIHttp TypeProject = "api-http"
	Grpc    TypeProject = "grpc"
	Thrift  TypeProject = "thrift"

	Th3Layer  ModelProject = "3layer"
	CleanCode ModelProject = "clean-code"
	Mvc       ModelProject = "mvc"

	Redis         DB = "redis"
	Mongo         DB = "mongo"
	Mysql         DB = "mysql"
	Postgres      DB = "postgres"
	Oracle        DB = "oracle"
	Elasticsearch DB = "elasticsearch"

	Json TypeConfig = "json"
	Yaml TypeConfig = "yaml"
	Env  TypeConfig = "env"
)
