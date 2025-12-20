package config

type HTTPServer struct{
	Addr string
}

//Used the Clean Env package for configuration

type Config struct{
	Env         string `yaml:"env" env-required:"true" env-deafult:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

