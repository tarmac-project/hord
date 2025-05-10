package bbolt

import (
	"fmt"
	"os"
	"testing"

	"github.com/tarmac-project/hord"
)

func BenchmarkDrivers(b *testing.B) {
	// Create some test data for Benchmarks
	data := []byte(`
  {
    "userId": 1,
    "id": 1,
    "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
    "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
  }
  `)

	b.Run("Bench_BoltDB", func(b *testing.B) {
		var db hord.Database
		var err error
		db, err = Dial(Config{
			Bucketname: "test",
			Filename:   "/tmp/bbolt-benchmark.db",
		})
		if err != nil {
			b.Fatalf("Got unexpected error when initializing hashmap - %s", err)
		}
		defer func() {
			if err := os.Remove("/tmp/bbolt-benchmark.db"); err != nil {
				b.Logf("Failed to remove benchmark db: %v", err)
			}
		}()
		defer db.Close()

		// Setup DB
		err = db.Setup()
		if err != nil {
			b.Fatalf("Unknown error setting up DB - %s", err)
		}

		// Execute HealthCheck
		err = db.HealthCheck()
		if err != nil {
			b.Fatalf("Error while checking health of DB - %s", err)
		}

		b.Run("SET", func(b *testing.B) {
			// Clean up Keys Created for Test
			b.Cleanup(func() {
				keys, _ := db.Keys()
				for _, d := range keys {
					_ = db.Delete(d)
				}
			})

			// Exec Benchmark
			for i := 0; i < b.N; i++ {
				err := db.Set("Test_Keys_"+fmt.Sprintf("%d", i), data)
				if err != nil {
					b.Fatalf("Error when executing Benchmark test - %s", err)
				}
			}
		})

		b.Run("GET", func(b *testing.B) {
			// Clean up Keys Created for Test
			b.Cleanup(func() {
				keys, _ := db.Keys()
				for _, d := range keys {
					_ = db.Delete(d)
				}
			})

			// Setup A Bunch of Keys
			b.StopTimer()
			for i := 0; i < 5000; i++ {
				_ = db.Set("Test_Keys_"+fmt.Sprintf("%d", i), data)
			}

			// Exec Benchmark
			count := 0
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				if count > 4999 {
					count = 0
				}
				_, err := db.Get("Test_Keys_" + fmt.Sprintf("%d", count))
				if err != nil {
					b.Fatalf("Error when executing Benchmark test - %s", err)
				}
			}
		})

	})
}
