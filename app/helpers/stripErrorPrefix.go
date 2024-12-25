package helpers

import (
	"fmt"
	"strings"
)

func StripGRPCErrorPrefix(err error) error {
	if err == nil {
		return nil
	}
	const grpcPrefix = "rpc error: code = Unknown desc = "
	return fmt.Errorf(strings.TrimPrefix(err.Error(), grpcPrefix))
}
