package helpers

import (
	"fmt"
	"strings"
)

func CreateErrorMessage(err error) error {
	errMessage := err.Error()

	err = fmt.Errorf(strings.Replace(errMessage, "rpc error: code = Unknown desc = ", "", -1))

	return err
}
