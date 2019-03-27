package configration

type Configuration struct {
	Version string `yaml:"version"`
	Storage map[string]interface{}
}
