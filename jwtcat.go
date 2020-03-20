package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print(processJwt(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}

// JWT is a triplet of Base64-encoded strings separated by dots.
// We simply decode each in turn.
func processJwt(jwtStr string) string {
	parts := strings.Split(jwtStr, ".")

	result := ""

	for i, part := range parts {
		if i == 2 {
			parts[i] = "[skipped decoding signature]"

		} else {
			decoded, err := base64.StdEncoding.DecodeString(part)
			if err != nil {
				parts[i] = fmt.Sprintf("[Base64 decode failed: %v]", err)

			} else {
				// https://stackoverflow.com/a/29046984/1432488
				var prettyJSON bytes.Buffer
				err = json.Indent(&prettyJSON, decoded, "", "\t")

				if err != nil {
					parts[i] = fmt.Sprintf("[JSON decode failed: %v]", err)

				} else {
					parts[i] = prettyJSON.String()
				}
			}
		}

		result += parts[i] + "\n"
	}

	return result
}
