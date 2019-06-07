package dep

// Move is an action that the dependency parser can take - whether to Shift, Attach-Left, or AttachRight
type Move byte

//go:generate stringer -type=Move

const (
	Shift Move = iota
	Left
	Right

	MAXMOVE
)

// AllMoves is the set of all possible moves
var AllMoves = []Move{Left, Right, Shift}
