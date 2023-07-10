package vmmcore

import "regexp"

const (
	IP_REGEX_STR          = "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}"
	IP_AND_NAME_REGEX_STR = IP_REGEX_STR + " \\S+"
)

var (
	ipRegex *regexp.Regexp
	ipAndNameRegex *regexp.Regexp
)

func getIpRegex() *regexp.Regexp {
	if ipRegex == nil {
		ipRegex = regexp.MustCompile(IP_REGEX_STR)
	}
	return ipRegex
}

func getIpAndNameRegex() *regexp.Regexp {
	if ipAndNameRegex == nil {
		ipAndNameRegex = regexp.MustCompile(IP_AND_NAME_REGEX_STR)
	}
	return ipAndNameRegex
}
