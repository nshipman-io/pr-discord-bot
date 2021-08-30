package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token string
	BotPrefix string
	GitToken string
	Repos []RepoStruct
	Config ConfigStruct
)

type ConfigStruct struct {
	Token string `json:"Token"`
	BotPrefix string `json:"BotPrefix""`
	GitToken string `json:"GitToken"`
	Repos []RepoStruct
}

type RepoStruct struct {
	Name string `json:"Name"`
	Owner string `json:"Owner"`
}

func ReadConfig() error{
	fmt.Println("Reading config file")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &Config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = Config.Token
	BotPrefix = Config.BotPrefix
	GitToken = Config.GitToken
	Repos = Config.Repos
	fmt.Println(Repos)

	return nil
}