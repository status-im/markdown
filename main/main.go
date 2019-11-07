package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomarkdown/markdown"
)

func main() {
	md := []byte("## markdown document")
	output := markdown.Parse(md, nil)
	fmt.Printf("## markdown document\n")
	fmt.Printf("%+v\n", output)
	fmt.Printf("%+v\n", output.GetChildren()[0])
	fmt.Printf("%+v\n", output.GetChildren()[0].GetChildren()[0])
	j, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(j)
}
