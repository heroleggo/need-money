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
	var stack Stack

	if !strings.HasPrefix(data, "<?xml") {
		return ""
	}
	// remove <?xml version="1.0" encoding="UTF-8"?> from data
	position := strings.Index(data, "?>")
	data = data[position+2:]

	// get information from data xml
	for i := 0; i < len(data); i++ {
		if data[i] == '<' {

		}
	}

	return "hehe"
}

func parseWithJson(data string) string {
	return "haha"
}

func parseWithDefault(data string) string {
	return "hihi"
}
