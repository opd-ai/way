package engine

// TransformComponent stores position and rotation.
type TransformComponent struct {
	X, Y, Z    float64
	Rotation   float64
	VelocityX  float64
	VelocityY  float64
	VelocityZ  float64
}

// VehiclePhysicsComponent stores vehicle physics state.
type VehiclePhysicsComponent struct {
	Speed       float64
	Acceleration float64
	MaxSpeed    float64
	Handling    float64
	Weight      float64
	DriftFactor float64
	BoostBar    float64
	IsDrifting  bool
	IsBoosting  bool
}

// ItemHolderComponent stores item pickup and usage state.
type ItemHolderComponent struct {
	HeldItem   int
	HasItem    bool
	Cooldown   float64
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
	TargetEntity Entity
	OffsetX      float64
	OffsetY      float64
	OffsetZ      float64
	Perspective  string // Must be "over-the-shoulder"
	Distance     float64
	Height       float64
	Angle        float64
}
