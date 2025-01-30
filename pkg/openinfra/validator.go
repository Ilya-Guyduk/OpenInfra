package openinfra

import (
	"github.com/xeipuuv/gojsonschema"
)

// ValidateSchema проверяет, соответствует ли спецификация JSON Schema.
func ValidateSchema(spec interface{}, schemaPath string) error {
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
	documentLoader := gojsonschema.NewGoLoader(spec)

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
