package toolkit

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

// PostForm simulate the front-end initiating a form request;
// @param  valMap map[string]string  stores the form field name and form value.
// @param  fileMap map[string]string  stores the form field name and file path.
func PostForm(postUrl string, valMap map[string]string, fileMap map[string]string) (httpStatusCode int, resBody string, err error) {
	//new 客户端
	client := http.Client{}
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)

	//fileMap -> form上传文件
	for fieldName, fPath := range fileMap {
		file, err := os.Open(fPath)
		defer file.Close()
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("Error while open file %s", fPath))
			return -1, "", err
		}

		// !!!关键代码是这里
		fileWrite, err := bodyWrite.CreateFormFile(fieldName, fPath)
		if err != nil {
			err = errors.Wrap(err, "error while CreateFormFile")
			return -1, "", err
		}
		_, err = io.Copy(fileWrite, file)
		if err != nil {
			err = errors.Wrap(err, "error while io.Copy")
			return -1, "", err
		}
	}

	//valMap -> 写入form
	for fieldName, value := range valMap {
		fieldWriter, err := bodyWrite.CreateFormField(fieldName)
		if err != nil {
			err = errors.Wrap(err, "error while CreateFormField")
			return -1, "", err
		}
		_, err = fieldWriter.Write([]byte(value))
		if err != nil {
			err = errors.Wrap(err, "error while fieldWriter.Write")
			return -1, "", err
		}

	}

	bodyWrite.Close() //要关闭，会将w.w.boundary刷写到w.writer中

	// 创建请求
	req, err := http.NewRequest(http.MethodPost, postUrl, bodyBuf)
	if err != nil {
		err = errors.Wrap(err, "error while http.NewRequest")
		return -1, "", err
	}
	// 设置头
	contentType := bodyWrite.FormDataContentType()
	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "error while client.Do")
		return -1, "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "error while ioutil.ReadAll(resp.Body)")
		return resp.StatusCode, "", err
	}
	return resp.StatusCode, string(b), nil
}

// get body from url
// if timeout==0,it means no timeout limit
func Get(getUrl string, timeout time.Duration) (httpStatusCode int, resBody string, err error) {
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(getUrl)
	if err != nil {
		err = errors.Wrap(err, "error while client.Get()")
		return -1, "", err
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "error while io.ReadAll(resp.Body)")
		return resp.StatusCode, "", err
	}

	return resp.StatusCode, string(bs), nil
}
