package storage

import (
	"errors"
)

var ErrUnableToFetch = errors.New("unable to fetch record")
var ErrUnableToUpdate = errors.New("unable to update record")
var ErrUnableToInsert = errors.New("unable to insert into database")
