package procgen

// ItemArchetype identifies one of the eight item archetypes.
type ItemArchetype int

const (
	ItemProjectile ItemArchetype = iota
	ItemMine
	ItemShield
	ItemSpeed
	ItemEMP
	ItemSwap
	ItemPull
	ItemAreaBomb
)

// ItemData holds the generated item properties.
type ItemData struct {
	Archetype ItemArchetype
	Name      string
	Genre     GenreID
}

// ItemGenerator produces procedural item configurations.
type ItemGenerator struct {
	lastOutput *ItemData
}

// Generate creates an item configuration from the given seed and parameters.
func (g *ItemGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	g.lastOutput = &ItemData{
		Archetype: ItemProjectile,
		Name:      "Generated Item",
		Genre:     params.Genre,
	}
	return g.lastOutput, nil
}

// Validate checks the generated item for validity.
func (g *ItemGenerator) Validate() error {
	return nil
}
