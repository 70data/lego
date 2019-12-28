package os

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// OverlapWriteFile is overlap write data to file once
func OverlapWriteFile(fileName, fileData string) {
	dirPathSlice := strings.Split(fileName, "/")
	_ = os.MkdirAll(strings.Trim(fileName, dirPathSlice[len(dirPathSlice)-1]), 0755)
	f, fErr := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if fErr != nil {
		log.Println(fErr)
	}
	_, wErr := f.WriteString(fileData)
	if wErr != nil {
		log.Println(wErr)
	}
	_ = f.Sync()
	_ = f.Close()
}

// MergeFile is merge file
func MergeFile(oldFile, newFile string) {
	fo, foErr := os.OpenFile(oldFile, os.O_RDWR|os.O_APPEND, 0777)
	if foErr != nil {
		log.Println(foErr)
	}
	// open new file
	fn, fnErr := os.Open(newFile)
	if fnErr != nil {
		log.Println(fnErr)
	}
	// read new file
	rd := bufio.NewReader(fn)
	for {
		oneDoc, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		log.Println("doc info:", oneDoc)
		// write old file
		_, fwErr := fo.WriteString(oneDoc)
		if fwErr != nil {
			log.Println(fwErr)
		}
	}
	// clone new file
	_ = fn.Close()
	// sync old file
	_ = fo.Sync()
	// close old file
	_ = fo.Close()
}

// DeleteFile is delete file
func DeleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		log.Println(err)
	}
}

// DeleteDir is delete dir
func DeleteDir(dirPath string) {
	err := os.RemoveAll(dirPath)
	if err != nil {
		log.Println(err)
	}
}
