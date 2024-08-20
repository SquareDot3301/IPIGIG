package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func findFileInDirs(fileName string, dirs []string) (string, error) {
	for _, dir := range dirs {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && info.Name() == fileName {
				return fmt.Errorf("found: %s", path)
			}
			return nil
		})
		if err != nil && err.Error()[:6] == "found:" {
			return err.Error()[7:], nil
		}
	}
	return "", fmt.Errorf("file not found: %s", fileName)
}

func readFileContent(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func GetFiles(fileNames []string) (map[string]string, error) {
	dirs := []string{
		filepath.Join(os.Getenv("HOME"), "Desktop"),
		filepath.Join(os.Getenv("HOME"), "Documents"),
		filepath.Join(os.Getenv("HOME"), "Downloads"),
		filepath.Join(os.Getenv("HOME"), "Music"),
		filepath.Join(os.Getenv("HOME"), "Pictures"),
		filepath.Join(os.Getenv("HOME"), "Videos"),
		filepath.Join(os.Getenv("HOME"), "Bureau"),
		filepath.Join(os.Getenv("HOME"), "Téléchargements"),
		filepath.Join(os.Getenv("HOME"), "Musique"),
		filepath.Join(os.Getenv("HOME"), "Photos"),
		filepath.Join(os.Getenv("HOME"), "Vidéos"),
	}

	contents := make(map[string]string)

	for _, fileName := range fileNames {
		filePath, err := findFileInDirs(fileName, dirs)
		if err != nil {
			fmt.Printf("Warning: %v\n", err)
			continue
		}

		content, err := readFileContent(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", fileName, err)
			continue
		}

		contents[fileName] = content
	}

	return contents, nil
}
