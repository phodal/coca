package cocatest

import (
	"encoding/json"
	"fmt"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"

	"io/ioutil"
	"reflect"
)

func JSONBytesEqual(actual, except []byte) (bool, error) {
	var actualInterface, exceptInterface interface{}
	if err := json.Unmarshal(actual, &actualInterface); err != nil {
		return false, err
	}
	if err := json.Unmarshal(except, &exceptInterface); err != nil {
		return false, err
	}
	isEqual := reflect.DeepEqual(exceptInterface, actualInterface)
	if !isEqual {
		if string(except) == "{}" {
			actualStr, _ := json.MarshalIndent(actualInterface, "", "\t")
			fmt.Println(string(actualStr))
		} else {
			formatNotEqualPrint(exceptInterface, actualInterface)
		}
	}
	return isEqual, nil
}

func formatNotEqualPrint(exceptInterface interface{}, actualInterface interface{}) {
	exceptStr, _ := json.Marshal(exceptInterface)
	actualStr, _ := json.Marshal(actualInterface)

	differ := diff.New()
	diffResult, _ := differ.Compare(exceptStr, actualStr)
	config := formatter.AsciiFormatterConfig{
		ShowArrayIndex: true,
		Coloring:       true,
	}
	var aJson map[string]interface{}
	_ = json.Unmarshal(actualStr, &aJson)

	aFormatter := formatter.NewAsciiFormatter(aJson, config)
	diffString, _ := aFormatter.Format(diffResult)
	fmt.Println(diffString)
}

func JSONFileBytesEqual(actualInterface interface{}, exceptFile string) (bool, error) {
	actual, err := json.MarshalIndent(actualInterface, "", "\t")
	if err != nil {
		return false, err
	}

	contents, err := ioutil.ReadFile(exceptFile)
	if err != nil {
		return false, err
	}

	return JSONBytesEqual(actual, contents)
}
