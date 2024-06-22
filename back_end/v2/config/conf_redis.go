package config

type Redis struct {
	Addr         string `yaml:"Addr"`
	Password     string `yaml:"Password"`
	DB           int    `yaml:"DB"`
	PoolSize     int    `yaml:"PoolSize"`
	MinIdleConns int    `yaml:"MinIdleConns"`
}
