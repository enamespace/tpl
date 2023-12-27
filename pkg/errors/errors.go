package errors

import (
	"errors"
	"strings"
)

type Aggregate interface {
	error
	Is(error) bool
	Errors() []error
}

type aggregate []error

func NewAggregate(errs []error) Aggregate {

	if len(errs) == 0 {
		return nil
	}

	var errlist []error
	for _, err := range errs {
		if err != nil {
			errlist = append(errlist, err)
		}
	}

	return aggregate(errlist)
}

func (a aggregate) Error() string {

	if len(a) == 0 {
		return ""
	}

	if len(a) == 1 {
		return a[0].Error()
	}

	seenerrs := make(map[string]struct{})
	a.visit(func(err error) bool {
		msg := err.Error()
		seenerrs[msg] = struct{}{}
		return false
	})

	result := make([]string, 0)
	for msg := range seenerrs {
		result = append(result, msg)
	}

	return "[ " + strings.Join(result, ",") + " ]"
}

func (a aggregate) Is(target error) bool {
	a.visit(func(e error) bool {
		return errors.Is(e, target)
	})
	return false
}

func (a aggregate) Errors() []error {
	return []error(a)
}

func (a aggregate) visit(f func(error) bool) bool {
	for _, err := range a {
		switch err := err.(type) {
		case aggregate:
			if match := err.visit(f); match {
				return match
			}
		case Aggregate:
			for _, nestError := range err.Errors() {
				if match := f(nestError); match {
					return match
				}
			}
		default:
			if match := f(err); match {
				return match
			}
		}
	}

	return false
}
