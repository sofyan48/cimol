package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sofyan48/cimol/src/util/helper/logging"
)

// Storage ...
type Storage struct {
	Logs logging.LogInterface
}

// StorageHandler ..
func StorageHandler() *Storage {
	return &Storage{
		Logs: logging.LogHandler(),
	}
}

// StorageInterface ...
type StorageInterface interface {
	CreateFolder(dir string) bool
	CreateFolderTree(dir string) bool
	CreateFile(path, fileName string)
	ReadDir(path string) []string
	RemoveFile(path string) bool
	RemoveContents(dir string) error
	CreateJSONFile(data interface{}, path, fileName string) error
}

// CreateFolderTree ...
func (file *Storage) CreateFolderTree(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			file.Logs.Write("Storage", err.Error())
			return false
		}
	}
	return true
}

// CreateFolder ...
func (file *Storage) CreateFolder(dir string) bool {
	err := os.Mkdir(dir, 0755)
	if err != nil {
		file.Logs.Write("Storage", err.Error())
		return false
	}
	return true
}

// IsDirectory ...
func (file *Storage) IsDirectory(path string) bool {
	fd, err := os.Stat(path)
	if err != nil {
		file.Logs.Write("Storage", err.Error())
	}
	switch mode := fd.Mode(); {
	case mode.IsDir():
		return true
	case mode.IsRegular():
		return false
	}
	return false
}

// ReadDir ...
func (file *Storage) ReadDir(path string) []string {
	fileList := []string{}
	filepath.Walk(path, func(filePath string, f os.FileInfo, err error) error {
		if !file.IsDirectory(path) {
			return nil
		}
		fileList = append(fileList, filePath)
		return nil
	})
	result := []string{}
	for _, i := range fileList {
		if path != i {
			result = append(result, i)
		}
	}
	return result
}

// RemoveDir ...
func (file *Storage) RemoveDir(path string) {

}

// CreateFile ...
func (file *Storage) CreateFile(path, fileName string) {

}

// CreateJSONFile ...
func (file *Storage) CreateJSONFile(data interface{}, path, fileName string) error {
	json, _ := json.MarshalIndent(data, "", " ")
	err := ioutil.WriteFile(path+"/"+fileName+".json", json, 0644)
	return err
}

// ReadFile ...
func (file *Storage) ReadFile(path, fileName string) {

}

// RemoveContents ...
func (file *Storage) RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveFile ...
func (file *Storage) RemoveFile(path string) bool {
	err := os.Remove(path)
	if err != nil {
		return false
	}
	return true
}

// CreateMetricFolder ...
func (file *Storage) CreateMetricFolder() {
	file.CreateFolder("./metric/sms")
	file.CreateFolder("./metric/push")
	file.CreateFolder("./metric/email")
}
