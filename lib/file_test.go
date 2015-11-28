package lib

import (
	"bufio"
	"os"
	"testing"
)

// Returns the directory relative to the directory the test is running in where
// resource files can be found.
func dataDir(dirname string) string {
	if len(os.Getenv("TEST_TMPDIR")) > 0 {
		return dirname
	}
	return "."
}

func TestHello(t *testing.T) {
	// The data file is put into the directory relative to the *runfiles
	// dir, in a directory corresponding to the package name.
	file, err := os.Open(dataDir("lib") + "/file.txt")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if text := scanner.Text(); text != SayHello() {
			t.Errorf("Mismatch: '%v'", text)
		}
	}

}
