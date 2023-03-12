package utils

import (
	"io/ioutil"
)

func ReadTemplate(f string) (*Tpl, error) {
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return Mx(string(bytes)), nil
}
