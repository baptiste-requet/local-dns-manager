package vmmcore

import (
	"os"
	"testing"
)

func TestAddLineInFile(t *testing.T) {
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	err = addLineInFile(file.Name(), "test line")
	if err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	expected := "test line\n"
	if string(content) != expected {
		t.Fatalf("expected %q, got %q", expected, string(content))
	}
}

func TestRemoveLine(t *testing.T) {
    // Create a temporary file for testing
    file, err := os.CreateTemp("", "testfile")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(file.Name())

    // Write some test data to the file
    _, err = file.WriteString("line 1\nline 2\nline 3\n")
    if err != nil {
        t.Fatal(err)
    }

    // Close the file
    err = file.Close()
    if err != nil {
        t.Fatal(err)
    }

    // Call the removeLine function to remove lines that match the regular expression
	line2 := "line 2"
    err = removeLine(file.Name(), line2)
    if err != nil {
        t.Fatal(err)
    }

    // Read the contents of the file and check that the line was removed
    data, err := os.ReadFile(file.Name())
    if err != nil {
        t.Fatal(err)
    }

    expected := "line 1\nline 3\n"
    if string(data) != expected {
        t.Fatalf("expected %q, got %q", expected, string(data))
    }
}
