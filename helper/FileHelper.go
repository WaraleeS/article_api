package helper

import (
	"io/ioutil"
)

func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err.Error(), nil
	} else {
		return string(bytes), nil
	}
}
