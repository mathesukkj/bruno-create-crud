package blocks

import "fmt"

func Meta(action, name string, seq int) string {
	return fmt.Sprintf(`meta {
  name: %s %s
  type: http
  seq: %d
}

`, action, name, seq)
}
