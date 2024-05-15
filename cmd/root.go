/*
Copyright © 2024 Matheus Kemuel <kemuel.g7363@gmail.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mathesukkj/bruno-create-crud/internal/blocks"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bcrud",
	Short: "A tool for quickly creating Bruno requests, such as POST, GET, PUT...all directly in your terminal!",
	Args: func(cmd *cobra.Command, args []string) error {
		return checkArgs(args)
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return preRun(args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
}

var headersStr string
var headers = make(map[string]string)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().
		StringVarP(&headersStr, "header", "H", "", "add headers to the request")
}

func run(args []string) {
	url := args[1]
	methods := []string{"get", "post", "get", "put", "delete"}
	paths := []string{"/", "/", "/1", "/1", "/1"}
	actions := []string{"List", "Create", "Show", "Update", "Delete"}

	for i := range len(actions) {
		fileContent := generateFileContents(args[0], url, methods[i], paths[i], actions[i], headers)
		filename := fmt.Sprintf(
			"%s %s.bru",
			actions[i],
			formatName(strings.Title(args[0]), actions[i]),
		)

		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}

		_, err = f.WriteString(fileContent)
		if err != nil {
			panic(err)
		}
	}
}

func generateFileContents(
	entity, url string,
	method string,
	path string,
	action string,
	headers map[string]string,
) string {
	var fileContents strings.Builder

	name := formatName(strings.Title(entity), action)
	meta := blocks.Meta(action, name, 0)
	fileContents.WriteString(meta)

	data := blocks.Method(method, url, strings.ToLower(name), path)
	fileContents.WriteString(data)

	if len(headers) > 0 {
		headersBlock := blocks.Headers(headers)
		fileContents.WriteString(headersBlock)
	}

	return fileContents.String()
}

func formatName(name, action string) string {
	if action == "List" {
		if string(name[len(name)-1]) == "y" {
			return string(name[0:len(name)-1]) + "ies"
		}
		return name + "s"
	}

	return name
}

func checkArgs(args []string) error {
	if len(args) < 1 {
		return errors.New("no entity name provided")
	}

	if len(args) < 2 {
		return errors.New("url not provided")
	}

	return nil
}

func preRun(args []string) error {
	args[1] = strings.TrimSuffix(args[1], "/")

	if headersStr == "" {
		return nil
	}

	return mapHeaders(headersStr)
}

func mapHeaders(headersStr string) error {
	headersArr := strings.Split(headersStr, ":")
	if len(headersArr) == 1 || len(headersArr)%2 != 0 {
		return errors.New("headers in wrong format!! use key:value")
	}

	for i := 0; i < len(headersArr); i += 2 {
		headers[headersArr[i]] = headersArr[i+1]
	}

	return nil
}
