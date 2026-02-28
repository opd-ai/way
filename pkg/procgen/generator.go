package procgen

// GenreID identifies one of the five supported genres.
type GenreID int

const (
	GenreFantasy        GenreID = iota
	GenreSciFi
	GenreHorror
	GenreCyberpunk
	GenrePostApocalyptic
)

// GenerationParams holds parameters for procedural content generation.
type GenerationParams struct {
	Genre      GenreID
	Difficulty float64
	Seed       int64
}

// Generator is the interface all procedural generators implement.
type Generator interface {
	Generate(seed int64, params GenerationParams) (interface{}, error)
	Validate() error
}
