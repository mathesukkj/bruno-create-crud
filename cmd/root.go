/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mathesukkj/gocat/internal/blocks"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bcrud",
	Short: "A tool for quickly creating Bruno requests, such as POST, GET, PUT...all directly in your terminal!",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("no entity name provided")
		}

		if len(args) < 2 {
			return errors.New("url not provided")
		}

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		args[1] = strings.TrimSuffix(args[1], "/")

		if headersStr == "" {
			return nil
		}

		headersArr := strings.Split(headersStr, ":")
		if len(headersArr) == 1 || len(headersArr)%2 != 0 {
			return errors.New("headers in wrong format!! use key:value")
		}

		for i := 0; i < len(headersArr); i += 2 {
			headers[headersArr[i]] = headersArr[i+1]
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		url := args[1]
		methods := []string{"get", "post", "get", "put", "delete"}
		paths := []string{"/", "/", "/1", "/1", "/1"}
		actions := []string{"List", "Create", "Show", "Update", "Delete"}

		for i := range len(actions) {
			name := getFormattedName(strings.Title(args[0]), actions[i])

			f, err := os.Create(fmt.Sprintf("%s %s.bru", actions[i], name))
			if err != nil {
				panic(err)
			}

			meta := blocks.Meta(actions[i], name, i)
			f.Write([]byte(meta))

			data := blocks.Method(methods[i], url, strings.ToLower(name), paths[i])
			f.Write([]byte(data))

			if len(headers) > 0 {
				headersBlock := blocks.Headers(headers)
				f.Write([]byte(headersBlock))
			}

			f.Write([]byte("\n"))
		}
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

func getFormattedName(name, action string) string {
	if action == "List" {
		if string(name[len(name)-1]) == "y" {
			return string(name[0:len(name)-1]) + "ies"
		}
		return name + "s"
	}

	return name
}