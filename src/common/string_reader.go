package common

import (
	"bufio"
	"bytes"
	"encoding/json"
	"strings"
)

func StringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func JsonPretty(result string) string {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(result), "", "  ")
	if err != nil {
		return result
	}
	pretty := string(prettyJSON.Bytes())
	return pretty
}
