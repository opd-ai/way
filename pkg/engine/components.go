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

// ItemHolderComponent stores item inventory state.
type ItemHolderComponent struct {
	CurrentItem int
	HasItem     bool
	Cooldown    float64
}

// RacePositionComponent stores race progress state.
type RacePositionComponent struct {
	Lap           int
	Checkpoint    int
	Position      int
	TotalLaps     int
	Finished      bool
	RaceTime      float64
}
