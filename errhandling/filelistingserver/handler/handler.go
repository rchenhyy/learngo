package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

const prefix = "/list/"

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path
	if strings.Index(path, prefix) != 0 {
		return userError(fmt.Sprintf("Path must start with: %s", prefix))
	}

	filename := path[len(prefix):]

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(content)
	return nil
}
