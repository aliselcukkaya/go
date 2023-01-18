package maps

type Dictionary map[string]string

const (
	ErrNotFound      = DictionaryError("could not find the word you were looking for")
	ErrAlreadyExists = DictionaryError("cannot add word because it already exists")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}
}
