package parser

import (
	"os"
	"testing"

	"github.com/Ilya-Guyduk/go-openinfra/pkg/infra"
	"github.com/stretchr/testify/assert"
)

const sampleYAML = `
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
  - name: local_virtualbox
    type: virtualbox
    connection_details:
      address: 192.168.1.10
      username: admin
      password: password

  - name: cloud_provider
    type: aws
    connection_details:
      api_endpoint: https://api.cloudprovider.com
      api_key: your_api_key_here

components:
  - type: virtual_machine
    name: local_vm
    provider: local_virtualbox
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
`

func TestProviders(t *testing.T) {
	// Создаём временный YAML-файл
	tmpFile, err := os.CreateTemp("", "openinfra-*.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name()) // Удаляем файл после теста

	_, err = tmpFile.WriteString(sampleYAML)
	assert.NoError(t, err)
	tmpFile.Close() // Закрываем файл, чтобы его можно было прочитать

	// Вызываем функцию ParseFile
	spec, err := ParseFile(tmpFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, spec)

	// Получаем список провайдеров
	providerList := spec.Providers
	assert.NotNil(t, providerList)
	assert.Len(t, providerList, 2, "Ожидалось 2 провайдера")

	// Ожидаемый список провайдеров
	expectedProviders := map[string]string{
		"local_virtualbox": "virtualbox",
		"cloud_provider":   "aws",
	}

	// Проверяем, что все провайдеры из YAML есть в списке
	for _, provider := range providerList {
		expectedType, exists := expectedProviders[provider.Name]
		assert.True(t, exists, "Провайдер %s не найден в ожидаемом списке", provider.Name)
		assert.Equal(t, expectedType, provider.Type, "Тип провайдера %s не соответствует ожидаемому", provider.Name)
	}

}

func TestInfo(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "openinfra-*.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name()) // Удаляем файл после теста

	_, err = tmpFile.WriteString(sampleYAML)
	assert.NoError(t, err)
	tmpFile.Close() // Закрываем файл, чтобы его можно было прочитать

	// Вызываем функцию ParseFile
	spec, err := ParseFile(tmpFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, spec)

	assert.Equal(t, "OpenInfra Specification", spec.Info.Title)
	assert.Equal(t, "A specification for describing infrastructure resources and components.", spec.Info.Description)
	assert.Equal(t, "1.0.0", spec.Info.Version)
	assert.Equal(t, "Your Name", spec.Info.Contact.Name)
	assert.Equal(t, "your.email@example.com", spec.Info.Contact.Email)
	assert.Equal(t, "Apache 2.0", spec.Info.License.Name)
	assert.Equal(t, "https://www.apache.org/licenses/LICENSE-2.0", spec.Info.License.URL)

}

func TestParseFile(t *testing.T) {
	// Создаем временный файл
	tmpFile, err := os.CreateTemp("", "openinfra-*.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name()) // Удаляем файл после теста

	_, err = tmpFile.WriteString(sampleYAML)
	assert.NoError(t, err)
	tmpFile.Close() // Закрываем файл, чтобы его можно было прочитать

	// Вызываем функцию ParseFile
	spec, err := ParseFile(tmpFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, spec)

	// Проверяем основную информацию
	assert.Equal(t, "1.0.0", spec.Version)

	// Проверяем ресурсы
	assert.Len(t, spec.Resources, 2)

	// Проверяем первую ВМ
	vm := spec.Resources[0]
	assert.Equal(t, "local_vm", vm.Name)
	assert.Equal(t, infra.ResourceVirtualMachine, vm.Type)

	properties := vm.Properties
	assert.Equal(t, "virtualbox", properties["provider"])
	if cpu, ok := properties["cpu"].(int); ok {
		assert.Equal(t, int(cpu), properties["cpu"])
	} else if cpu, ok := properties["cpu"].(float64); ok {
		assert.Equal(t, float64(cpu), properties["cpu"])
	}
	assert.Equal(t, "4GB", properties["memory"])
	assert.Equal(t, "50GB", properties["disk_size"])
	assert.Equal(t, "ubuntu-22.04", properties["os"])
	assert.Equal(t, "local_network", properties["network"])
	assert.ElementsMatch(t, []string{"start", "stop", "restart"}, vm.Actions)

	// Проверяем сеть
	network := spec.Resources[1]
	assert.Equal(t, "local_network", network.Name)
	assert.Equal(t, infra.ResourceNetwork, network.Type)

	networkProperties := network.Properties
	assert.Equal(t, "192.168.1.0/24", networkProperties["cidr"])
	assert.Equal(t, "192.168.1.1", networkProperties["gateway"])
	assert.ElementsMatch(t, []string{"8.8.8.8", "8.8.4.4"}, networkProperties["dns_servers"])

	// Проверяем зависимости
	assert.Len(t, spec.Dependencies, 1)
	dep := spec.Dependencies[0]
	assert.Equal(t, "local_vm", dep.Resource)
	assert.ElementsMatch(t, []string{"test_network"}, dep.DependsOn)
}
