package golibrary

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// GetBodyBytes 获取返回包的body, bytes形式
func GetBodyBytes(r *http.Response) ([]byte, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = r.Body.Close() //  must close
	if err != nil {
		return nil, err
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes, nil
}

// GetBodyString 获取返回包的body, string形式, 如果是读取图片可能会出现乱码
func GetBodyString(r *http.Response) (string, error) {
	bodyBytes, err := GetBodyBytes(r)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}

// GetTitle 获取返回包的title, 可能会匹配多组, 以数组形式返回
func GetTitle(r *http.Response) ([]string, error) {
	body, err := GetBodyString(r)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile("<title>[\\s\\S]*?</title>")
	match := re.FindStringSubmatch(body)
	return match, err
}

// GetHeaderRaw 获取返回包的header, string形式
func GetHeaderRaw(r *http.Response) string {
	headerRaw := ""
	for key := range r.Header {
		headerRaw += key + ": " + strings.Join(r.Header[key], ";") + "\r\n"
	}
	return headerRaw
}

// GetBodyMd5 计算返回包body的md5值, 返回string
func GetBodyMd5(r *http.Response) (string, error) {
	bodyBytes, err := GetBodyBytes(r)
	if err != nil {
		return "", err
	}
	digest := fmt.Sprintf("%x", md5.Sum(bodyBytes))
	return digest, nil
}
