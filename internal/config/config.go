package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Db *Db `json:"db"`
}

type Db struct {
	DBType   string `json:"db_type"`
	DBName   string `json:"db_name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func InitConfig(p string) *Config {
	var Conf = &Config{
		Db: &Db{},
	}

	f, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(Conf); err != nil {
		log.Fatal(err)
	}

	if err := check(Conf); err != nil {
		log.Fatal(err)
	}

	log.Println("Config loaded")
	return Conf
}

func check(c *Config) error {
	// todo:检查字段完整
	return nil
}
