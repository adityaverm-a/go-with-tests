package maps

const (
	ErrWordExists = DictionaryErr("can not add word because it already exisits")
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

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

	case nil:
		return ErrWordExists

	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) {

}
