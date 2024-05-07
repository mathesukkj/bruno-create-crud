package cli

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Cli(args []string) {
	if len(args) == 0 {
		log.Fatal("No name given!")
	}

	url := "http://localhost"
	if len(args) > 1 {
		url = args[1]
	}

	var headers []string
	if len(args) > 2 {
		headers = strings.Split(args[2], ":")
	}

	methods := []string{"get", "post", "get", "put", "delete"}
	paths := []string{"/", "/", "/1", "/1", "/1"}
	actions := []string{"list", "create", "show", "update", "delete"}

	for i := range 5 {
		name := args[0] + "s"

		f, err := os.Create(fmt.Sprintf("%s %s.bru", actions[i], name))
		if err != nil {
			panic(err)
		}

		meta := Meta(actions[i], name, i)
		f.Write([]byte(meta))

		data := Method(methods[i], url, name, paths[i])
		f.Write([]byte(data))

		if headers != nil {
			f.Write([]byte(Headers(headers[0], headers[1])))
		}

		f.Write([]byte("\n"))
	}
}

func Meta(action, name string, seq int) string {
	return fmt.Sprintf(`meta {
  name: %s %s
  type: http
  seq: %d
}

`, action, name, seq)
}

func Method(method, url, name, path string) string {
	return fmt.Sprintf(`%s {
  url: %s/%s%s
  body: none
  auth: none
}`, method, url, name, path)
}

func Headers(headerName, headerValue string) string {
	return fmt.Sprintf(`headers {
  %s: %s
}`, headerName, headerValue)
}
