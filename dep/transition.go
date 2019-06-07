package dep

import (
	"fmt"

	"github.com/chewxy/lingo"
)

// transition is a tuple of Move and label
type transition struct {
	Move
	lingo.DependencyType
}

var transitions []transition
var MAXTRANSITION int

func buildTransitions(labels []lingo.DependencyType) []transition {
	ts := make([]transition, 0)

	for _, m := range AllMoves {
		for _, l := range labels {
			if (m == Shift && l != lingo.NoDepType) || (m != Shift && l == lingo.NoDepType) {
				continue
			}
			ts = append(ts, transition{m, l})
		}
	}
	return ts
}

func (t transition) String() string {
	return fmt.Sprintf("(%s, %s)", t.Move, t.DependencyType)
}

func lookupTransition(t transition, table []transition) (int, error) {
	for i, v := range table {
		if v == t {
			return i, nil
		}
	}
	return 0, fmt.Errorf("transition %v not found", t)
}

// this builds the default transitions
func init() {
	lbls := make([]lingo.DependencyType, lingo.MAXDEPTYPE)

	for i := 0; i < int(lingo.MAXDEPTYPE); i++ {
		lbls[i] = lingo.DependencyType(i)
	}

	transitions = buildTransitions(lbls)
	MAXTRANSITION = len(transitions)
}
