package openinfra

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

// OpenInfra представляет собой структуру для хранения всей спецификации.
type OpenInfra struct {
	Version      string     `yaml:"openinfra"`
	Info         Info       `yaml:"info"`
	Resources    []Resource `yaml:"resources"`
	Dependencies []Dependency `yaml:"dependencies"`
}

// Info содержит метаданные спецификации.
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

// Resource описывает ресурс инфраструктуры.
type Resource struct {
	Type       string                 `yaml:"type"`
	Name       string                 `yaml:"name"`
	Properties map[string]interface{} `yaml:"properties"`
	Actions    []string               `yaml:"actions"`
}

// Dependency описывает зависимости между ресурсами.
type Dependency struct {
	Resource  string   `yaml:"resource"`
	DependsOn []string `yaml:"depends_on"`
}

// Load загружает спецификацию из YAML-файла.
func Load(filename string) (*OpenInfra, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var spec OpenInfra
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &spec, nil
}

// Validate проверяет спецификацию на соответствие JSON Schema.
func (o *OpenInfra) Validate(schemaPath string) error {
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
	documentLoader := gojsonschema.NewGoLoader(o)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("failed to validate schema: %w", err)
	}

	if !result.Valid() {
		var errs string
		for _, desc := range result.Errors() {
			errs += fmt.Sprintf("- %s\n", desc)
		}
		return errors.New("validation errors:\n" + errs)
	}

	return nil
}
