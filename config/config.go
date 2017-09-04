package config

import (
	"github.com/naoina/toml"
	"os"
)

func LoadConfig() *GlobalConfig {
	dir, _ := os.Getwd()
	file := dir + "/config/prod.toml"

	io, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer io.Close()

	var c GlobalConfig
	if err := toml.NewDecoder(io).Decode(&c); err != nil {
		panic(err)
	}

	return &c
}

type GlobalConfig struct {
	Title string
	Owner struct {
		Name string
	}
	Database DbOptions
	Ipfs IpfsOptions
}

type IpfsOptions struct {
	Server string
	Port int
	Topic string
}

type DbOptions struct {
	Server string
	Port int
	Name string
	CacheCapacity int
	BufferCapacity int
}

func defaultConfig() {

}