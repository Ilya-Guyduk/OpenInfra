# go-openinfra

`go-openinfra` is a Go library for managing infrastructure components and resources, based on the OpenInfra specification. It allows you to define, parse, and generate YAML configurations for infrastructure resources and components, including cloud providers, hypervisors, virtual machines, networks, and more.

## Features

- Parse OpenInfra YAML specifications.
- Generate YAML configurations from Go structs.
- Support for defining providers (e.g., cloud, hypervisor) and components (e.g., virtual machines, networks).
- Define dependencies between components.
- Flexible and extensible design to add new resource types and providers.

## Installation

To install the `go-openinfra` library, you can use `go get`:

```bash
go get github.com/Ilya-Guyduk/go-openinfra
```

## Usage

### Example Specification

Here is an example of an OpenInfra specification that defines providers and components:

```yaml
openinfra: 1.0.0
info:
  title: OpenInfra Specification
  description: A specification for describing infrastructure resources and components.
  version: 1.0.0
  contact:
    name: Your Name
    email: your.email@example.com
  license:
    name: Apache 2.0
    url: https://www.apache.org/licenses/LICENSE-2.0

providers:
  - name: virtualbox
    type: hypervisor
    connection_details:
      address: 192.168.1.10
      username: admin
      password: password

  - name: cloud_provider
    type: cloud
    connection_details:
      api_endpoint: https://api.cloudprovider.com
      api_key: your_api_key_here

components:
  - type: virtual_machine
    name: local_vm
    provider: virtualbox
    properties:
      cpu: 2
      memory: 4GB
      disk_size: 50GB
      os: ubuntu-22.04
      network: local_network
    actions:
      - start
      - stop
      - restart

  - type: network
    name: local_network
    provider: cloud_provider
    properties:
      cidr: 192.168.1.0/24
      gateway: 192.168.1.1
      dns_servers:
        - 8.8.8.8
        - 8.8.4.4

dependencies:
  - component: local_vm
    depends_on:
      - local_network
```

### Parsing a Specification

To parse a YAML specification, you can use the following code:

```go
package main

import (
    "fmt"
    "log"
    "github.com/Ilya-Guyduk/go-openinfra/internal/parser"
)

func main() {
    specData := []byte(`
    openinfra: 1.0.0
    providers:
      - name: virtualbox
        type: hypervisor
        connection_details:
          address: 192.168.1.10
          username: admin
          password: password
    components:
      - type: virtual_machine
        name: local_vm
        provider: virtualbox
        properties:
          cpu: 2
          memory: 4GB
          disk_size: 50GB
          os: ubuntu-22.04
    `)

    spec, err := parser.ParseSpec(specData)
    if err != nil {
        log.Fatalf("Error parsing specification: %v", err)
    }

    fmt.Println("Parsed Specification:")
    fmt.Printf("%+v\n", spec)
}
```

### Generating YAML from Go Structs

To generate YAML from Go structs, you can use the `ToYAML` function:

```go
package main

import (
    "fmt"
    "log"
    "github.com/Ilya-Guyduk/go-openinfra/internal/parser"
    "github.com/Ilya-Guyduk/go-openinfra/infra"
)

func main() {
    vm := infra.VirtualMachine{
        Name:       "test_vm",
        Hypervisor: "virtualbox",
        CPU:        2,
        Memory:     "4GB",
        DiskSize:   "50GB",
        OS:         "ubuntu-22.04",
        Network:    "local_network",
        Actions:    []string{"start", "stop", "restart"},
    }

    yamlData, err := parser.ToYAML(vm)
    if err != nil {
        log.Fatalf("Error generating YAML: %v", err)
    }

    fmt.Println("Generated YAML:")
    fmt.Println(yamlData)
}
```

### Running Tests

To run the tests for the `go-openinfra` library, you can use the following command:

```bash
go test ./internal/parser -v
```

## Contributing

We welcome contributions to `go-openinfra`. To contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new pull request.

## License

`go-openinfra` is licensed under the Apache 2.0 License. See [LICENSE](LICENSE) for more details.
