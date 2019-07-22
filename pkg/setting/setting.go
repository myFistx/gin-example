package setting

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

//var (
//	config *ini.File
//
//	RunMode string
//	HttpPort int
//	ReadTimeout time.Duration
//	WriteTimeout time.Duration
//	PageSize int
//	JwtSecret string
//)

func Init() {
	path := os.Getenv("GOPATH")
	fmt.Println(path)

	//config, err := ini.Load("F:/www/golang/src/gin-blog/conf/app.ini")
	//if err != nil {
	//	log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	//}
	//fmt.Println(config)
}

func CurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	return file
}
