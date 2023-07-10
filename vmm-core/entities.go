package vmmcore

import (
	"fmt"
	"strings"
)

type Entry struct {
	Name    string
	Ip      string
	Comment string
}

func (entry Entry) isNull() bool {
	return entry.Name == "" && entry.Ip == ""
}

func (entry Entry) isNotNull() bool {
	return !entry.isNull()
}

func (entry Entry) isIpValid() bool {
	return entry.Ip != "" && getIpRegex().MatchString(entry.Ip)
}

func (entry Entry) isNameValid() bool {
	return entry.Name != "" // TODO check that name does not contains spaces (only url like)
}

func (entry Entry) isValid() bool {
	return entry.isNameValid() && entry.isIpValid()
}

func (entry Entry) toString() string {
	if entry.Comment == "" {
		return fmt.Sprintf("%s %s", entry.Ip, entry.Name)
	} else {
		return fmt.Sprintf("%s %s # %s", entry.Ip, entry.Name, entry.Comment)
	}
}

func parseEntryFromString(entryStr string) (entry Entry) {
	parts := strings.Split(entryStr, "#")
	if len(parts) > 1 {
		entry.Comment = strings.TrimSpace(parts[1])
	}

	ip := getIpRegex().FindString(parts[0])
	entry.Ip = strings.TrimSpace(ip)

	lineWithoutIp := parts[0][len(ip):]
	entry.Name = strings.TrimSpace(lineWithoutIp)

	return entry
}
