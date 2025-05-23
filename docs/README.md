# Hord

![Hord Gopher](hord-banner.png)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tarmac-project/hord)
[![codecov](https://codecov.io/gh/tarmac-project/hord/branch/main/graph/badge.svg?token=0TTTEWHLVN)](https://codecov.io/gh/tarmac-project/hord)
[![Go Report Card](https://goreportcard.com/badge/github.com/tarmac-project/hord)](https://goreportcard.com/report/github.com/tarmac-project/hord)
[![Go Reference](https://pkg.go.dev/badge/github.com/tarmac-project/hord.svg)](https://pkg.go.dev/github.com/tarmac-project/hord)
[![License](https://img.shields.io/github/license/tarmac-project/hord)](https://choosealicense.com/licenses/apache-2.0/)

Package hord provides a simple and extensible interface for interacting with various database systems in a uniform way.

Hord is designed to be a database-agnostic library that provides a common interface for interacting with different database systems. It allows developers to write code that is decoupled from the underlying database technology, making it easier to switch between databases or support multiple databases in the same application.

## Features

- **Driver-based**: Hord follows a driver-based architecture, where each database system is implemented as a separate driver. This allows for easy extensibility to support new databases.
- **Uniform API**: Hord provides a common API for database operations, including key-value operations, setup, and configuration. The API is designed to be simple and intuitive.
- **Pluggable**: Developers can choose and configure the desired database driver based on their specific needs.
- **Error handling**: Hord provides error types and constants for consistent error handling across drivers.
- **Testing with Mock Driver**: Hord provides a mock driver in the `mock` package, which can be used for testing purposes. The `mock` driver allows users to define custom functions executed when calling the `Database` interface methods, making it easier to test code that relies on the Hord interface.
- **Documentation**: Each driver comes with its own package documentation, providing guidance on how to use and configure the driver.

### Evolving Features

- **Cache Implementations**: Combine database drivers with pre-defined cache implementations.

## Database Drivers

| Database | Support | Docs | Comments | Protocol Compatible Alternatives |
| -------- | ------- | ---- | -------- | -------------------------------- |
| [BoltDB](https://github.com/etcd-io/bbolt) | ✅ | [![Go Reference](https://pkg.go.dev/badge/github.com/tarmac-project/hord/drivers/bbolt.svg)](https://pkg.go.dev/github.com/tarmac-project/hord/drivers/bbolt) | | |
| [Cassandra](https://cassandra.apache.org/) | ✅ | [![Go Reference](https://pkg.go.dev/badge/github.com/tarmac-project/hord/drivers/cassandra.svg)](https://pkg.go.dev/github.com/tarmac-project/hord/drivers/cassandra) | | [ScyllaDB](https://www.scylladb.com/), [YugabyteDB](https://www.yugabyte.com/), [Azure Cosmos DB](https://learn.microsoft.com/en-us/azure/cosmos-db/introduction) |
| Hashmap | ✅ | [![Go Reference](https://pkg.go.dev/badge/github.com/tarmac-project/hord/drivers/hashmap.svg)](https://pkg.go.dev/github.com/tarmac-project/hord/drivers/hashmap) | In-memory, Optional storage to YAML or JSON file ||
| [Mock](https://pkg.go.dev/github.com/tarmac-project/hord/mock) | ✅ | [![Go Reference](https://pkg.go.dev/badge/github.com/tarmac-project/hord/drivers/mock)](https://pkg.go.dev/github.com/tarmac-project/hord/drivers/mock) | Mock Database interactions within unit tests ||
| [NATS](https://nats.io/) | ✅ | [![Go Reference](https://pkg.go.dev/badge/github.com/tarmac-project/hord/drivers/nats)](https://pkg.go.dev/github.com/tarmac-project/hord/drivers/nats) | Experimental ||
| [Redis](https://redis.io/) | ✅ | [![Go Reference](https://pkg.go.dev/badge/github.com/tarmac-project/hord/drivers/redis)](https://pkg.go.dev/github.com/tarmac-project/hord/drivers/redis) || [Dragonfly](https://www.dragonflydb.io/), [KeyDB](https://docs.keydb.dev/) |

## Caching Implementations

[![Go Reference](https://pkg.go.dev/badge/github.com/tarmac-project/hord/cache)](https://pkg.go.dev/github.com/tarmac-project/hord/cache)

Hord provides a set of cache implementations that can be used with the Hord interface to allow combining datastores. For example, you can use the Look Aside cache driver to check Redis before fetching from Cassandra.

| Cache Strategy | Docs | Comments |
| -------------- | ---- | -------- |
| Look Aside | [![Go Reference](https://pkg.go.dev/badge/github.com/tarmac-project/hord/cache/lookaside)](https://pkg.go.dev/github.com/tarmac-project/hord/cache/lookaside) | Cache is checked before database, if not found in cache, database is checked and cache is updated |

## Usage

To use Hord, import it as follows:

    import "github.com/tarmac-project/hord"

### Creating a Database Client

To create a database client, you need to import and use the appropriate driver package along with the `hord` package.

For example, to use the Redis driver:

```go
import (
    "github.com/tarmac-project/hord"
    "github.com/tarmac-project/hord/drivers/redis"
)

func main() {
    var db hord.Database
    db, err := redis.Dial(redis.Config{})
    if err != nil {
        // Handle connection error
    }

    // Use the db client for database operations
    // ...
}
```

Each driver provides its own `Dial` function to establish a connection to the database. Refer to the specific driver documentation for more details.

### Database Operations

Once you have a database client, you can use it to perform various database operations. The API is consistent across different drivers.

```go
// Set a value
err = db.Set("key", []byte("value"))
if err != nil {
    // Handle error
}

// Retrieve a value
value, err := db.Get("key")
if err != nil {
    // Handle error
}
```

Refer to the `hord.Database` interface documentation for a complete list of available methods.

## Contributing
Thank you for your interest in helping develop Hord. The time, skills, and perspectives you contribute to this project are valued.

Please reference our [Contributing Guide](CONTRIBUTING.md) for details.

## License
[Apache License 2.0](https://choosealicense.com/licenses/apache-2.0/)
