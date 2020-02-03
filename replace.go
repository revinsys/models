package main

import "strings"

func GenerateFromTemplate(template string, keys map[string]string) string {
	generated := template
	for key, value := range keys {
		generated = strings.ReplaceAll(generated, key, value)
	}

	return generated
}
