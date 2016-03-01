// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	Rabbitbeat RabbitbeatConfig
}

type RabbitbeatConfig struct {
	Period string `yaml:"period"`
}
