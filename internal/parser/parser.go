package parser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/Ilya-Guyduk/go-openinfra/pkg/provider"
	"github.com/Ilya-Guyduk/go-openinfra/pkg/provider/aws"
	"github.com/Ilya-Guyduk/go-openinfra/pkg/provider/virtualbox"
)

// ParseFile читает и парсит YAML-файл OpenInfra
func ParseFile(filename string) (*OpenInfraSpec, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл: %w", err)
	}

	var rawSpec struct {
		Version   string               `yaml:"openinfra"`
		Info      Info                 `yaml:"info"`
		Providers []*Provider          `yaml:"providers"`
		Resources []ResourceDefinition `yaml:"components"`
	}

	if err := yaml.Unmarshal(data, &rawSpec); err != nil {
		return nil, fmt.Errorf("ошибка при разборе YAML: %w", err)
	}

	spec := &OpenInfraSpec{
		Version:   rawSpec.Version,
		Info:      rawSpec.Info,
		Resources: ConvertResources(rawSpec.Resources),
		Providers: make(map[string]*Provider),
	}

	// Заполняем карту провайдеров
	for _, p := range rawSpec.Providers {
		spec.Providers[p.Name] = p
	}

	// Инициализируем Executor у провайдеров
	for _, p := range rawSpec.Providers {
		exec, err := createProviderExecutor(p)
		if err != nil {
			return nil, fmt.Errorf("ошибка при инициализации провайдера %s: %w", p.Name, err)
		}

		// Назначаем Executor, который реализует интерфейс provider.Provider
		spec.Providers[p.Name].Executor = exec
	}

	return spec, nil
}

// ProviderDef — временная структура для парсинга YAML
type ProviderDef struct {
	Name              string            `yaml:"name"`
	Type              string            `yaml:"type"`
	ConnectionDetails map[string]string `yaml:"connection_details"`
}

// createProvider создает провайдер на основе типа
func createProviderExecutor(p *Provider) (provider.Provider, error) {
	switch p.Type {
	case "aws":
		return aws.New(p.Name, p.ConnectionDetails), nil
	case "virtualbox":
		return virtualbox.New(p.Name, p.ConnectionDetails), nil
	default:
		return nil, fmt.Errorf("неизвестный тип провайдера: %s", p.Type)
	}
}
