package errors

import (
	"github.com/pkg/errors"
)

func New(msg string) error {
	return errors.New(msg)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

func Cause(err error) error {
	return errors.Cause(err)
}

type causer interface {
	Cause() error
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func StackTracer(err error) errors.StackTrace {
	var stack errors.StackTrace

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		if err, ok := err.(stackTracer); ok {
			stack = err.StackTrace()
		}
		err = cause.Cause()
	}

	return stack
}
