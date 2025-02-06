package parser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// ParseFile читает и парсит YAML-файл OpenInfra
func ParseFile(filename string) (*OpenInfraSpec, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл: %w", err)
	}

	var rawSpec struct {
		Version   string     `yaml:"openinfra"`
		Info      Info       `yaml:"info"`
		Providers []Provider `yaml:"providers"`
	}

	if err := yaml.Unmarshal(data, &rawSpec); err != nil {
		return nil, fmt.Errorf("ошибка при разборе YAML: %w", err)
	}

	spec := &OpenInfraSpec{
		Version:   rawSpec.Version,
		Info:      rawSpec.Info,
		Providers: make(map[string]Provider),
	}

	// Добавляем провайдеров в карту
	for _, p := range rawSpec.Providers {
		spec.Providers[p.Name] = p
	}

	return spec, nil
}
