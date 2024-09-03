package config_test

type User struct {
	Name string `json:"name,omitempty" xml:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	Age  int    `json:"age,omitempty" xml:"age,omitempty" yaml:"age,omitempty" toml:"age,omitempty"`
}
