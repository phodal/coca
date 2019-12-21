package support

import (
	"github.com/phodal/coca/config"
	"fmt"
	"io"
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

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
