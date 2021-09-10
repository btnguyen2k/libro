package respicite

import (
	"errors"
	"time"
)

const (
	// MappingFieldSrc is name of BO field to hold 'source' value.
	MappingFieldSrc = "Src"

	// MappingFieldDest is name of BO field to hold 'destination' value.
	MappingFieldDest = "Dest"

	// MappingFieldCreatedOn is name of BO field to hold 'created-on' value.
	MappingFieldCreatedOn = "time"
)

// Mapping represents a mapping {source -> destination}
type Mapping struct {
	Src       string    `json:"Src"`  // source
	Dest      string    `json:"Dest"` // destination
	CreatedOn time.Time `json:"time"` // timestamp when this mapping is created
}

var (
	// ErrNotFound is returned if the mapping does not exist
	ErrNotFound = errors.New("mapping not found")

	// ErrDuplicated is returned if the mapping already existed
	ErrDuplicated = errors.New("mapping already existed")
)
