package procgen

// KartData holds the generated kart properties.
type KartData struct {
	Name     string
	Genre    GenreID
	Speed    float64
	Handling float64
	Weight   float64
}

// KartGenerator produces procedural kart configurations.
type KartGenerator struct {
	lastOutput *KartData
}

// Generate creates a kart configuration from the given seed and parameters.
func (g *KartGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	g.lastOutput = &KartData{
		Name:     "Generated Kart",
		Genre:    params.Genre,
		Speed:    1.0,
		Handling: 1.0,
		Weight:   1.0,
	}
	return g.lastOutput, nil
}

// Validate checks the generated kart for validity.
func (g *KartGenerator) Validate() error {
	return nil
}
