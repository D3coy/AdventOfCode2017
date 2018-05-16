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
	size := len(data)
	for index, strnum := range data {
		compl := (index + int(size/2)) % size
		if int(strnum) == int(data[compl]) {
			i, err := strconv.Atoi(string(strnum))
			if err != nil {
				log.Fatal(err)
			}
			capcha += i
		}
	}

	fmt.Printf("Capcha: %v\n", capcha)
}
