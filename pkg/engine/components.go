package engine

// TransformComponent stores position and rotation.
type TransformComponent struct {
	X, Y, Z   float64
	Rotation  float64
	VelocityX float64
	VelocityY float64
	VelocityZ float64
}

// VehiclePhysicsComponent stores vehicle physics state.
type VehiclePhysicsComponent struct {
	Speed        float64
	Acceleration float64
	MaxSpeed     float64
	Handling     float64
	Weight       float64
	DriftFactor  float64
	BoostBar     float64
	IsDrifting   bool
	IsBoosting   bool
}

// ItemHolderComponent stores item pickup and usage state.
type ItemHolderComponent struct {
	HeldItem int
	HasItem  bool
	Cooldown float64
}

// RacePositionComponent stores race progress for an entity.
type RacePositionComponent struct {
	Lap        int
	Checkpoint int
	Position   int
	Finished   bool
	RaceTime   float64
}

// CameraComponent defines the camera position and perspective mode.
type CameraComponent struct {
	TargetEntity Entity  // Entity to follow (typically the player kart)
	Perspective  string  // Must be "over-the-shoulder"
	Distance     float64 // Distance behind the target entity
	Height       float64 // Height above the target entity
	Angle        float64 // Camera tilt angle in degrees
	PositionX    float64 // Current camera world X position
	PositionY    float64 // Current camera world Y position
	PositionZ    float64 // Current camera world Z position
}
