package golibrary

import (
	"bufio"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadFile 读取文件内容,返回bytes
func ReadFile(filePath string) ([]byte, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// WriteFile 将bytes写入目标文件
func WriteFile(filePath string, content []byte) error {
	return os.WriteFile(filePath, content, 0644)
}

// WriteFileString 将string写入目标文件
func WriteFileString(filePath string, content string) error {
	contentBytes := []byte(content)
	return os.WriteFile(filePath, contentBytes, 0644)
}

// ReadFileString 读取文件内容,返回string
func ReadFileString(filePath string) (string, error) {
	content, err := ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// ParseYamlToStruct 将yaml文件解析成结构体
func ParseYamlToStruct(yamlPath string, s *struct{}) error {
	data, err := ReadFile(yamlPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, s)
	if err != nil {
		return err
	}
	return nil
}

// WalkDir 获取指定目录下的所有文件
func WalkDir(folderPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// ReadFileLines 读取文件,每行分割
func ReadFileLines(filePath string) ([]string, error) {
	file, err := os.Open("test.txt")

	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
