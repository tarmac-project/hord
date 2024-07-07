module github.com/tarmac-project/hord/cache

go 1.20

require (
	github.com/tarmac-project/hord v0.5.0
	github.com/tarmac-project/hord/cache/lookaside v0.5.0
	github.com/tarmac-project/hord/drivers/cassandra v0.5.0
	github.com/tarmac-project/hord/drivers/hashmap v0.5.0
	github.com/tarmac-project/hord/drivers/mock v0.5.0
	github.com/tarmac-project/hord/drivers/redis v0.5.0
)

require (
	github.com/FZambia/sentinel v1.1.1 // indirect
	github.com/gocql/gocql v1.6.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gomodule/redigo v1.9.2 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
