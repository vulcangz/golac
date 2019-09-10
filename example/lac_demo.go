package main

import (
	"fmt"
	"time"

	"github.com/vulcangz/golac"
)

func main() {
	text := "天气预报说今天要下雨"
	//text := "test.txt"

	// LocalExec connector is responsible to run PaddleHub LAC process.
	c := golac.NewLocalExec(nil)
	c.Option = "--input_text"
	//c.Option = "--input_file"

	// Annotate text
	doc, err := c.Run(text)
	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Printf("Cmd Reponse = %v\n", doc)

	start := time.Now()
	d, _ := golac.Decode(doc)
	elapsed := time.Now().Sub(start)
	fmt.Printf("After Decode: %v\nElapsed time: %v\n", d, elapsed)
	for _, v := range d {
		for i := 0; i < len(v.Word); i++ {
			fmt.Printf("%s(%s) ", v.Word[i], v.Tag[i])
		}
	}
}
