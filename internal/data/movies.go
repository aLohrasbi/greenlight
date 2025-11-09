package data

import (
	"time"
)

// It’s crucial to point out here that all the fields in our Movie struct are exported
// (i.e. start with a capital letter),
// which is necessary for them to be visible to Go’s encoding/json package.
// Any fields which aren’t exported won’t be included when encoding a struct to JSON
// Annotate the Movie struct with struct tags to control how the keys appear in the
// JSON-encoded output.
type Movie struct {
    ID        int64     `json:"id"`
    CreatedAt time.Time `json:"-"` // Use the - directive  never want a particular struct field to appear in the JSON output.
    Title     string    `json:"title"`
    Year      int32     `json:"year,omitzero"`    // Add the omitzero directive
	 // Use the Runtime type instead of int32. Note that the omitzero directive will
    // still work on this: if the Runtime field has the underlying value 0, then it will
    // be considered zero and omitted -- and the MarshalJSON() method we just made 
    // won't be called at all.
    Runtime Runtime  `json:"runtime,omitzero"` // Add the omitzero directive and not change the key name
    Genres    []string  `json:"genres,omitempty"`  // Add the omitzero directive
    Version   int32     `json:"version,string"`
}
