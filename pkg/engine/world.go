package engine

// Entity is an integer ID representing a game entity.
type Entity int

// System defines the interface all ECS systems must implement.
type System interface {
	Update(w *World, dt float64)
}

// World holds all entities, component stores, and registered systems.
type World struct {
	nextEntity  Entity
	Entities    []Entity
	Transforms  map[Entity]*TransformComponent
	Vehicles    map[Entity]*VehiclePhysicsComponent
	Items       map[Entity]*ItemHolderComponent
	Positions   map[Entity]*RacePositionComponent
	systems     []System
	TickRate    float64
	accumulator float64
}

// NewWorld creates a new ECS world with a 60 Hz fixed timestep.
func NewWorld() *World {
	return &World{
		nextEntity: 1,
		Entities:   make([]Entity, 0),
		Transforms: make(map[Entity]*TransformComponent),
		Vehicles:   make(map[Entity]*VehiclePhysicsComponent),
		Items:      make(map[Entity]*ItemHolderComponent),
		Positions:  make(map[Entity]*RacePositionComponent),
		systems:    make([]System, 0),
		TickRate:   1.0 / 60.0,
	}
}

// CreateEntity allocates a new entity ID and returns it.
func (w *World) CreateEntity() Entity {
	e := w.nextEntity
	w.nextEntity++
	w.Entities = append(w.Entities, e)
	return e
}

// DeleteEntity removes an entity and all its components.
func (w *World) DeleteEntity(e Entity) {
	for i, ent := range w.Entities {
		if ent == e {
			w.Entities = append(w.Entities[:i], w.Entities[i+1:]...)
			break
		}
	}
	delete(w.Transforms, e)
	delete(w.Vehicles, e)
	delete(w.Items, e)
	delete(w.Positions, e)
}

// AddSystem registers a system for the tick loop.
func (w *World) AddSystem(s System) {
	w.systems = append(w.systems, s)
}

// Tick advances the world by dt seconds using fixed-timestep accumulation.
func (w *World) Tick(dt float64) {
	w.accumulator += dt
	for w.accumulator >= w.TickRate {
		for _, s := range w.systems {
			s.Update(w, w.TickRate)
		}
		w.accumulator -= w.TickRate
	}
}
