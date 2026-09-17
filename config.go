//revive:disable:package-comments
package service

// Configuration for the service
type Configuration struct {
	Name    string `env:"SERVICE_NAME"`
	Version string `env:"SERVICE_VERSION" envDefault:"dev"`
	Address string `env:"SERVICE_ADDRESS" envDefault:"50051"`
}
