package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	b, err := ioutil.ReadFile("Input")
	if err != nil {
		log.Fatal(err)
	}
	data := string(b)
	capcha := 0
	for index, strnum := range data {
		if index == len(data)-1 {
			continue
		}
		if int(strnum) == int(data[index+1]) {
			i, err := strconv.Atoi(string(strnum))
			if err != nil {
				log.Fatal(err)
			}
			capcha += i
		}
	}
	if data[0] == data[len(data)-1] {
		i, err := strconv.Atoi(string(data[0]))
		if err != nil {
			log.Fatal(err)
		}
		capcha += i
	}

	fmt.Printf("Capcha: %v\n", capcha)
}
