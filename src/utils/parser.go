package utils

func Parse(str string, method string) string {
	if method == "xml" {
		return parseWithXml(str)
	} else if method == "json" {
		return parseWithJson(str)
	} else {
		return parseWithDefault(str)
	}
}

func parseWithXml(str string) string {
	return "hehe"
}

func parseWithJson(str string) string {
	return "haha"
}

func parseWithDefault(str string) string {
	return "hihi"
}
