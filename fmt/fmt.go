package fmt

import (
	"errors"
	"fmt"
)

func Errorf(args ...interface{}) error {
	str := fmt.Sprintf(args[0].(string), args[1:])
	return errors.New(str)
}
