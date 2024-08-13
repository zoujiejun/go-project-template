package config

type Config struct {
	App       App        `yaml:"app"`
	Server    Server     `yaml:"server"`
	Databases []Database `yaml:"databases"`
}

type App struct {
	Name   string `yaml:"name"`
	ApiKey string `yaml:"api_key"`
}

type Server struct {
	Mode string `yaml:"mode"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Database struct {
	Name         string `yaml:"name"`
	Driver       string `yaml:"name"`
	Dsn          string `yaml:"dsn"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Database     string `yaml:"database"`
	Location     string `yaml:"location"`
	Charset      string `yaml:"charset"`
	ParseTime    bool   `yaml:"parse_time"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}
