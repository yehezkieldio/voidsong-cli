package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func CheckPackageJSON() bool {
	if _, err := os.Stat("package.json"); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func CheckBunProject() bool {
	if _, err := os.Stat("bun.lockb"); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func findFilesByRegex(root, pattern string) ([]string, error) {
	var matches []string
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %v", err)
	}

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && re.MatchString(info.Name()) {
			matches = append(matches, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return matches, nil
}

func CheckConfigurationFile(regex string) bool {
	files, err := findFilesByRegex(".", regex)
	if err != nil {
		return false
	}
	if len(files) == 0 {
		return false
	}
	return true
}
