package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)

type options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type book struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []options `json:"options"`
}

type bookList map[string]book

func main() {
	startTime := time.Now()
	fileName := flag.String("json", "books.json", "read specific json file")
	flag.Parse()
	bl := &bookList{}

	file, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, bl)

	// f, err := os.Open(*fileName)
	// if err != nil {
	// 	panic(err)
	// }
	// d := json.NewDecoder(f)
	// err = d.Decode(bl)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Printf("%v\n", bl)
	duration := time.Now().Sub(startTime)
	fmt.Println(duration)
}
