package cocatest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

// JSONBytesEqual compares the JSON in two byte slices.
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
		exceptStr, _ := json.Marshal(exceptInterface)
		actualStr, _ := json.Marshal(actualInterface)
		fmt.Println(string(actualStr))
		fmt.Println(string(exceptStr))
	}
	return isEqual, nil
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

//func JSONEqual(a, b io.Reader) (bool, error) {
//	var j, j2 interface{}
//	d := json.NewDecoder(a)
//	if err := d.Decode(&j); err != nil {
//		return false, err
//	}
//	d = json.NewDecoder(b)
//	if err := d.Decode(&j2); err != nil {
//		return false, err
//	}
//	return reflect.DeepEqual(j2, j), nil
//}
//
