package pkg

import (
	"os"
	"sort"
)

func GetFiles(dir string, sortAlphabetic bool) ([]string, error) {
	file, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	list, _ := file.Readdirnames(0) // 0 to read all files and folders

	if sortAlphabetic {
		sort.Strings(list)
	}

	return list, nil
}

func ReadFile(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
