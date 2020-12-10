package config

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type App struct {
	WorkDir string `yaml:"workDir"`
	Oss     struct {
		AccessKey       string `yaml:"accessKey"`
		AccessKeySecret string `yaml:"accessKeySecret"`
		Endpoint        string `yaml:"endpoint"`
		Bucket          string `yaml:"bucket"`
		Domain          string `yaml:"domain"`
	} `yaml:"oss"`

	Db struct {
		Default Db
	}

	Redis struct {
		Default Redis
	}
}

type Db struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}
