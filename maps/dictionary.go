package main

// Interesting property of maps: You can modify them without passing as an address to it e.g. &myMap
// Attempts to write to a nil map will cause a runtime panic. Never initialise an empty map variable
// var m map[string]string BAD

// var dictionary = map[string]string{} GOOD
// var dictionary = make(map[string]string) GOOD

//not used, deprecated
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does nto exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string { // By implementing the Error() method, this custom Err type automatically satisfies the error interface
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

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition

	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word) // built in map method that removes the key from the map
}
