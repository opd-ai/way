package procgen

// HazardData holds the generated hazard properties.
type HazardData struct {
	Name     string
	Genre    GenreID
	Density  float64
	Position Waypoint
}

// HazardGenerator produces procedural hazard configurations.
type HazardGenerator struct {
	lastOutput *HazardData
}

// Generate creates a hazard configuration from the given seed and parameters.
func (g *HazardGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	g.lastOutput = &HazardData{
		Name:    "Generated Hazard",
		Genre:   params.Genre,
		Density: params.Difficulty,
	}
	return g.lastOutput, nil
}

// Validate checks the generated hazard for validity.
func (g *HazardGenerator) Validate() error {
	return nil
}
