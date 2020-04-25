package stringsvc

import "strings"

// Svc defines the shape of StringSvc
type Svc interface {
	toUppercase(string) string
}

type stringSvc struct {
}

// New returns a new string service
func New() Svc {
	return &stringSvc{}
}

func (s stringSvc) toUppercase(str string) string {
	return strings.ToUpper(str)
}
