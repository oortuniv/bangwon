package config

type Bangwon struct {
	Id     string `yaml:"id"`
	Server string `yaml:"server"`
}

type Config struct {
	Id       string    `yaml:"id"`
	Port     int       `yaml:"port"`
	Bangwons []Bangwon `yaml:"bangwons"`

	SrcPath string `yaml:"src_path"`
}
