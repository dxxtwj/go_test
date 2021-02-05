package har

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)
const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string  {
	return string(e)
}

func HandlerFileList (writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start" + "with" + prefix)
	}
	path := request.URL.Path[len(prefix):] // list/fib.txt
	file, err := os.Open(path)
	if err != nil {
		// 把错误打印在浏览器中给用户看（这样是不好的）
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		//return
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}
	writer.Write(all)
	return nil
}
