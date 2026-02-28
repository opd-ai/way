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
	Cameras     map[Entity]*CameraComponent
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
		Cameras:    make(map[Entity]*CameraComponent),
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
	delete(w.Cameras, e)
}

// AddSystem registers a system for the tick loop.
func (w *World) AddSystem(s System) {
	w.systems = append(w.systems, s)
}

// Tick advances the world by dt seconds using fixed-timestep accumulation.
// Delta time is clamped to avoid spiral-of-death after long pauses.
func (w *World) Tick(dt float64) {
	const maxAccumulated = 0.25
	if dt > maxAccumulated {
		dt = maxAccumulated
	}

	w.accumulator += dt

	const maxStepsPerTick = 5
	steps := 0
	for w.accumulator >= w.TickRate && steps < maxStepsPerTick {
		for _, s := range w.systems {
			s.Update(w, w.TickRate)
		}
		w.accumulator -= w.TickRate
		steps++
	}

	if steps == maxStepsPerTick {
		w.accumulator = 0
	}
}
