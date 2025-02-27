package config

type Mysql struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	Db        string `yaml:"db"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Log_level string `yaml:"log_level"`
	Config    string `yaml:"config"`
}

func (m *Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Db + "?" + m.Config
}
