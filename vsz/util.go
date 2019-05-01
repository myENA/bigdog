package vsz

import (
	"fmt"
	"strings"
)

func buildQueryParamString(queryParams map[string]string) string {
	// TODO: Make more efficient?
	paramParts := make([]string, 0)

	for name, value := range queryParams {
		value = strings.TrimSpace(value)
		if "" != value {
			paramParts = append(paramParts, fmt.Sprintf("%s=%s", name, value))
		}
	}

	if 0 < len(paramParts) {
		return fmt.Sprintf("?%s", strings.Join(paramParts, "&"))
	} else {
		return ""
	}
}
