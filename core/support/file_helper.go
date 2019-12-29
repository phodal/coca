package support

import (
	"fmt"
	"github.com/phodal/coca/config"
	"io/ioutil"
	"os"
)

var reporterPath = config.CocaConfig.ReporterPath

func WriteToCocaFile(fileName string, payload string) {
	if _, err := os.Stat(reporterPath); os.IsNotExist(err) {

		err := os.Mkdir(reporterPath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
	_ = ioutil.WriteFile(reporterPath+"/"+fileName, []byte(payload), os.ModePerm)
}

func ReadCocaFile(fileName string) []byte {
	return ReadFile(reporterPath + "/" + fileName)
}

func ReadFile(fileName string) []byte {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		_ = fmt.Errorf("Failed removing original file: %s", err)
		return nil
	}
	return contents
}
