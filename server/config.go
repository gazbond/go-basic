package server

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/spf13/viper"
)

// Config structure
type Config struct {
	Server struct {
		Host string
		Port int
	}
}

// Config variables
var (
	ServerUrl string
)

// Load config
func LoadConfig(name string) {

	// Set config path, name, type
	viper.AddConfigPath("config")
	viper.SetConfigName(name)
	viper.SetConfigType("json")

	// Read config file
	err := viper.ReadInConfig()
	if err != nil {
		color.Red.Println("Error reading config")
		panic(err)
	}

	// Parse config
	c := &Config{}
	err = viper.Unmarshal(c)
	if err != nil {
		color.Red.Println("Error parsing config")
		panic(err)
	}

	// Set config variables
	ServerUrl = fmt.Sprintf("%v:%v", c.Server.Host, c.Server.Port)

	color.Yellow.Println("Config loaded")
}
