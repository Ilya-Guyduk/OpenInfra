package parser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/Ilya-Guyduk/go-openinfra/pkg/infra"
)

// OpenInfraSpec описывает структуру корневого документа OpenInfra
type OpenInfraSpec struct {
	Version      string               `yaml:"openinfra"`
	Info         Info                 `yaml:"info"`
	Providers    []Provider           `yaml:"providers"`
	Resources    []ResourceDefinition `yaml:"components"`
	Dependencies []Dependency         `yaml:"dependencies"`
}

// Info содержит общую информацию о спецификации
type Info struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
	Contact     struct {
		Name  string `yaml:"name"`
		Email string `yaml:"email"`
	} `yaml:"contact"`
	License struct {
		Name string `yaml:"name"`
		URL  string `yaml:"url"`
	} `yaml:"license"`
}

type Provider struct {
	Name              string            `yaml:"name"`
	Type              string            `yaml:"type"`
	ConnectionDetails map[string]string `yaml:"connection_details"`
}

type Component struct {
	Type       string                 `yaml:"type"`
	Name       string                 `yaml:"name"`
	Provider   string                 `yaml:"provider"` // Ссылается на провайдера
	Properties map[string]interface{} `yaml:"properties"`
	Actions    []string               `yaml:"actions"`
}

// ResourceDefinition описывает ресурс в конфигурации
type ResourceDefinition struct {
	Type         infra.ResourceType     `yaml:"type"`
	Provider     string                 `yaml:"provider"`
	Name         string                 `yaml:"name"`
	Properties   map[string]interface{} `yaml:"properties"`
	Actions      []string               `yaml:"actions"`
	Dependencies []Dependency           `yaml:"dependencies"`
}

// GetName возвращает имя ресурса
func (rd ResourceDefinition) GetName() string {
	return rd.Name
}

// GetType возвращает тип ресурса
func (rd ResourceDefinition) GetType() infra.ResourceType {
	return rd.Type
}

// Dependency описывает зависимости между ресурсами
type Dependency struct {
	Resource  string   `yaml:"resource"`
	DependsOn []string `yaml:"depends_on"`
}

// ParseFile читает и парсит YAML-файл OpenInfra
func ParseFile(filename string) (*OpenInfraSpec, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл: %w", err)
	}

	var spec OpenInfraSpec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("ошибка при разборе YAML: %w", err)
	}

	// Преобразуем ресурсы в конкретные типы
	for i, resource := range spec.Resources {
		switch resource.Type {
		case infra.ResourceVirtualMachine:
			// Преобразуем в VirtualMachine
			var vm infra.VirtualMachine
			vm.Name = resource.Name
			vm.Hypervisor = resource.Provider
			vm.CPU = convertToInt(resource.Properties["cpu"]) // используем функцию convertToInt
			vm.Memory = resource.Properties["memory"].(string)
			vm.DiskSize = resource.Properties["disk_size"].(string)
			vm.OS = resource.Properties["os"].(string)
			vm.Network = resource.Properties["network"].(string)
			vm.Actions = resource.Actions

			// Преобразуем в ResourceDefinition
			spec.Resources[i] = ResourceDefinition{
				Type:         infra.ResourceVirtualMachine,
				Name:         vm.Name,
				Properties:   map[string]interface{}{"hypervisor": vm.Hypervisor, "cpu": vm.CPU, "memory": vm.Memory, "disk_size": vm.DiskSize, "os": vm.OS, "network": vm.Network},
				Actions:      vm.Actions,
				Dependencies: resource.Dependencies,
			}
		case infra.ResourceNetwork:
			// Преобразуем в NetworkResource
			var network infra.NetworkResource
			network.Name = resource.Name
			network.CIDR = resource.Properties["cidr"].(string)
			network.Gateway = resource.Properties["gateway"].(string)
			network.DNSServers = make([]string, 0)
			if dns, ok := resource.Properties["dns_servers"].([]interface{}); ok {
				for _, d := range dns {
					network.DNSServers = append(network.DNSServers, d.(string))
				}
			}
			network.Actions = resource.Actions

			// Преобразуем в ResourceDefinition
			spec.Resources[i] = ResourceDefinition{
				Type:         infra.ResourceNetwork,
				Name:         network.Name,
				Properties:   map[string]interface{}{"cidr": network.CIDR, "gateway": network.Gateway, "dns_servers": network.DNSServers},
				Actions:      network.Actions,
				Dependencies: resource.Dependencies,
			}
		}
	}

	return &spec, nil
}

// convertToInt безопасно преобразует значение в целое число
func convertToInt(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case float64:
		return int(v)
	default:
		return 0
	}
}
