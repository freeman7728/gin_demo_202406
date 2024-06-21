package config

type Logger struct {
	Level          string `yaml:"level"`
	Prefix         string `yaml:"prefix"`
	Director       string `yaml:"director"`
	Show_line      bool   `yaml:"show_line"`
	Log_in_console bool   `yaml:"log_in_console"`
}
