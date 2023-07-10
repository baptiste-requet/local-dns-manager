package vmmcore

import (
	"fmt"
)

const (
	HOST_FILE_PATH = "c:\\windows\\System32\\drivers\\etc\\hosts"

	NAME_PRINT_LEN    = 40
	IP_PRINT_LEN      = 15
	COMMENT_PRINT_LEN = 50
)

func GetAllEntries() (entries []Entry) {
	lines, err := getAllLinesMatchingRegexp(HOST_FILE_PATH, getIpAndNameRegex())
	if err != nil {
		fmt.Println(err)
		return []Entry{}
	}

	for _, line := range lines {
		entry := parseEntryFromString(line)
		if entry.isNotNull() {
			entries = append(entries, entry)
		}
	}

	return entries
}

func AddEntry(ip, name, comment string) {
	entry := Entry{
		Name:    name,
		Ip:      ip,
		Comment: comment,
	}

	if !entry.isValid() {
		fmt.Println("Cannot save this entry: it is unvalid. Check that the IP is correct and that its name is not empty.")
		return
	}

	allEntries := GetAllEntries()
	if isIpUsed(entry.Ip, allEntries) {
		fmt.Printf("The IP '%s' is already registered.\n", entry.Ip)
		return
	}
	if isNameUsed(entry.Name, allEntries) {
		fmt.Printf("The name '%s' is already used.\n", entry.Name)
		return
	}

	if err := addLineInFile(HOST_FILE_PATH, entry.toString()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Entry added.")
	}
}

func RemoveEntryByIp(ip string) {
	allEntries := GetAllEntries()
	if !isIpUsed(ip, allEntries) {
		fmt.Printf("The IP '%s' is not registered.\n", ip)
		return
	}

	if err := removeLine(HOST_FILE_PATH, ip + " \\S+"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Entry removed.")
	}
}

func RemoveEntryByName(name string) {
	allEntries := GetAllEntries()
	if !isNameUsed(name, allEntries) {
		fmt.Printf("The name '%s' is not registered.\n", name)
		return
	}

	if err := removeLine(HOST_FILE_PATH, IP_REGEX_STR + " " + name); err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Entry removed.")
	}
}

func PrintRegisteredEntries() {
	entries := GetAllEntries()
	fmt.Printf("%d IP(s) registered\n", len(entries))
	fmt.Printf("%-*s %-*s %-*s\n",
		NAME_PRINT_LEN, "Name",
		IP_PRINT_LEN, "IP",
		COMMENT_PRINT_LEN, "Comment")
	for _, entry := range entries {
		fmt.Printf("%-*s ", NAME_PRINT_LEN, entry.Name)
		fmt.Printf("%-*s ", IP_PRINT_LEN, entry.Ip)
		if entry.Comment != "" {
			fmt.Printf("%-*s", COMMENT_PRINT_LEN, entry.Comment)
		}
		fmt.Println()
	}
}

func isIpUsed(ip string, entries []Entry) bool {
	for _, entry := range entries {
		if entry.Ip == ip {
			return true
		}
	}
	return false
}

func isNameUsed(name string, entries []Entry) bool {
	for _, entry := range entries {
		if entry.Name == name {
			return true
		}
	}
	return false
}
