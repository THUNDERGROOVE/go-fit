// config is to generate unique config files to encrypt session data
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

type config struct {
	SessionSecret string
}

var Config config

func init() {
	LoadConfig()
}

func LoadConfig() {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		fmt.Println("No config file found.  Generating SessionSecret")
		secret := randSeq(25)
		Config = config{
			SessionSecret: secret,
		}
		if data, err := json.Marshal(Config); err == nil {
			ioutil.WriteFile("config.json", data, 0777)
		} else {
			fmt.Println(err.Error())
		}
	} else {
		if data, err := ioutil.ReadFile("config.json"); err == nil {
			if err := json.Unmarshal(data, &Config); err == nil {
				fmt.Println("Config loaded from file.")
			} else {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
