package main

import (
	"fmt"
	"log"

	"github.com/your-username/openinfra/pkg/openinfra"
)

func main() {
	// Загрузка спецификации
	spec, err := openinfra.Load("examples/cloud_infra.yaml")
	if err != nil {
		log.Fatalf("Failed to load spec: %v", err)
	}

	// Валидация спецификации
	if err := spec.Validate("schemas/openinfra_schema.json"); err != nil {
		log.Fatalf("Validation failed: %v", err)
	}

	// Получение ресурса по имени
	resource, err := spec.GetResourceByName("web_server")
	if err != nil {
		log.Fatalf("Failed to get resource: %v", err)
	}

	// Выполнение действия над ресурсом
	if err := resource.ExecuteAction("start"); err != nil {
		log.Fatalf("Failed to execute action: %v", err)
	}

	fmt.Println("OpenInfra specification loaded and validated successfully!")
}
