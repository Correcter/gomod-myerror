package myerror

import "errors"

func ReturnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function!")
		return err
	}
	return nil
}
