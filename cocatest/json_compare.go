package cocatest

import (
	"encoding/json"
	"fmt"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"io/ioutil"
	"reflect"
)

func JSONBytesEqual(actual, except []byte, exceptFile string) (bool, error) {
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
			actualStr, _ := json.MarshalIndent(actualInterface, "", "  ")
			fmt.Println(string(actualStr))
			ioutil.WriteFile(exceptFile, actualStr, 0644)
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
	fmt.Println(string(actualStr))
}

func JSONFileBytesEqual(actualInterface interface{}, exceptFile string) bool {
	actual, err := json.MarshalIndent(actualInterface, "", "  ")
	if err != nil {
		fmt.Println(err)
		return false
	}

	contents, err := ioutil.ReadFile(exceptFile)
	if err != nil {
		fmt.Println(err)
		_ = ioutil.WriteFile(exceptFile, []byte(`{}`), 0644)
		return false
	}

	equal, err := JSONBytesEqual(actual, contents, exceptFile)
	if err !=nil {
		fmt.Println(err)
		return false
	}

	return equal
}
