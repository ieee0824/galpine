package main

import (
	"encoding/json"
	"fmt"

	"github.com/ieee0824/galpine"
)

func main() {
	d := galpine.NewDatas()
	bin, _ := json.MarshalIndent(d, "", "    ")
	fmt.Println(string(bin))
}
