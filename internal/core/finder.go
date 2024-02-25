package core

import (
	"log"
	"os"
	"path/filepath"
)

func GetFullPathToFile(path string) string {
	return filepath.Join(path, "cert.pem")

}

func FindCertificate(path string) []string {

	list := folderList(path)

	return list

}

func folderList(path string) []string {
	var listDir []string
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		fullPath := filepath.Join(path, f.Name())
		listDir = append(listDir, fullPath)
		if f.IsDir() {
			folderList(fullPath)
		}

	}
	return listDir
}
