package conf

type app struct {
	Desc        string `yaml:"desc"`
	Version     string `yaml:"version"`
	Addr        string `yaml:"addr"`
	ConfigFile  string `yaml:"configFile"`
	LogsAddress string `yaml:"logsAddress"`
	ShortHost   string `yaml:"ShortHost"`
	DefaultUrl  string `yaml:"defaultUrl"`
}

type mysql struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbName"`
	MaxConn  int    `yaml:"maxConn"`
	MaxOpen  int    `yaml:"maxOpen"`
}

type redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type ViperConfig struct {
	App   app   `yaml:"app"`
	Mysql mysql `yaml:"mysql"`
	Redis redis `yaml:"redis"`
}
