package utils

import "strings"

func Parse(str string, method string) string {
	if method == "xml" {
		return parseWithXml(str)
	} else if method == "json" {
		return parseWithJson(str)
	} else {
		return parseWithDefault(str)
	}
}

func parseWithXml(data string) string {
	splittedData := strings.Split(data, "\n")
	return "hehe"
}

func parseWithJson(data string) string {
	return "haha"
}

func parseWithDefault(data string) string {
	return "hihi"
}
