package resource

// Error type for resource calls
type Error int

// Types of errors for resource calls
const (
	ErrorNone Error = iota
	ErrorTooManyResults
	ErrorCharacterNotFound
)

// Resource API for searching character
type Resource interface {
	CharacterSearch(name string) (*Character, Error)
}
