package audio

// SFXType identifies a sound effect category.
type SFXType int

const (
	SFXCollision SFXType = iota
	SFXItemUse
	SFXBoost
	SFXDriftSqueal
	SFXFinishFanfare
)

// GenerateSFX produces a procedural sound effect buffer for the given type and seed.
func GenerateSFX(sfxType SFXType, seed int64) []float64 {
	// Skeleton: SFX synthesis will be implemented in Phase 3
	return nil
}
