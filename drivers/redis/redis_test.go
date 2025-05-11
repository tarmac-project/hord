package redis

import (
	"crypto/tls"
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/tarmac-project/hord"
	"testing"
	"time"
)

func TestHealthCheck(t *testing.T) {
	t.Run("No connection", func(t *testing.T) {
		db := &Database{}
		err := db.HealthCheck()
		if !errors.Is(err, hord.ErrNoDial) {
			t.Errorf("Expected ErrNoDial, got %v", err)
		}
	})

	t.Run("Failed health check", func(t *testing.T) {
		// Create a database instance with nil connection but non-nil pool
		// to simulate a health check failure
		db := &Database{
			pool: &redis.Pool{
				// TestOnBorrow will fail because there's no real connection
				Dial: func() (redis.Conn, error) {
					return nil, errors.New("connection error")
				},
			},
		}

		err := db.HealthCheck()
		if err == nil {
			t.Fatal("Expected health check to fail, got nil error")
		}

		// Verify that ErrHealthCheckFailure is in the error chain
		if !errors.Is(err, hord.ErrHealthCheckFailure) {
			t.Errorf("Expected error to contain ErrHealthCheckFailure, got %v", err)
		}
	})
}

func TestConnectivity(t *testing.T) {
	t.Run("No Config", func(t *testing.T) {
		_, err := Dial(Config{})
		if err == nil {
			t.Errorf("Expected error when Dialing with no config set, got nil")
		}
	})

	t.Run("Just Redis", func(t *testing.T) {
		db, err := Dial(Config{
			ConnectTimeout: time.Duration(5) * time.Second,
			Server:         "redis:6379",
		})
		if err != nil {
			t.Fatalf("Failed to connect to Redis - %s", err)
		}
		defer db.Close()

		// Test a connection
		c := db.pool.Get()
		defer c.Close() // nolint:errcheck

		_, err = c.Do("PING")
		if err != nil {
			t.Errorf("Failed to ping Redis server - %s", err)
		}
	})

	t.Run("Fake TLS", func(_ *testing.T) {
		_, _ = Dial(Config{
			ConnectTimeout: time.Duration(5) * time.Second,
			Server:         "redis:6379",
			TLSConfig:      &tls.Config{},
		})
	})

	t.Run("Sentinel Connection No Master", func(t *testing.T) {
		db, err := Dial(Config{
			ConnectTimeout: time.Duration(5) * time.Second,
			SentinelConfig: SentinelConfig{
				Servers: []string{"redis-sentinel:26379"},
			},
		})
		if err == nil {
			defer db.Close()
			t.Fatalf("Failed to connect to Redis via Sentinel - %s", err)
		}
	})

	t.Run("Sentinel Connection", func(t *testing.T) {
		db, err := Dial(Config{
			ConnectTimeout: time.Duration(5) * time.Second,
			SentinelConfig: SentinelConfig{
				Servers: []string{"redis-sentinel:26379"},
				Master:  "mymaster",
			},
		})
		if err != nil {
			t.Fatalf("Failed to connect to Redis via Sentinel - %s", err)
		}
		defer db.Close()

		// Test a connection
		c := db.pool.Get()
		defer c.Close() // nolint:errcheck

		_, err = c.Do("PING")
		if err != nil {
			t.Errorf("Failed to ping Redis server - %s", err)
		}

		// Check TestOnBorrow
		err = db.pool.TestOnBorrow(c, time.Now())
		if err != nil {
			t.Errorf("Error returned when testing pool connection - %s", err)
		}
	})
}
