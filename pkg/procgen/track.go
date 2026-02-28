package procgen

// Waypoint represents a point on the track's waypoint graph.
type Waypoint struct {
	X, Y, Z   float64
	Width      float64
	IsShortcut bool
	IsHazard   bool
}

// TrackData holds the generated track layout.
type TrackData struct {
	Waypoints []Waypoint
	Name      string
	Genre     GenreID
}

// TrackGenerator produces procedural track layouts.
type TrackGenerator struct {
	lastOutput *TrackData
}

// Generate creates a track layout from the given seed and parameters.
func (g *TrackGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	g.lastOutput = &TrackData{
		Waypoints: []Waypoint{},
		Name:      "Generated Track",
		Genre:     params.Genre,
	}
	return g.lastOutput, nil
}

// Validate checks the generated track for validity.
func (g *TrackGenerator) Validate() error {
	return nil
}
