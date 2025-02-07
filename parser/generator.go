package parser

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// GenerateYAML генерирует YAML-строку из OpenInfraSpec
func GenerateYAML(spec *OpenInfraSpec) (string, error) {
	// Маршаллинг данных в YAML
	data, err := yaml.Marshal(spec)
	if err != nil {
		return "", fmt.Errorf("ошибка при генерации YAML: %w", err)
	}
	return string(data), nil
}
