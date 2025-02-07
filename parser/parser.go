package parser

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// ParseFile читает и парсит YAML-файл OpenInfra.
func ParseFile(filename string) (*OpenInfraSpec, error) {
	// Проверяем, существует ли файл
	fileInfo, err := os.Stat(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("ошибка: файл %s не найден", filename)
		}
		return nil, fmt.Errorf("ошибка при получении информации о файле: %w", err)
	}

	// Проверяем права на чтение файла
	file, err := os.Open(filename)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			return nil, fmt.Errorf("ошибка: недостаточно прав для чтения файла %s", filename)
		}
		return nil, fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	defer file.Close()

	// Проверяем, не пуст ли файл
	if fileInfo.Size() == 0 {
		return nil, fmt.Errorf("ошибка: файл %s пуст", filename)
	}

	// Читаем содержимое файла
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла %s: %w", filename, err)
	}

	var rawSpec struct {
		Version      string       `yaml:"openinfra"`
		Info         Info         `yaml:"info"`
		Providers    []Provider   `yaml:"providers"`
		Resources    []Resource   `yaml:"components"`
		Dependencies []Dependency `yaml:"dependencies"`
	}

	// Парсим YAML
	if err := yaml.Unmarshal(data, &rawSpec); err != nil {
		return nil, fmt.Errorf("ошибка: некорректное форматирование YAML в файле %s: %w", filename, err)
	}

	// Создаём структуру с провайдерами в виде карты
	spec := &OpenInfraSpec{
		Version:      rawSpec.Version,
		Info:         rawSpec.Info,
		Providers:    make(map[string]Provider),
		Resources:    make(map[string]Resource),
		Dependencies: rawSpec.Dependencies,
	}

	for _, p := range rawSpec.Providers {
		spec.Providers[p.Name] = p
	}
	for _, r := range rawSpec.Resources {
		spec.Resources[r.Name] = r
	}

	return spec, nil
}
