package engine

// System is the interface for all ECS systems.
type System interface {
	// Update processes one tick of the system's logic.
	Update(w *World, dt float64)
}

// PhysicsSystem handles vehicle movement, drifting, boost, and collision.
type PhysicsSystem struct{}

// Update processes physics for all entities with transform and vehicle components.
func (s *PhysicsSystem) Update(w *World, dt float64) {
	// Skeleton: physics logic will be implemented in Phase 2
}

// RaceSystem handles lap counting, position ranking, and finish detection.
type RaceSystem struct{}

// Update processes race state for all entities with race position components.
func (s *RaceSystem) Update(w *World, dt float64) {
	// Skeleton: race logic will be implemented in Phase 2
}

// ItemSystem handles item pickup, targeting, hit validation, and cooldowns.
type ItemSystem struct{}

// Update processes item logic for all entities with item holder components.
func (s *ItemSystem) Update(w *World, dt float64) {
	// Skeleton: item logic will be implemented in Phase 3
}

// TrackSystem handles waypoint graph traversal, shortcuts, and hazard triggers.
type TrackSystem struct{}

// Update processes track logic.
func (s *TrackSystem) Update(w *World, dt float64) {
	// Skeleton: track logic will be implemented in Phase 2
}
