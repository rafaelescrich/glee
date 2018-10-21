package position

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertAlgebriacToIndex(t *testing.T) {
	index := convertAlgebriacToIndex("a8")
	assert.Equal(t, 0, index)
	index = convertAlgebriacToIndex("h1")
	assert.Equal(t, 63, index)
	index = convertAlgebriacToIndex("h8")
	assert.Equal(t, 7, index)
}

func TestTokenizeFen(t *testing.T) {
	position, activeSide, castlingRights, enPassante, moveCt, halfMoveCt := getFenStringTokens("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	assert.Equal(t, position, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	assert.Equal(t, activeSide, White)
	assert.Equal(t, castlingRights, "KQkq")
	assert.Equal(t, enPassante, 64)
	assert.Equal(t, moveCt, 0)
	assert.Equal(t, halfMoveCt, 1)

	position, activeSide, castlingRights, enPassante, moveCt, halfMoveCt = getFenStringTokens("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b q e3 1 2")
	assert.Equal(t, position, "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R")
	assert.Equal(t, activeSide, Black)
	assert.Equal(t, castlingRights, "q")
	assert.Equal(t, enPassante, 44)
	assert.Equal(t, moveCt, 1)
	assert.Equal(t, halfMoveCt, 2)
}

func TestPositionContructorFen(t *testing.T) {
	position, _ := NewPositionFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Kkq - 0 1")
	assert.Equal(t, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Kkq - 0 1", position.GetFenString())

	position, _ = NewPositionFen("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b q e3 1 2")
	assert.Equal(t, "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b q e3 1 2", position.GetFenString())

	position, _ = NewPositionFen("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b - - 1 2")
	assert.Equal(t, "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b - - 1 2", position.GetFenString())

	position, _ = NewPositionFen("7k/8/8/8/8/8/8/6KB w q - 0 1")
	assert.Equal(t, "7k/8/8/8/8/8/8/6KB w q - 0 1", position.GetFenString())

	position, _ = NewPositionFen("7k/8/8/8/8/8/8/Rq6 w KQkq - 0 1")
	assert.Equal(t, "7k/8/8/8/8/8/8/Rq6 w KQkq - 0 1", position.GetFenString())
}

func TestPositionUpdate(t *testing.T) {
	position, _ := NewPositionFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	position.MakeMoveAlgebraic("e2", "e3", White)
	assert.Equal(t, position.GetFenString(), "rnbqkbnr/pppppppp/8/8/8/4P3/PPPP1PPP/RNBQKBNR b KQkq - 1 1")
}

func TestUnMakeMove(t *testing.T) {
	// unmake single move
	position, _ := NewPositionFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	position.MakeMoveAlgebraic("e2", "e3", White)
	position = position.UnMakeMove()
	assert.Equal(t, position.GetFenString(), "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	// unmaking multiple moves in a row
	position, _ = NewPositionFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	position.MakeMoveAlgebraic("e2", "e3", White)
	position.MakeMoveAlgebraic("e7", "e6", Black)
	position.MakeMoveAlgebraic("d2", "d4", White)
	assert.Equal(t, "rnbqkbnr/pppp1ppp/4p3/8/3P4/4P3/PPP2PPP/RNBQKBNR b KQkq d3 2 1", position.GetFenString())
	position = position.UnMakeMove()
	position = position.UnMakeMove()
	position = position.UnMakeMove()
	assert.Equal(t, position.GetFenString(), "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	// unmake attacking move
	position, _ = NewPositionFen("7k/8/8/8/8/8/7p/6KR w q - 0 1")
	position.MakeMoveAlgebraic("h1", "h2", White)
	assert.Equal(t, position.GetFenString(), "7k/8/8/8/8/8/7R/6K1 b q - 1 1")
	position = position.UnMakeMove()
	assert.Equal(t, position.GetFenString(), "7k/8/8/8/8/8/7p/6KR w q - 0 1")

	//unmake en passante move
	position, _ = NewPositionFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	position.MakeMoveAlgebraic("e2", "e4", White)
	assert.Equal(t, position.GetFenString(), "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 1 1")
	position = position.UnMakeMove()
	assert.Equal(t, position.GetFenString(), "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
}

func TestPrintPos(t *testing.T) {
	position, _ := NewPositionFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	position.Print()
	position, _ = NewPositionFen("7k/8/8/8/8/8/7p/6KR w q - 0 1")
	position.Print()
	assert.Equal(t, 1, 1)
}
