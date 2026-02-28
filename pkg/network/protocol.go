// Package network provides the authoritative server, client prediction,
// and network protocol types.
package network

// MessageType identifies the kind of network message.
type MessageType uint8

const (
	MsgInput         MessageType = iota // Client input
	MsgSnapshot                         // Server state snapshot
	MsgRaceState                        // Race state broadcast
	MsgItemFire                         // Item fire event
	MsgItemHit                          // Item hit confirmation
	MsgConnect                          // Client connect
	MsgDisconnect                       // Client disconnect
)

// Message is a network message envelope.
type Message struct {
	Type    MessageType
	Tick    uint64
	Payload []byte
}

// RaceSnapshot holds a delta-compressed race state.
type RaceSnapshot struct {
	Tick       uint64
	Positions  []EntityState
	ItemStates []ItemState
	LapData    []LapState
}

// EntityState holds the network-synced state of one entity.
type EntityState struct {
	EntityID   int
	X, Y, Z    float64
	VX, VY, VZ float64
	Rotation   float64
}

// ItemState holds the network-synced item state of one entity.
type ItemState struct {
	EntityID    int
	CurrentItem int
	HasItem     bool
}

// LapState holds the network-synced lap state of one entity.
type LapState struct {
	EntityID int
	Lap      int
	Finished bool
}
