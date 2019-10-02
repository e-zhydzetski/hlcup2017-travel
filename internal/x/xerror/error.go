package xerror

import "github.com/pkg/errors"

type Error string

// for const errors: const ErrNotFound = Error("not found")
func (e Error) Error() string {
	return string(e)
}

func Combine(err, nextErr error) error {
	if err == nil {
		err = nextErr
	} else if nextErr != nil {
		err = errors.Wrap(err, nextErr.Error())
	}
	return err
}
