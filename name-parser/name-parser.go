package name_parser

type NameParser interface {
	ParseFrom(string) error
	GetParts() map[string]string
}
