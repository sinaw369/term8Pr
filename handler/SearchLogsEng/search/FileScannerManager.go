package search

import (
	"fmt"
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/backscanner"
	"os"
)

type FileScannerManager struct {
	fileNames   []string
	scannerMap  map[string]*backscanner.Scanner
	isEndOfFile map[string]bool
}

func NewFileScannerManager() *FileScannerManager {
	return &FileScannerManager{
		fileNames:   []string{},
		scannerMap:  make(map[string]*backscanner.Scanner),
		isEndOfFile: make(map[string]bool),
	}
}
func (FSM *FileScannerManager) AddFileName(path []string) error {

	for _, fileName := range path {
		if err := FSM.addFile(fileName); err != nil {
			return err
		}

	}

	return nil
}
func (FSM *FileScannerManager) addFile(fileName string) error {
	// Check if the file name already exists in the slice

	// Open the file and create a new scanner.
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	fileStatus, _ := file.Stat()
	if fileStatus.Size() == 0 {
		return nil
	}
	for _, existingFileName := range FSM.fileNames {
		if existingFileName == fileStatus.Name() {
			return nil // File name already exists, do nothing
		}
	}
	// File name does not exist, add it to the slice
	FSM.fileNames = append(FSM.fileNames, fileStatus.Name())

	scanner := backscanner.NewOptions(
		file,
		int(fileStatus.Size()),
		&backscanner.Options{
			ChunkSize:     0,
			MaxBufferSize: 0,
			FileName:      fileStatus.Name(),
			File:          file,
		},
	)
	FSM.scannerMap[fileStatus.Name()] = scanner

	return nil
}
func (FSM *FileScannerManager) CloseAll() {
	for _, scanner := range FSM.scannerMap {
		err := scanner.Close()
		if err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}
}
func (FSM *FileScannerManager) ProcessNextFile() (*backscanner.Scanner, error) {
	if len(FSM.fileNames) == 0 {
		return nil, fmt.Errorf("no more files to process")
	}

	fileName := FSM.fileNames[0]
	FSM.fileNames = FSM.fileNames[1:] // Remove the first file name from the slice

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := int(fileInfo.Size())

	scanner := backscanner.New(file, fileSize)
	FSM.scannerMap[fileName] = scanner

	return scanner, nil
}
