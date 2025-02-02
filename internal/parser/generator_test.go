package parser

import (
	"testing"

	"github.com/Ilya-Guyduk/go-openinfra/pkg/infra"
	"github.com/stretchr/testify/assert"
)

func TestGenerateYAMLVirtualMachine(t *testing.T) {
	vm := infra.VirtualMachine{
		Name:       "test_vm",
		Hypervisor: "virtualbox",
		CPU:        2,
		Memory:     "4GB",
		DiskSize:   "50GB",
		OS:         "ubuntu-22.04",
		Network:    "test_network",
		Actions:    []string{"start", "stop", "restart"},
	}

	yamlData, err := ToYAML(vm)
	assert.NoError(t, err)
	assert.Contains(t, yamlData, "name: test_vm")
	assert.Contains(t, yamlData, "provider: virtualbox")
	assert.Contains(t, yamlData, "cpu: 2")
	assert.Contains(t, yamlData, "memory: 4GB")
	assert.Contains(t, yamlData, "disk_size: 50GB")
	assert.Contains(t, yamlData, "os: ubuntu-22.04")
	assert.Contains(t, yamlData, "network: test_network")
	assert.Contains(t, yamlData, "actions:")
	assert.Contains(t, yamlData, "  - start")
	assert.Contains(t, yamlData, "  - stop")
	assert.Contains(t, yamlData, "  - restart")
}

func TestGenerateYAMLNetworkResource(t *testing.T) {
	network := infra.NetworkResource{
		Name:       "test_network",
		CIDR:       "192.168.1.0/24",
		Gateway:    "192.168.1.1",
		DNSServers: []string{"8.8.8.8", "8.8.4.4"},
	}

	yamlData, err := ToYAML(network)
	assert.NoError(t, err)
	assert.Contains(t, yamlData, "name: test_network")
	assert.Contains(t, yamlData, "cidr: 192.168.1.0/24")
	assert.Contains(t, yamlData, "gateway: 192.168.1.1")
	assert.Contains(t, yamlData, "dns_servers:")
	assert.Contains(t, yamlData, "  - 8.8.8.8")
	assert.Contains(t, yamlData, "  - 8.8.4.4")
}

func TestGenerateYAMLOpenInfraSpec(t *testing.T) {
	spec := &OpenInfraSpec{
		Version: "1.0.0",
		Info: Info{
			Title:       "OpenInfra Specification",
			Description: "A specification for describing infrastructure resources and components.",
			Version:     "1.0.0",
			Contact: struct {
				Name  string `yaml:"name"`
				Email string `yaml:"email"`
			}{
				Name:  "Test User",
				Email: "test@example.com",
			},
			License: struct {
				Name string `yaml:"name"`
				URL  string `yaml:"url"`
			}{
				Name: "Apache 2.0",
				URL:  "https://www.apache.org/licenses/LICENSE-2.0",
			},
		},
		Resources: []ResourceDefinition{
			{
				Type:    infra.ResourceVirtualMachine,
				Name:    "test_vm",
				Actions: []string{"start", "stop"},
			},
			{
				Type:    infra.ResourceNetwork,
				Name:    "test_network",
				Actions: []string{"configure"},
			},
		},
		Dependencies: []Dependency{
			{
				Resource:  "test_vm",
				DependsOn: []string{"test_network"},
			},
		},
	}

	yamlData, err := GenerateYAML(spec)
	assert.NoError(t, err)
	assert.Contains(t, yamlData, "openinfra: 1.0.0")
	assert.Contains(t, yamlData, "title: OpenInfra Specification")
	assert.Contains(t, yamlData, "description: A specification for describing infrastructure resources and components.")
	assert.Contains(t, yamlData, "version: 1.0.0")
	assert.Contains(t, yamlData, "name: test_vm")
	assert.Contains(t, yamlData, "type: virtual_machine")
	assert.Contains(t, yamlData, "actions:")
	assert.Contains(t, yamlData, "  - start")
	assert.Contains(t, yamlData, "  - stop")
}

// UnknownResource - структура, которая реализует интерфейс infra.Resource
type UnknownResource struct {
	Name string `yaml:"name"`
}

func (ur UnknownResource) GetName() string {
	return ur.Name
}

func (ur UnknownResource) GetType() infra.ResourceType {
	return "unknown" // Или другой подходящий тип
}

func (ur UnknownResource) GetProperties() map[string]interface{} {
	return make(map[string]interface{})
}

func TestGenerateYAMLUnknownResource(t *testing.T) {
	// Создаем неизвестный ресурс для теста
	unknownResource := UnknownResource{
		Name: "unknown_resource",
	}

	_, err := ToYAML(unknownResource)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "неизвестный тип ресурса")
}
