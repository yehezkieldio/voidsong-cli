package utils

import "os"

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
