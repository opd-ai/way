// Package procgen provides the procedural content generation framework.
package procgen

// GenreID identifies one of the five supported genres.
type GenreID int

const (
	GenreFantasy       GenreID = iota // Fantasy genre
	GenreSciFi                        // Sci-fi genre
	GenreHorror                       // Horror genre
	GenreCyberpunk                    // Cyberpunk genre
	GenrePostApocalyptic              // Post-apocalyptic genre
)

// GenerationParams holds parameters for procedural generation.
type GenerationParams struct {
	Genre      GenreID
	Difficulty float64
	Seed       int64
}

// Generator is the interface all procedural generators implement.
type Generator interface {
	// Generate produces content from a seed and parameters.
	Generate(seed int64, params GenerationParams) (interface{}, error)
	// Validate checks that the generated content is valid.
	Validate() error
}
