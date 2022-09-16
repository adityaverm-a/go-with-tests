package maps

const (
	ErrWordExists        = DictionaryErr("can not add word because it already exisits")
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordDoesNotExists = DictionaryErr("can not update word because it does not exist")
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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists

	case nil:
		d[word] = definition

	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

// Quick Summary

// Throughout the process I learned How to:

// - Create maps
// - Search for items in maps
// - Add new items to maps
// - Update items in maps
// - Delete items from a map

// Learned more about Errors:
// - How to create errors that are constants
// - Writing error wrappers
