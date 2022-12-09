package utils

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func LoadFile(path string, dest interface{}) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, dest)
	if err != nil {
		return err
	}
	return nil
}
