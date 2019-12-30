package coca_file

import (
	"fmt"
	"github.com/phodal/coca/cmd/config"
	"io/ioutil"
	"os"
	"path/filepath"
)

var reporterPath = config.CocaConfig.ReporterPath

func WriteToCocaFile(fileName string, payload string) {
	if _, err := os.Stat(reporterPath); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(reporterPath, os.ModePerm)
		if mkdirErr != nil {
			fmt.Println(mkdirErr)
		}
	}
	_ = ioutil.WriteFile(filepath.FromSlash(reporterPath+"/"+fileName), []byte(payload), os.ModePerm)
}

func ReadCocaFile(fileName string) []byte {
	return ReadFile(filepath.FromSlash(reporterPath + "/" + fileName))
}

func ReadFile(fileName string) []byte {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		_ = fmt.Errorf("Failed removing original file: %s", err)
		return nil
	}
	return contents
}
