package examples

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	chessboard "github.com/sylvainsausse/chess-engine"
)

func main(){
	c := chessboard.NewChessboard()
	//c.Load("0xB0009DCBEEEAEEEE00C0D000000E000000000000666660000000066634521543")
	i := 0
	for true {
		if c.CheckMate(i%2==1){
			fmt.Println("CHECKMATE")
			os.Exit(0)
		}
		if !(i%2==1) {
			var P1,P2 string
			
			c.Disp()
			fmt.Scanf("%s %s",&P1,&P2)
			l1,c1 := chessboard.PosToCord(P1)
			l2,c2 := chessboard.PosToCord(P2)
			err := c.Make_move(i%2 == 1,l1,c1,l2,c2)
			for err != nil {
				k,l := chessboard.PosToCord(P1)
				fmt.Println(P1,"->",P2,err.Error(),chessboard.RuneArray[c.GetPieceInPos(P1)])
				fmt.Println("les coups possible de cette pieces sont :",c.GetAllPossiblePlaysFrom(i%2==1,k,l))
				fmt.Scanf("%s %s",&P1,&P2)
				l1,c1 = chessboard.PosToCord(P1)
				l2,c2 = chessboard.PosToCord(P2)
				err = c.Make_move(i%2 == 1,l1,c1,l2,c2)
			}
		} else {
			list := c.GetAllPlays(i%2==1)
			chosen := list[rand.Int()%len(list)]
			fmt.Println(chosen)
			err := c.Make_move(i%2==1,chosen[0][0],chosen[0][1],chosen[1][0],chosen[1][1])
			if err != nil {
				log.Fatal(err.Error())
			}
		}
		i ++
	}
	

}