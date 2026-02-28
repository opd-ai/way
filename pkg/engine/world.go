// Package engine provides the ECS core: entity management, component stores,
// system registration, and the fixed-timestep tick loop.
package engine

// Entity is an integer ID representing a game object.
type Entity int

// World manages entities, component stores, and systems.
type World struct {
	nextEntity  Entity
	entities    map[Entity]bool
	transforms  map[Entity]*TransformComponent
	vehicles    map[Entity]*VehiclePhysicsComponent
	items       map[Entity]*ItemHolderComponent
	positions   map[Entity]*RacePositionComponent
	systems     []System
}

// NewWorld creates and returns an empty World.
func NewWorld() *World {
	return &World{
		nextEntity: 1,
		entities:   make(map[Entity]bool),
		transforms: make(map[Entity]*TransformComponent),
		vehicles:   make(map[Entity]*VehiclePhysicsComponent),
		items:      make(map[Entity]*ItemHolderComponent),
		positions:  make(map[Entity]*RacePositionComponent),
	}
}

// CreateEntity allocates a new entity ID and returns it.
func (w *World) CreateEntity() Entity {
	e := w.nextEntity
	w.nextEntity++
	w.entities[e] = true
	return e
}

// DestroyEntity removes an entity and all its components.
func (w *World) DestroyEntity(e Entity) {
	delete(w.entities, e)
	delete(w.transforms, e)
	delete(w.vehicles, e)
	delete(w.items, e)
	delete(w.positions, e)
}

// AddTransform attaches a TransformComponent to an entity.
func (w *World) AddTransform(e Entity, c *TransformComponent) {
	w.transforms[e] = c
}

// GetTransform returns the TransformComponent for an entity, or nil.
func (w *World) GetTransform(e Entity) *TransformComponent {
	return w.transforms[e]
}

// AddVehicle attaches a VehiclePhysicsComponent to an entity.
func (w *World) AddVehicle(e Entity, c *VehiclePhysicsComponent) {
	w.vehicles[e] = c
}

// GetVehicle returns the VehiclePhysicsComponent for an entity, or nil.
func (w *World) GetVehicle(e Entity) *VehiclePhysicsComponent {
	return w.vehicles[e]
}

// AddItemHolder attaches an ItemHolderComponent to an entity.
func (w *World) AddItemHolder(e Entity, c *ItemHolderComponent) {
	w.items[e] = c
}

// GetItemHolder returns the ItemHolderComponent for an entity, or nil.
func (w *World) GetItemHolder(e Entity) *ItemHolderComponent {
	return w.items[e]
}

// AddRacePosition attaches a RacePositionComponent to an entity.
func (w *World) AddRacePosition(e Entity, c *RacePositionComponent) {
	w.positions[e] = c
}

// GetRacePosition returns the RacePositionComponent for an entity, or nil.
func (w *World) GetRacePosition(e Entity) *RacePositionComponent {
	return w.positions[e]
}

// RegisterSystem adds a system to the world's update loop.
func (w *World) RegisterSystem(s System) {
	w.systems = append(w.systems, s)
}

// Tick advances the world by one fixed timestep, updating all registered systems.
func (w *World) Tick(dt float64) {
	for _, s := range w.systems {
		s.Update(w, dt)
	}
}

// Entities returns all live entity IDs.
func (w *World) Entities() []Entity {
	out := make([]Entity, 0, len(w.entities))
	for e := range w.entities {
		out = append(out, e)
	}
	return out
}
