package chessboard

type Piece byte

type Team bool

const (
	WHITE_TEAM Team = false
	BLACK_TEAM Team = true
)

const (
	NOTHING Piece = 0x0
	KING Piece = 0x1 // 0000 0001
	QUEEN Piece = 0x2 // 0000 0010
	TOWER Piece = 0x3 // 0000 0011
	KNIGH Piece = 0x4 // 0000 0100
	FOOL Piece = 0x5  // 0000 0101
	PAWN Piece = 0x6 //  0000 0110

	WHITE_KING Piece = 0x1 // 0000 0001
	WHITE_QUEEN Piece = 0x2 // 0000 0010
	WHITE_TOWER Piece = 0x3 // 0000 0011
	WHITE_KNIGH Piece = 0x4 // 0000 0100
	WHITE_FOOL Piece = 0x5  // 0000 0101
	WHITE_PAWN Piece = 0x6 //  0000 0110

	BLACK_KING Piece = 0x9 // 0000 1001
	BLACK_QUEEN Piece = 0xA // 0000 1010
	BLACK_TOWER Piece = 0xB // 0000 1011
	BLACK_KNIGH Piece = 0xC // 0000 1100
	BLACK_FOOL Piece = 0xD // 0000 1101
	BLACK_PAWN Piece = 0xE // 0000 1110
)

func (this Piece) getTeam() Team {
	return this & 0x8 >> 3 == 1 // recuperation du bit d'equipe
}

func (this Piece) isLegalPiece() bool {
	p := this & 0x7
	return p > 0 && p <= 6
}

func (this Piece) hasAlreadyMoved() bool {
	return this & 0x10 >> 4 == 1
}

func (this *Piece) moveit(){
	(*this) |= 0x10
}

func (this Piece) Teampiece() Piece{
	return this & 0xF
}

func (this Piece) Type() Piece {
	return this & 0x7
}