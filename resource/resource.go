package resource

// Error type for resource calls
type Error int

// Types of errors for resource calls
const (
	ErrorNone Error = iota
	ErrorTooManyResults
	ErrorCharacterNotFound
)

type Resource interface {
	CharacterSearch(name string) (*Character, Error)
}
