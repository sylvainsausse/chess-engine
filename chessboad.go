package chessboard

import (
	"fmt"
)

type Chessboard [64]Piece


var RuneArray = [...]string{"  "," ♔" ," ♕" ," ♖" ," ♘" ," ♗" ," ♙" ," "," "," ♚" ," ♛" ," ♜" ," ♞" ," ♝" ," ♟" } 

var Cases = [...]string{" □" , " ■" } 

var Blank = Chessboard{
	0xB,0xC,0xD,0xA,0x9,0xD,0xC,0xB,
	0xE,0xE,0xE,0xE,0xE,0xE,0xE,0xE,
	0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,
	0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,
	0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,
	0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,
	0x6,0x6,0x6,0x6,0x6,0x6,0x6,0x6,
	0x3,0x4,0x5,0x2,0x1,0x5,0x4,0x3,
}

const (
	BLANK_1 uint64 = 0x9AB87BA9
	BLANK_2 uint64 = 0xCCCCCCCC
	BLANK_3456 uint64 = 0x00000000
	BLANK_7 uint64 = 0x66666666
	BLANK_8 uint64 = 0x34521543
)

func NewChessboard() Chessboard {
	return Blank.clone()
}

func (this *Chessboard) Load(c string) error {
	for i := 2; i<len(c) ; i++ {
		if c[i] >= '0' && c[i] <= '9' {
			this[i-2] = Piece(c[i] - '0')
		} else if c[i] >= 'A' && c[i] <= 'F' {
			this[i-2] = Piece(c[i] - 'A' + 10)
		} else {
			return ChessError{"Not a chessboard"}
		}
	}
	return nil
}

func (this Chessboard) clone() Chessboard {
	c := Chessboard{}
	for i,j := range this {
		c[i] = j
	}
	return c
}

func (this Chessboard) Convert() []byte {
	ret := make([]byte,64)
	for i,j := range this {
		ret[i] = byte(j)
	}
	return ret
}

func (this *Chessboard) LoadFromBytes(bytes []byte){
	for i := 0 ; i < 64 ; i++ {
		this[i] = Piece(bytes[i])
	}
	return
}


func (this Chessboard) GetPieceAt(line int, column int) Piece {
	if line >= 8 || line < 0 || column >= 8 || column < 0 {
		return 0
	}
	return this[line*8+column]
}

func PosToCord(pos string) (int, int){
	if len(pos) > 2 || len(pos) == 0{
		return -1,-1
	}
	if pos[0] > 'H' || pos[0] < 'A' {
		return -1,-1
	}
	if pos[1] > '8' || pos[0] < '1' {
		return -1,-1
	}
	return 7-int(pos[1]-'1'),int(pos[0]-'A')
}

func CordToPos(i int, j int) string {
	return string([]byte{'A'+byte(j),'1'+(7-byte(i))})
}

func (this Chessboard) GetPieceInPos(pos string) Piece {
	i,j := PosToCord(pos)
	if i == -1 || j == -1 {
		return 0
	}
	println(i,j)
	return this.GetPieceAt(i,j)
}

func (this Chessboard) Disp(){
	fmt.Println(this.ToString())
	fmt.Println("   A B C D E F G H")
	for i := 0; i<8; i++ {
		fmt.Printf("%d ",8-i)
		for j := 0; j<8; j++{
			value := this.GetPieceAt(i,j)
			if value == 0 {
				fmt.Printf("%s",Cases[(i+j)%2])
			} else {
				fmt.Printf("%s",RuneArray[this.GetPieceAt(i,j).Teampiece()])
			}
		}
		fmt.Println()
	}
}

func (this Chessboard) ToString() string {
	conv := "0123456789ABCDEF"
	ret := "0x"
	for _,j := range this {
		ret += string(conv[j.Teampiece()])
	}
	return ret
}

type ChessError struct {
	msg string
}