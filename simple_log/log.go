package simple_log

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
)

func Init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime)
}

func ErrLog(errInfo ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	log.Println(file, line, "[ERROR]", errInfo)
}

func InfoLog(infoMsg ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	// funcName := runtime.FuncForPC(ptrFuncName).Name()
	log.Println(file, line, "[INFO]", infoMsg)
}

func InfoLogConvertToJson(src interface{}) {
	js, _ := json.Marshal(src)
	_, file, line, _ := runtime.Caller(1)
	log.Println(file, line, "[INFO]", string(js))
}
