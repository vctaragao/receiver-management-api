package storage

import (
	"errors"
)

var ErrUnableToFetch = errors.New("unable to fetch record")
var ErrUnableToUpdate = errors.New("unable to update record")
var ErrUnableToFind = errors.New("unable to find database records")
var ErrUnableToInsert = errors.New("unable to insert into database")
var ErrUnableToDelete = errors.New("unable to delete database records")
