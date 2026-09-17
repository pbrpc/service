//revive:disable:package-comments
package service

// Configuration for the server
type Configuration struct {
	Name    string `env:"SERVER_NAME"`
	Version string `env:"SERVER_VERSION" envDefault:"dev"`
}
