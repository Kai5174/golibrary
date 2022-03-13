package golibrary

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"path"
	"reflect"
	"testing"
)

func TestReadFileLines(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"1", args{filePath: path.Join("tests", "test1.txt")}, []string{"1", "11", "111"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileLines(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFileLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalkDir(t *testing.T) {
	type args struct {
		folderPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"1", args{folderPath: path.Join("tests", "walking")}, []string{
			path.Join("tests", "walking", "123.aaa"),
			path.Join("tests", "walking", "123.asdf"),
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WalkDir(tt.args.folderPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("WalkDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WalkDir() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseYamlToStruct(t *testing.T) {
	type args struct {
		yamlPath string
		s        interface{}
	}
	type yamlStruct struct {
		A string   `yaml:"a"`
		B string   `yaml:"b"`
		C []string `yaml:"c"`
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", args{
			yamlPath: path.Join("tests", "test3.yaml"),
			s:        &yamlStruct{},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ParseYamlToStruct(tt.args.yamlPath, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("ParseYamlToStruct() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				a := tt.args.s.(*yamlStruct)
				assert.Equal(t, a.A, "1")
				assert.Equal(t, a.B, "2")
				assert.Equal(t, a.C, []string{"3", "4"})
			}
		})
	}
}

func TestReadFileString(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{"1", args{filePath: path.Join("tests", "test2.txt")}, "111111", assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileString(tt.args.filePath)
			if !tt.wantErr(t, err, fmt.Sprintf("ReadFileString(%v)", tt.args.filePath)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ReadFileString(%v)", tt.args.filePath)
		})
	}
}
