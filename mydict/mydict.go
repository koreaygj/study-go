package mydict

import "errors"

type Dictionary map[string]string

var (
	errNotFount = errors.New("Not Found")
	errWordExists = errors.New("Word already exists")
	errCantUpdate = errors.New("Cant update non-existing word")
)
func(d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if( exists){
		return value, nil
	}
	return "", errNotFount
}

func(d Dictionary) Add(word, def string) error{
	_, err := d.Search(word)
	if(err == errNotFount){
		d[word] = def
	} else if(err == nil){
		return errWordExists
	}
	return nil
}

func(d Dictionary) Update(word, def string) error{
	_, err := d.Search(word)
	if(err == errNotFount){
		return errCantUpdate
	}else if(err == nil){
		d[word] = def
	}
	return nil
}