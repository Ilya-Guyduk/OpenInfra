package parser

// convertToInt безопасно преобразует значение в целое число
func convertToInt(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case float64:
		return int(v)
	default:
		return 0
	}
}
