package system

import (
	"bufio"
	"io"
	"os"
	"strings"

	"k8s.io/klog/v2"
)

func OverlapWriteFile(fileName, fileData string) {
	dirPathSlice := strings.Split(fileName, "/")
	_ = os.MkdirAll(strings.Trim(fileName, dirPathSlice[len(dirPathSlice)-1]), 0755)
	f, fErr := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if fErr != nil {
		klog.Infoln(fErr)
	}
	_, wErr := f.WriteString(fileData)
	if wErr != nil {
		klog.Infoln(wErr)
	}
	_ = f.Sync()
	_ = f.Close()
}

func MergeFile(oldFile, newFile string) {
	fo, foErr := os.OpenFile(oldFile, os.O_RDWR|os.O_APPEND, 0777)
	if foErr != nil {
		klog.Infoln(foErr)
	}
	// open new file
	fn, fnErr := os.Open(newFile)
	if fnErr != nil {
		klog.Infoln(fnErr)
	}
	// read new file
	rd := bufio.NewReader(fn)
	for {
		oneDoc, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		klog.Infoln("doc info:", oneDoc)
		// write old file
		_, fwErr := fo.WriteString(oneDoc)
		if fwErr != nil {
			klog.Infoln(fwErr)
		}
	}
	// clone new file
	_ = fn.Close()
	// sync old file
	_ = fo.Sync()
	// close old file
	_ = fo.Close()
}

func DeleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		klog.Infoln(err)
	}
}

func DeleteDir(dirPath string) {
	err := os.RemoveAll(dirPath)
	if err != nil {
		klog.Infoln(err)
	}
}
