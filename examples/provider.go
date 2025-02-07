package main

import (
	"fmt"

	"github.com/Ilya-Guyduk/openinfra/parser"
)

func main() {
	spec, err := parser.ParseFile("./some_file.yaml")
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

	providerList := spec.GetProviderList()

	for _, provider := range providerList {
		fmt.Println(provider)
	}

}
