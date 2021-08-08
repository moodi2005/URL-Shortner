package databaseconfig

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

type databaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

// read database config yaml file and save in databaseconfig
func readConfigFile(fileName string, databaseconfig *databaseConfig) error {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error in Read Config yaml file: %s", err)
	}

	err = yaml.Unmarshal(yamlFile, &databaseconfig)
	if err != nil {
		return fmt.Errorf("error in Unmarshal Config yaml Content: %s", err)
	}

	return nil
}

// get database config info from readConfigFile function
// and genrate config info string
func getConfigInfo() (string, error) {
	var databaseconfig databaseConfig
	err := readConfigFile("../database/config.yaml", &databaseconfig)
	if err != nil {
		return "", err
	}
	configInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		databaseconfig.Host, databaseconfig.Port, databaseconfig.User, databaseconfig.Password, databaseconfig.DBName, databaseconfig.SSLMode)

	return configInfo, nil
}

func ConnectToDB() (*sql.DB, error) {
	// get config info
	configInfo, err := getConfigInfo()
	if err != nil {
		return nil, err
	}
	// connect to postgres
	DB, err := sql.Open("postgres", configInfo)
	if err != nil {
		return nil, fmt.Errorf("error in connect to database %s", err)
	}
	
	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	return DB, nil
}
