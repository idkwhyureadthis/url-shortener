package linkverify

import "strings"

const PATTERN = "12345678910ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func VerifyLink(link string) bool {
	for _, elem := range PATTERN {
		link = strings.ReplaceAll(link, string(elem), "")
	}
	return len(link) == 0
}
