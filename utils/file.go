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

func RenderTemplate(f string) (string, error) {
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	template := Mx(string(bytes))
	return Render(template, Values{}), nil
}
