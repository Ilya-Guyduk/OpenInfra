package parser

import (
	"fmt"

	"github.com/Ilya-Guyduk/go-openinfra/pkg/infra"
	"github.com/Ilya-Guyduk/go-openinfra/pkg/provider"
)

// OpenInfraSpec описывает структуру корневого документа OpenInfra
type OpenInfraSpec struct {
	Version      string               `yaml:"openinfra"`
	Info         Info                 `yaml:"info"`
	Providers    map[string]*Provider `yaml:"providers"`
	Resources    []ResourceDefinition `yaml:"components"`
	Dependencies []Dependency         `yaml:"dependencies"`
}

func (ois *OpenInfraSpec) GetProviderList() []*Provider {
	var providerList []*Provider
	for _, provider := range ois.Providers {
		providerList = append(providerList, provider)
	}
	return providerList
}

func (ois *OpenInfraSpec) GetProviderMap() map[string]*Provider {
	return ois.Providers
}

func (ois *OpenInfraSpec) GetProviderByName(name string) (*Provider, error) {
	if provider := ois.Providers[name]; provider != nil {
		return provider, nil
	} else {
		return nil, fmt.Errorf("Provider not found!")
	}
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

// Provider — провайдер ресурсов
type Provider struct {
	Name              string            `yaml:"name"`
	Type              string            `yaml:"type"`
	ConnectionDetails map[string]string `yaml:"connection_details"`
	Executor          provider.Provider
}

// Component — описание компонента инфраструктуры
type Component struct {
	Type       string                 `yaml:"type"`
	Name       string                 `yaml:"name"`
	Provider   string                 `yaml:"provider"`
	Properties map[string]interface{} `yaml:"properties"`
	Actions    []string               `yaml:"actions"`
}

// ResourceDefinition описывает конкретный ресурс
type ResourceDefinition struct {
	Type         infra.ResourceType     `yaml:"type"`
	Provider     string                 `yaml:"provider"`
	Name         string                 `yaml:"name"`
	Properties   map[string]interface{} `yaml:"properties"`
	Actions      []string               `yaml:"actions"`
	Dependencies []Dependency           `yaml:"dependencies"`
}

// Dependency описывает зависимости между ресурсами
type Dependency struct {
	Resource  string   `yaml:"resource"`
	DependsOn []string `yaml:"depends_on"`
}
