package blocks

import "fmt"

func Method(method, url, name, path string) string {
	return fmt.Sprintf(`%s {
  url: %s/%s%s
  body: json
  auth: none
}

`, method, url, name, path)
}
