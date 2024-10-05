package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ZipAndRename takes a folder or file, renames all files by appending ".rename",
// and then compresses the files into a zip archive.
func ZipAndRename(src string, destZip string) error {
	// Check if src is a file or directory
	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Rename all files (in case of a directory)
	if fileInfo.IsDir() {
		err := renameFilesInDir(src)
		if err != nil {
			return fmt.Errorf("failed to rename files: %v", err)
		}
	} else {
		// For a single file, rename it
		err := renameFile(src)
		if err != nil {
			return fmt.Errorf("failed to rename file: %v", err)
		}
	}

	// Create a zip file
	zipFile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Add the file or directory to the zip
	if fileInfo.IsDir() {
		err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			return addToZip(zipWriter, path, src)
		})
		if err != nil {
			return err
		}
	} else {
		err = addToZip(zipWriter, src, filepath.Dir(src))
		if err != nil {
			return err
		}
	}

	return nil
}

// renameFilesInDir renames all the files in the directory by appending ".rename".
func renameFilesInDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			oldPath := filepath.Join(dir, file.Name())
			err := renameFile(oldPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// renameFile renames the file by appending ".rename" to the filename.
func renameFile(filePath string) error {
	newPath := filePath + ".rename"
	err := os.Rename(filePath, newPath)
	if err != nil {
		return err
	}
	return nil
}

// addToZip adds a file or directory to the zip archive.
func addToZip(zipWriter *zip.Writer, filePath string, baseDir string) error {
	relativePath, err := filepath.Rel(baseDir, filePath)
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return nil // Skip directories, only compress files
	}

	// Open the file to be zipped
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create the zip header
	zipHeader, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	zipHeader.Name = strings.ReplaceAll(relativePath, "\\", "/") // For cross-platform compatibility
	zipWriterFile, err := zipWriter.Create(zipHeader.Name)
	if err != nil {
		return err
	}

	// Copy the file content to the zip
	_, err = io.Copy(zipWriterFile, file)
	if err != nil {
		return err
	}

	return nil
}
