package match_name

import (
	"github.com/uberblah/go-demo/name-parser"
)

var (
	registry = map[string]name_parser.NameParser{
		"person": &name_parser.Person{},
		"email":  &name_parser.Email{},
	}
)

func MatchName(name string) (nameType string, parts map[string]string) {
	var parser name_parser.NameParser
	for nameType, parser = range registry {
		if parser.ParseFrom(name) == nil {
			parts = parser.GetParts()
			return
		}
	}
	return "", nil
}
