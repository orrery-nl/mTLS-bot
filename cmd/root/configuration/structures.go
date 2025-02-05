package configuration

type RootConfiguration struct {
	// CLIs - list of CLI clients.
	CLIs []CliClientConfiguration `yaml:"clis"`
}

type CliClientConfiguration struct {
	Id        string `yaml:"id"`
	State     string `yaml:"state"`
	PublicKey string `yaml:"public_key"`
}
