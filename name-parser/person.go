package name_parser

import (
	"fmt"
	"regexp"
)

type Person map[string]string

const (
	nameRegexString = `[A-Z][a-z]*`

	PersonFirstNameKey  = "first-name"
	PersonMidNamePrefix = "mid-name-"
	PersonLastNameKey   = "last-name"
)

var (
	nameRegex = regexp.MustCompilePOSIX(nameRegexString)
)

func (p *Person) ParseFrom(s string) error {
	// find all person name parts
	parts := nameRegex.FindAllString(s, -1)
	if len(parts) < 3 {
		return fmt.Errorf("Cannot parse '%s' as a person's name", s)
	}

	// grab the first and last, and put them in special entries
	*p = make(map[string]string, len(parts)-1)
	(*p)[PersonFirstNameKey] = parts[0]
	(*p)[PersonLastNameKey] = parts[len(parts)-1]

	// read out the middle parts in order, into appropriately named entries
	parts = parts[1 : len(parts)-1]
	for i, part := range parts {
		(*p)[fmt.Sprintf("%s%d", PersonMidNamePrefix, i+1)] = part
	}

	return nil
}

func (p *Person) GetParts() map[string]string {
	return *p
}
