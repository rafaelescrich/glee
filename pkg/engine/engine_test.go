package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tonyOreglia/glee/pkg/moves"
	"github.com/tonyOreglia/glee/pkg/position"
)

var flagtests = []struct {
	fen           string
	depth         int
	expectedNodes int
	name          string
}{
	{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", 1, 20, ""},
	{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", 2, 400, ""},
	{"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", 1, 48, ""},
	{"n1n5/PPPk4/8/8/8/8/4Kppp/5N1N b - - 0 1", 1, 24, ""},
	{"n1n5/PPPk4/8/8/8/8/4Kppp/5N1N b - - 0 1", 2, 496, ""},
	// // castling through check
	// {"r3k2r/p1ppqNb1/bn2pnp1/3P4/1p2P3/2N2Q1p/PPPBBPPP/R3K2R b KQkq - 0 1", 1, 44, "cannot castle through check"},
	// misses e8f7

	// {"r3k2r/p1ppqNb1/bn2pnp1/3P4/1p2P3/2N2Q1p/PPPBBPPP/R3K2R b KQkq - 0 1", 2, 2080, "good test position"},
	// {"r3k2r/p1ppqNb1/1n2pnp1/3P4/1p2P3/2N2Q1p/PPPBbPPP/R3K2R w KQkq - 0 1", 1, 41, "good test position --> a6e2"},
	// {"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", 2, 2039, ""},
	// {"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", 3, 8092, ""},
}

func setup() (int, int, *moves.Move) {
	perft := 0
	singlePlyPerft := 0
	mv := new(moves.Move)
	return perft, singlePlyPerft, mv
}

func TestMinMax(t *testing.T) {
	for _, tt := range flagtests {
		perft, singlePlyPerft, mv := setup()
		pos, _ := position.NewPositionFen(tt.fen)
		minMax(searchParams{
			depth:          tt.depth,
			ply:            tt.depth,
			pos:            &pos,
			engineMove:     &mv,
			perft:          &perft,
			singlePlyPerft: &singlePlyPerft,
		})
		assert.Equal(t, tt.expectedNodes, perft)
	}
}
