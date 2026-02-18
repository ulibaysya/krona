package config

type Log struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

type Service struct {
	TemplatesPath string `yaml:"templates_path"`
	Static        struct {
		Serve bool   `yaml:"serve"`
		Path  string `yaml:"path"`
	} `yaml:"static"`
	AdminPanel struct {
		Enabled bool `yaml:"enabled"`
		Server  `yaml:"server"`
	} `yaml:"admin_panel"`
}

type Server struct {
	Address string `yaml:"address"`
	Network string `yaml:"network"`
}

type Storage struct {
	Method string `yaml:"method"`
	Cache  Cache  `yaml:"cache"`
	RDBMS  RDBMS  `yaml:"rdbms"`
}

type RDBMS struct {
	Engine  string `yaml:"engine"`
	Connstr string `yaml:"connstr"`
}

type Cache struct {
	Engine  string `yaml:"engine"`
	Connstr string `yaml:"connstr"`
}
