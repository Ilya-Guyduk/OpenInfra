package examples

import (
	"fmt"

	"github.com/Ilya-Guyduk/openinfra/parser"
)

func main() {
	spec := parser.ParseFile("./some_file.yaml")

	providerList := spec.GetProviderList()

	fmt.Printf(providerList)
}
