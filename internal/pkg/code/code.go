package code

import (
	"github.com/pkg/errors"
	"fmt"
)


type Coder interface {
	HTTPStatus() int
	String() string
	Reference() string
	Code() string
}

type defaultCoder struct {
	C int
	HTTP int
	Ext string
	Ref string
}


func WithCode(err error, c int) error {
	return errors.Wrap(err, fmt.Sprintf("with code %v ", c))
} 

