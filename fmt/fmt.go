package fmt

import "errors"

func Errorf(args ...interface{}) error {
	return errors.New("yolo")
}
