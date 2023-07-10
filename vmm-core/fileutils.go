package vmmcore

import (
	"bufio"
	"os"
	"regexp"
)

func getAllLinesMatchingRegexp(filePath string, regExpr *regexp.Regexp) (lines []string, err error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matched := regExpr.MatchString(line)
		if matched {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func addLineInFile(filePath, line  string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(line + "\n")
	if err != nil {
		return err
	}

	return nil
}

func removeLine(filePath string, regex string) error {
    re := regexp.MustCompile(regex)
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string
    for scanner.Scan() {
        line := scanner.Text()
        if !re.MatchString(line) {
            lines = append(lines, line)
        }
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    output := []byte{}
    for _, line := range lines {
        output = append(output, []byte(line+"\n")...)
    }

    return os.WriteFile(filePath, output, 0644)
}
