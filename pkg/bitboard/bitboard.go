// Package bitboard implements utility routines for
// using chess engine bitboards.
package bitboard

import (
	"fmt"
	"math/bits"
)

// Bitboard struct exposes uint64 "bitboard" with associated getter, setter, and helper fxns
type Bitboard struct {
	bitboard uint64
}

func NewBitboard(bb uint64) *Bitboard {
	return &Bitboard{bitboard: bb}
}

func NewBitboardFromIndex(i int) *Bitboard {
	bb := &Bitboard{}
	bb.SetBit(i)
	return bb
}

// NewBbFromMovesSlice takes a legal moves slice and returns a bitboard representing those moves
func NewBbFromMovesSlice(mvs [][2]int) *Bitboard {
	bitboard := &Bitboard{}
	for i := 0; i < len(mvs); i++ {
		bitboard.SetBit(mvs[i][1])
	}
	return bitboard
}

func (b *Bitboard) Set(bb uint64) error {
	b.bitboard = bb
	return nil
}

func (b *Bitboard) Combine(bb *Bitboard) *Bitboard {
	b.bitboard |= bb.Value()
	return b
}

func ReturnCombined(bb *Bitboard, bb2 *Bitboard) *Bitboard {
	combinedBb := NewBitboard(bb.Value() | bb2.Value())
	return combinedBb
}

func (b *Bitboard) BitwiseAnd(bb *Bitboard) *Bitboard {
	b.bitboard &= bb.bitboard
	return b
}

func ReturnBitwiseAnd(bb *Bitboard, b *Bitboard) *Bitboard {
	overlap := NewBitboard(b.bitboard & bb.bitboard)
	return overlap
}

func ReturnOverlapBb(bb1 *Bitboard, bb2 *Bitboard) *Bitboard {
	overlapBb := NewBitboard(bb1.Value() & bb2.Value())
	return overlapBb
}

func (b *Bitboard) PopulationCount() int {
	copy := NewBitboard(b.Value())
	count := 0
	for copy.Value() != 0 {
		copy.RemoveBit(copy.Lsb())
		count++
	}
	return count
}

func (b *Bitboard) RemoveOverlappingBits(bb *Bitboard) *Bitboard {
	b.bitboard &^= bb.Value()
	return b
}

func (b *Bitboard) SetBit(bit int) {
	b.bitboard |= uint64(1) << uint(bit)
}

func GetShiftedLeftBb(b *Bitboard, shift uint) *Bitboard {
	shiftedBb := new(Bitboard)
	shiftedBb.Set(b.bitboard << shift)
	return shiftedBb
}

func GetShiftedRightBb(b *Bitboard, shift uint) *Bitboard {
	shiftedBb := new(Bitboard)
	shiftedBb.Set(b.bitboard >> shift)
	return shiftedBb
}

func (b *Bitboard) IsZero() bool {
	return b.Value() == uint64(0)
}

func (b *Bitboard) BitIsSet(bit int) bool {
	return ((uint64(1) << uint(bit)) & b.bitboard) != uint64(0)
}

func (b *Bitboard) BitIsNotSet(bit int) bool {
	return !b.BitIsSet(bit)
}

func (b *Bitboard) ReturnBitsFlipped() *Bitboard {
	newBB := new(Bitboard)
	newBB.Set(^b.Value())
	return newBB
}

func (b *Bitboard) GetBitValue(bit int) uint {
	if b.BitIsSet(bit) {
		return 1
	}
	return 0
}

func (b *Bitboard) Value() uint64 {
	return b.bitboard
}

// count leading zeros
func (b *Bitboard) Msb() int {
	return 63 - int(bits.LeadingZeros64(b.Value()))
}

// count leading zeros
func (b *Bitboard) Lsb() int {
	return int(bits.TrailingZeros64(b.Value()))
}

func (b *Bitboard) RemoveBit(pos int) {
	b.bitboard &= ^(uint64(1) << uint(pos))
}

func (b *Bitboard) Print() {
	var i int
	fmt.Println("")
	for i = 0; i < 64; i++ {
		var sq int
		if b.BitIsSet(i) {
			sq = 1
		}
		fmt.Print(sq)
		if ((i + 1) % 8) == 0 {
			fmt.Println("")
		}
	}
	fmt.Println("")
	fmt.Println("")
}
