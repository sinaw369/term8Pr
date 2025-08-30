package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/backscanner"
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/params"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// ByTime is a type that allows sorting of LogEntry by time.
type ByTime []map[string]interface{}

// Len, Swap, and Less are methods required by sort.Interface.
func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[j]["time"].(time.Time).Before(a[i]["time"].(time.Time)) }

type SearchLogs struct {
	FilesPath []string
	ValidPath []string
	BasePath  string
	Manager   *FileScannerManager
	Mutex     sync.Mutex
}

func New(BasePath string) *SearchLogs {
	return &SearchLogs{
		FilesPath: []string{},
		ValidPath: []string{},
		BasePath:  BasePath,
		Manager:   NewFileScannerManager(),
		Mutex:     sync.Mutex{},
	}
}

// Clear resets all fields of the SearchLogs struct to their zero values.
func (SL *SearchLogs) Clear() {
	SL.FilesPath = []string{}            // Reset the slice to nil
	SL.Manager = NewFileScannerManager() // Reset the pointer to nil
}

func (SL *SearchLogs) GetFilesInFolder(SLReq params.SearchLogRequest) ([]string, int64, error) {
	startTime := time.Now()
	isend := false
	var jsonData []string
	responseCounter := int64(0)
	wg := new(sync.WaitGroup)
	inProgress := make(chan *backscanner.Scanner, 10)
	inProgressLine := make(chan map[string]interface{}, 2000)

	pathLen, err := SL.GetFileNames(SLReq)
	if pathLen == 0 {
		return jsonData, responseCounter, nil
	}
	if err != nil {
		return nil, 0, err
	}

	for {

		wg.Add(1)
		go func() {
			defer wg.Done()
			err = SL.CreateReaderForFile(inProgress, wg)
			if err != nil {
				fmt.Println("CreateReaderForFile:", err)
			}
		}()
		//2
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			SL.ProcessFile(inProgress, inProgressLine, wg)
		}()
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			SL.ProcessFile(inProgress, inProgressLine, wg)
		}()

		go func() {
			wg.Wait()
			//	close(inProgressLine)
			//	close(isEndOfFileCh)
		}()

		jsonData, responseCounter = SL.ProcessLine2(inProgress, inProgressLine, &isend, SLReq)
		if isend == true {
			break
		}

	}

	//jsonData, responseCounter = FindWordInTexts(SLReq, jsonData)

	//

	//SL.Manager.CloseAll()
	if SLReq.CheckBox == true {
		if SLReq.Limit == 0 {
			fmt.Println("process limit:", "All File", " and spent Time :", time.Now().Sub(startTime).Seconds(), "Seconds")
		} else {
			fmt.Println("process limit:", SLReq.Limit, " and spent Time :", time.Now().Sub(startTime).Seconds(),
				"Seconds =", time.Now().Sub(startTime).Milliseconds(), "Milliseconds")
		}

		return nil, responseCounter, nil
	}
	if SLReq.Limit == 0 {
		fmt.Println("process limit:", "All File", " and spent Time :", time.Now().Sub(startTime).Seconds(), "Seconds")
	} else {
		fmt.Println("process limit:", SLReq.Limit, " and spent Time :", time.Now().Sub(startTime).Seconds(),
			"Seconds =", time.Now().Sub(startTime).Milliseconds(), "Milliseconds")
	}

	return jsonData, responseCounter, nil
}

func (SL *SearchLogs) GetFileNames(filter params.SearchLogRequest) (int, error) {
	SL.FilesPath = nil
	SL.ValidPath = nil
	var folderPath string
	if SL.BasePath != params.EmptyString {
		folderPath = fmt.Sprintf("%s/logs/%s-%s-%s", SL.BasePath, filter.Year, filter.Month, filter.Day)
	} else {
		folderPath = fmt.Sprintf("logs/%s-%s-%s", filter.Year, filter.Month, filter.Day)
	}
	if filter.LogType == "errors" {
		folderPath += "/errors"
	}
	fmt.Println("folderPath:", folderPath)
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		if !strings.Contains(folderPath, "errors") {
			if info.IsDir() && info.Name() == "errors" {
				return filepath.SkipDir
			}
		}
		if filter.NotIncludeFileName == params.EmptyString {
			if info.Mode().IsRegular() && strings.Contains(info.Name(), filter.FileName) {
				SL.FilesPath = append(SL.FilesPath, path)
			}
		} else {
			if info.Mode().IsRegular() && strings.Contains(info.Name(), filter.FileName) && !strings.Contains(info.Name(), filter.NotIncludeFileName) {
				SL.FilesPath = append(SL.FilesPath, path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("err:", err)
		return 0, err
	}
	pathLenth := len(SL.FilesPath)

	for i := 0; i < pathLenth/2; i++ {
		j := pathLenth - i - 1
		SL.FilesPath[i], SL.FilesPath[j] = SL.FilesPath[j], SL.FilesPath[i]
	}
	for _, fileName := range SL.FilesPath {
		// Open the file and create a new scanner
		file, err := os.Open(fileName)
		if err != nil {
			continue
		}
		fileStatus, _ := file.Stat()
		if fileStatus.Size() == 0 {
			//SL.FilesPath = removeElement(SL.FilesPath, fileName)
			continue
		}
		SL.ValidPath = append(SL.ValidPath, fileName)
	}
	return pathLenth, nil
}

func (SL *SearchLogs) CreateReaderForFile(inProgress chan *backscanner.Scanner, wg *sync.WaitGroup) error {

	for _, fileName := range SL.ValidPath {
		// Open the file and create a new scanner
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		fileStatus, _ := file.Stat()
		if fileStatus.Size() == 0 {
			//SL.FilesPath = removeElement(SL.FilesPath, fileName)
			continue
		}
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

		inProgress <- scanner
	}
	return nil
}

func (SL *SearchLogs) ProcessFile(inProgress chan *backscanner.Scanner, inProgressLine chan map[string]interface{}, wg *sync.WaitGroup) {
	wg1 := &sync.WaitGroup{}
	semaphore := make(chan struct{}, 10) // Create a semaphore with capacity 10

	for scanner := range inProgress {
		semaphore <- struct{}{} // Acquire a token from the semaphore
		wg1.Add(1)

		go func(scanner *backscanner.Scanner) {
			defer func() {
				wg1.Done()
				<-semaphore // Release the token back to the semaphore
			}()

			SL.processFile(scanner, inProgress, inProgressLine)
		}(scanner)

	}

	wg1.Wait()
}

func (SL *SearchLogs) processFile(scanner *backscanner.Scanner, inProgress chan *backscanner.Scanner, inProgressLine chan map[string]interface{}) {
	counter := 0
	//wg1 := &sync.WaitGroup{}
	//semaphore := make(chan struct{}, 100) // Create a semaphore with capacity 10
	for i := 0; i < params.DefultReadLine; i++ {
		line, _, err := scanner.LineBytes()
		if err != nil {
			if err == io.EOF {
				inProgressLine <- map[string]interface{}{"eof": true}
				break
			}
		}
		counter++
		if len(line) == 0 {
			continue
		}
		//wg1.Add(1)
		//semaphore <- struct{}{} // Acquire a token from the semaphore
		//go func(line []byte) {
		//	defer func() {
		//		wg1.Done()
		//		<-semaphore // Release the token back to the semaphore
		//	}()
		//	entry := SL.CreateLine(line, scanner.FileName(), mu)
		//	if entry != nil {
		//		inProgressLine <- entry
		//	}
		//}(line)
		entry := SL.CreateLine(line, scanner.FileName())
		if entry != nil {
			inProgressLine <- entry
		}
	}
	//wg1.Wait()
	if counter == params.DefultReadLine {
		inProgress <- scanner
	}
}
func (SL *SearchLogs) CreateLine(line []byte, fileName string) map[string]interface{} {
	entry := make(map[string]interface{})
	entry["fileName"] = fileName

	var err error

	//time.Sleep(time.Microsecond * 500)
	//fmt.Println(string(line))
	if json.Valid(line) {
		if err = json.Unmarshal(line, &entry); err != nil {
			fmt.Printf("Error unmarshaling JSON in file %s: %v\n", fileName, err)
			return nil
		}
		timeStr, ok := entry["time"].(string)
		if !ok {
			fmt.Println("Error converting time field to string")
			return nil
		}
		entry["time"], err = time.Parse("2006-01-02T15:04:05", timeStr)
		if err != nil {

			fmt.Printf("Error parsing time1 in file %s: %v\n", fileName, err)
			return nil
		}

	} else {
		isJson := true
		// Extract JSON content from the line
		jsonStart := bytes.IndexByte(line, '{')
		if jsonStart == -1 {
			isJson = false
			jsonStart = bytes.IndexByte(line, '[')
			if jsonStart != -1 {
				entry["text"] = string(line[:jsonStart])
			}
		}
		if jsonStart != -1 {
			if isJson {
				//fmt.Println(string(line[jsonStart:]))
				if err := json.Unmarshal(line[jsonStart:], &entry); err != nil {
					fmt.Printf("Error parsing time2 in file %s: %v\n", fileName, err)
					return nil
				}
				timeStr := string(line[:19])
				parsedTime, perr := time.Parse("2006/01/02 15:04:05", timeStr)
				if perr != nil {
					fmt.Printf("Error parsing time3 in file %s: %v\n", fileName, err)
					return nil
				}
				entry["time"] = parsedTime
			} else {
				timeStr := string(line[:19])
				parsedTime, perr := time.Parse("2006/01/02 15:04:05", timeStr)
				if perr != nil {
					fmt.Printf("Error parsing time4 in file %s: %v\n", fileName, err)
					return nil
				}
				entry["time"] = parsedTime

			}

		}
	}
	return entry
}

func (SL *SearchLogs) ProcessLine(inProgress chan *backscanner.Scanner, inProgressLine chan map[string]interface{}, isEnd *bool, SLReq params.SearchLogRequest) ([]string, int64) {
	//var full int
	var entryLines []map[string]interface{}
	var jsonData []string
	responseCounter := int64(0)
	c := 0
	counter := int64(0)

	for {
		entryLine := <-inProgressLine

		if _, ok := entryLine["eof"]; ok {
			c++
			if len(SL.ValidPath) == c {
				*isEnd = true
				ClearInProgress(inProgress)
				//close(inProgress)
				break
			}
			continue
		}
		entryLines = append(entryLines, entryLine)
		counter++
		if counter == 1000 {
			sort.Sort(ByTime(entryLines))
			counter = 0

		}

		//full++
		//fmt.Println(full)
		if SLReq.Limit == len(entryLines) && SLReq.Limit != 0 {
			//fmt.Println("break 2")
			ClearInProgress(inProgress)
			//close(inProgress)
			*isEnd = true
			break
		}
	}
	sort.Sort(ByTime(entryLines))
	// Process sorted entries
	for _, entry := range entryLines {
		fullContent := formatEntry(entry)
		jsonData = append(jsonData, fullContent)
		responseCounter++
	}
	return jsonData, responseCounter

}

func (SL *SearchLogs) ProcessLine2(inProgress chan *backscanner.Scanner, inProgressLine chan map[string]interface{}, isEnd *bool, SLReq params.SearchLogRequest) ([]string, int64) {
	//var full int
	var entryLines []map[string]interface{}
	var jsonData []string
	responseCounter := int64(0)
	c := 0
	counter := int64(0)

	for {

		entryLine := <-inProgressLine

		if _, ok := entryLine["eof"]; ok {
			c++
			if len(SL.ValidPath) == c {
				*isEnd = true
				ClearInProgress(inProgress)
				//close(inProgress)
				break
			}
			continue
		}

		fullContent := JustformatEntry(entryLine)
		teBool := FindWordInTexts2(SLReq, fullContent)
		if teBool {
			entryLines = append(entryLines, entryLine)
			counter++
			if counter == 1000 {
				//	sort.Sort(ByTime(entryLines))
				counter = 0

			}
			if SLReq.Limit == len(entryLines) && SLReq.Limit != 0 {
				//fmt.Println("break 2")
				ClearInProgress(inProgress)
				//close(inProgress)
				*isEnd = true
				break
			}
		}
		//full++
		//fmt.Println(full)

	}
	sort.Sort(ByTime(entryLines))
	// Process sorted entries
	for _, entry := range entryLines {
		fullContent := formatEntry(entry)
		jsonData = append(jsonData, fullContent)
		responseCounter++
	}
	return jsonData, responseCounter

}

func findWordInText(word string, text string) bool {
	return strings.Contains(text, word)
}

func convertValueToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case []interface{}:
		var strValues []string
		for _, item := range v {
			strValues = append(strValues, fmt.Sprint(item))
		}
		return fmt.Sprintf("[%s]", strings.Join(strValues, ", "))
	default:
		return fmt.Sprint(value)
	}
}

// formatEntry formats a log entry into JSON format
func formatEntry(entry map[string]interface{}) string {
	// Remove the "fileName" entry from the map temporarily
	fileName, ok := entry["fileName"].(string)
	if ok {
		delete(entry, "fileName")
	}

	// Marshal the entry map into JSON
	entryJSON, err := json.Marshal(entry)
	if err != nil {
		return fmt.Sprintf("Error formatting entry: %v", err)
	}

	// Concatenate the "fileName" entry with a newline character
	result := fmt.Sprintf("%s\n%s", fileName, entryJSON)

	return result
}
func JustformatEntry(entry map[string]interface{}) string {
	// Remove the "fileName" entry from the map temporarily
	fileName, _ := entry["fileName"].(string)

	// Marshal the entry map into JSON
	entryJSON, err := json.Marshal(entry)
	if err != nil {
		return fmt.Sprintf("Error formatting entry: %v", err)
	}

	// Concatenate the "fileName" entry with a newline character
	result := fmt.Sprintf("%s\n%s", fileName, entryJSON)

	return result
}
func removeElement(slice []string, element string) []string {
	var indexToRemove int
	for i, v := range slice {
		if v == element {
			indexToRemove = i
			break
		}
	}
	return append(slice[:indexToRemove], slice[indexToRemove+1:]...)
}
func ClearInProgress(inp chan *backscanner.Scanner) {
	for {
		select {
		case <-inp:
		case <-time.After(35 * time.Millisecond):

			return
		}
	}
}
func planB(Badline []byte) []byte {
	pain := 0
	for i, val := range Badline {
		if val == '\n' { // '\n' represents newline character
			pain = i
			break
		}
	}

	Goodline := append(Badline[pain+1:], Badline[:pain]...)
	return Goodline
}
func FindWordInTexts(SLReq params.SearchLogRequest, texts []string) ([]string, int64) {
	newTexts := make([]string, 0)
	responseCounter := int64(0)
	textsEx := texts
	fmt.Printf("SearchKey:%+v\n", SLReq.SearchKey)
	fmt.Printf("SearchKey:%+v\n", SLReq.NotIncludeSearchKey)
	fmt.Printf("NotIncludeFileName:%+v\n", SLReq.NotIncludeFileName)
	fmt.Printf("FileName:%+v\n", SLReq.FileName)

	for _, text := range textsEx {
		if SLReq.NotIncludeSearchKey != params.EmptyString {
			if findWordInText(SLReq.NotIncludeSearchKey, text) {
				continue
			}
		}

		if SLReq.SearchKey != params.EmptyString {
			if !findWordInText(SLReq.SearchKey, text) {
				continue
			}
		}

		responseCounter++
		newTexts = append(newTexts, text)
	}

	return newTexts, responseCounter

}
func FindWordInTexts2(SLReq params.SearchLogRequest, text string) bool {
	if SLReq.NotIncludeSearchKey != params.EmptyString {
		if findWordInText(SLReq.NotIncludeSearchKey, text) {
			return false
		}
	}

	if SLReq.SearchKey != params.EmptyString {
		if !findWordInText(SLReq.SearchKey, text) {
			return false
		}
	}
	return true

}
