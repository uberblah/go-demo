package name_parser

import (
	"fmt"
	"regexp"
)

type Email map[string]string

const (
	emailRegexString = `([[:alnum:]\._-]+)@([[:alnum:].]+)`

	EmailUserNameKey   = "username"
	EmailDomainNameKey = "domain-name"
)

var (
	emailRegex = regexp.MustCompilePOSIX(emailRegexString)
)

func (e *Email) ParseFrom(s string) error {
	parts := emailRegex.FindStringSubmatch(s)
	if len(parts) != 3 {
		return fmt.Errorf("Could not parse '%s' as an email", s)
	}

	*e = make(map[string]string, 2)
	(*e)[EmailUserNameKey] = parts[1]
	(*e)[EmailDomainNameKey] = parts[2]

	return nil
}

func (e *Email) GetParts() map[string]string {
	return *e
}
