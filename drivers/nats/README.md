# Hord NATS driver

The NATS driver is a separate Go module. The driver on `main` requires Go 1.26.0 or later and uses `github.com/nats-io/nats.go` v1.56.0. This Go minimum applies only to the NATS driver module, not to Hord core or the other drivers.

Projects that cannot yet move to Go 1.26 can continue to use the previous [NATS driver release, v0.8.1](https://github.com/tarmac-project/hord/releases/tag/drivers/nats/v0.8.1).
