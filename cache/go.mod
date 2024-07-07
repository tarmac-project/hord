module github.com/tarmac-project/hord/cache

go 1.20

require (
	github.com/tarmac-project/hord v0.4.0
	github.com/tarmac-project/hord/cache/lookaside v0.4.0
	github.com/tarmac-project/hord/drivers/cassandra v0.4.0
	github.com/tarmac-project/hord/drivers/hashmap v0.4.0
	github.com/tarmac-project/hord/drivers/mock v0.4.0
	github.com/tarmac-project/hord/drivers/redis v0.4.0
)

require (
	github.com/FZambia/sentinel v1.1.1 // indirect
	github.com/gocql/gocql v1.6.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gomodule/redigo v1.9.2 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/madflojo/hord v0.3.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
