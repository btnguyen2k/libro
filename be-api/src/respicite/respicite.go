package respicite

import (
	"errors"
	"time"
)

const (
	// MappingFieldSrc is name of BO field to hold 'source' value.
	MappingFieldSrc = "src"

	// MappingFieldDest is name of BO field to hold 'destination' value.
	MappingFieldDest = "dest"

	// MappingFieldCreatedOn is name of BO field to hold 'created-on' value.
	MappingFieldCreatedOn = "time"
)

// Mapping represents a mapping {source -> destination}
type Mapping struct {
	src       string    `json:"src"`  // source
	dest      string    `json:"dest"` // destination
	createdOn time.Time `json:"time"` // timestamp when this mapping is created
}

var (
	// ErrNotFound is returned if the mapping does not exist
	ErrNotFound = errors.New("mapping not found")

	// ErrDuplicated is returned if the mapping already existed
	ErrDuplicated = errors.New("mapping already existed")
)
