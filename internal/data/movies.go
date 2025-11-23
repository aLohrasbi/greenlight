package data

import (
	"time"

	"github.com/alohrasbi/greenlight/internal/data/validator"
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
	Year      int32     `json:"year,omitzero"` // Add the omitzero directive
	// Use the Runtime type instead of int32. Note that the omitzero directive will
	// still work on this: if the Runtime field has the underlying value 0, then it will
	// be considered zero and omitted -- and the MarshalJSON() method we just made
	// won't be called at all.
	Runtime Runtime  `json:"runtime,omitzero"` // Add the omitzero directive and not change the key name
	Genres  []string `json:"genres,omitempty"` // Add the omitzero directive
	Version int32    `json:"version,string"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	// Use the Check() method to execute our validation checks. This will add the
	// provided key and error message to the errors map if the check does not evaluate
	// to true. For example, in the first line here we "check that the title is not
	// equal to the empty string". In the second, we "check that the length of the title
	// is less than or equal to 500 bytes" and so on.
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	// Note that we're using the Unique helper in the line below to check that all
	// values in the input.Genres slice are unique.
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
