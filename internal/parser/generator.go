package parser

import (
	"fmt"

	"github.com/Ilya-Guyduk/go-openinfra/pkg/infra"
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

// ToYAML генерирует YAML для одного ресурса
func ToYAML(resource infra.Resource) (string, error) {
	var data []byte
	var err error

	switch r := resource.(type) {
	case infra.VirtualMachine:
		data, err = yaml.Marshal(r)
	case infra.NetworkResource:
		data, err = yaml.Marshal(r)
	default:
		err = fmt.Errorf("неизвестный тип ресурса: %T", r)
	}

	if err != nil {
		return "", fmt.Errorf("ошибка при генерации YAML для ресурса: %w", err)
	}
	return string(data), nil
}
