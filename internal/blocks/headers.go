package blocks

import "fmt"

func Headers(headers map[string]string) string {
	block := fmt.Sprint("headers {\n")

	for key, value := range headers {
		block += fmt.Sprintf("  %s: %s\n", key, value)
	}

	block += "}"

	return block
}
