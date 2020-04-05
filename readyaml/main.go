package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type pathurl struct {
	PATH string `ymal:"path"`
	URL  string `yaml:"url"`
}

func main() {
	fileName := flag.String("yaml", "config.yml", "the configuration of the project")
	flag.Parse()

	data, err := ioutil.ReadFile(*fileName)
	if err != nil {
		fmt.Println("cannot read yaml file", *fileName)
	}

	pu := pathurl{}
	err = yaml.Unmarshal(data, &pu)
	if err != nil {
		fmt.Println("cannot parse yaml file", *fileName)
	}
	fmt.Println(pu)
}
