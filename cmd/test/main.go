package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ieee0824/galpine"
)

func main() {
	f, err := os.Open("./data.js")
	if err != nil {
		panic(err)
	}
	d := galpine.NewDatas(f)
	bin, _ := json.MarshalIndent(d, "", "    ")
	fmt.Println(string(bin))
}
