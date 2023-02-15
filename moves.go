package chessboard

func (this *Chessboard) IsLegalMove(team Team, line1 int, col1 int, line2 int, col2 int) bool {
	piece := this.GetPieceAt(line1,col1) 
	real_team :=  piece.getTeam()// recuperation du bit d'equipe


	if real_team != team || !piece.isLegalPiece() { // La pièce est-elle valide et est elle de la bonne equipe ?
		return false
	}

	piecedest := this.GetPieceAt(line2,col2)
	destteam := piecedest.getTeam()

	if destteam == team && piecedest.isLegalPiece(){ // Si la destination contient un pièce de la meme equipe 
		return false
	}
	b,l1,c1,l2,c2 := this.IsALegalRook(team,line1,col1,line2,col2)
	if !this.PossiblyMove(line1,col1,line2,col2){
		return false
	}
	c := this.clone()
	if !this.IsEnPassant(team,line1,col1,line2,col2){
		c.move(line1,col1,line2,col2)
		c.delete(line1,col2)
	} else if b {
		c.move(line1,col1,line2,col2)
		c.move(l1,c1,l2,c2)
	} else {
		c.move(line1,col1,line2,col2)
	}
	if c.CheckForChecks(team){
		return false
	}

	return true
}

func (this *Chessboard) IsLegalMove_NoCheck(team Team, line1 int, col1 int, line2 int, col2 int) bool {
	piece := this.GetPieceAt(line1,col1) 
	real_team :=  piece.getTeam()// recuperation du bit d'equipe


	if real_team != team || !piece.isLegalPiece() { // La pièce est-elle valide et est elle de la bonne equipe ?
		return false
	}

	piecedest := this.GetPieceAt(line2,col2)
	destteam := piecedest.getTeam()

	if destteam == team && piecedest.isLegalPiece(){ // Si la destination contient un pièce de la meme equipe 
		return false
	}

	if this.IsEnPassant(team,line1,col1,line2,col2){
		return true
	}

	b,_,_,_,_ := this.IsALegalRook(team,line1,col1,line2,col2)

	if b {
		return true
	}


	if !this.PossiblyMove(line1,col1,line2,col2){
		return false
	}

	return true
}

func (this Chessboard) PossiblyMove(line1 int, col1 int, line2 int, col2 int) bool {
	piece := this.GetPieceAt(line1,col1)
	if line2 < 0 || line2 >= 8 || col2 < 0 || col2 >= 8 || line1 < 0 || line1 >= 8 || col1 < 0 || col1 >= 8 {
		return false
	}
	if line1-line2 == 0 && col1-col2 == 0 {
		return false
	}
	switch(piece.Type()){ // Filtration du bit d'equipe (si pas deja fait au praravent)
		case KING: 
			if abs(line1 - line2)  <= 1 && abs(col1-col2) <= 1 {
				return true
			}
			break
		case QUEEN :
			// deplacement vertical
			if line1-line2 == 0 { 
				minimum := min(col1,col2)
				maximum := max(col1,col2)
				for i := minimum+1  ; i < maximum ; i++ {
					// si il y a une piece sur le chemin
					if this.GetPieceAt(line2,i).isLegalPiece() { 
						return false
					}
				}
				return true
			}
			// deplacement horizontal
			if col1-col2 == 0 {
				minimum := min(line1,line2)
				maximum := max(line1,line2)
				for i := minimum+1  ; i < maximum ; i++ {
					// si il y a une piece sur le chemin
					if this.GetPieceAt(i,col2).isLegalPiece() { 
						return false
					}
				}
				return true
			}
			// deplacement diagonal
			var mil,mic int
			if col1-col2 == line1-line2 {
				if min(col1,col2) == col1 {
					mil = line1
					mic = col1
				} else {
					mil = line2
					mic = col2
				}
				for i := 1 ; i < abs(line1-line2); i++ {
					if this.GetPieceAt(mil+i,mic+i).isLegalPiece() {
						//println("+ piece sur le chemin",CordToPos(mil+i,mic+i))
						return false
					}
					//println("+",CordToPos(mil+i,mic+i))
				}
				return true
			}
			if col1-col2 == line2-line1 {
				if min(col1,col2) == col1 {
					mil = line1
					mic = col1
				} else {
					mil = line2
					mic = col2
				}
				for i := 1 ; i < abs(line1-line2) ; i++ {
					if this.GetPieceAt(mil-i,mic+i).isLegalPiece() {
						//println("- piece sur le chemin",CordToPos(mil-i,mic+i))
						return false
					}
					//println("-",CordToPos(mil-i,mic+i))
				}
				return true
			}
			break
		case TOWER :
			// deplacement vertical
			if line1-line2 == 0 { 
				minimum := min(col1,col2)
				maximum := max(col1,col2)
				for i := minimum+1  ; i < maximum ; i++ {
					// si il y a une piece sur le chemin
					if this.GetPieceAt(line2,i).isLegalPiece() { 
						return false
					}
				}
				return true
			}
			// deplacement horizontal
			if col1-col2 == 0 {
				minimum := min(line1,line2)
				maximum := max(line1,line2)
				for i := minimum+1  ; i < maximum ; i++ {
					// si il y a une piece sur le chemin
					if this.GetPieceAt(i,col2).isLegalPiece() { 
						return false
					}
				}
				return true
			}
			break
		case FOOL :
			// deplacement diagonal
			var mil,mic int
			if col1-col2 == line1-line2 {
				if min(col1,col2) == col1 {
					mil = line1
					mic = col1
				} else {
					mil = line2
					mic = col2
				}
				for i := 1 ; i < abs(line1-line2); i++ {
					if this.GetPieceAt(mil+i,mic+i).isLegalPiece() {
						//println("+ piece sur le chemin",CordToPos(mil+i,mic+i))
						return false
					}
					//println("+",CordToPos(mil+i,mic+i))
				}
				return true
			}
			if col1-col2 == line2-line1 {
				if min(col1,col2) == col1 {
					mil = line1
					mic = col1
				} else {
					mil = line2
					mic = col2
				}
				for i := 1 ; i < abs(line1-line2) ; i++ {
					if this.GetPieceAt(mil-i,mic+i).isLegalPiece() {
						//println("- piece sur le chemin",CordToPos(mil-i,mic+i))
						return false
					}
					//println("-",CordToPos(mil-i,mic+i))
				}
				return true
			}
			break
		case KNIGH :
			l := abs(line1-line2)
			c := abs(col1-col2)
			if (l == 2 && c == 1)||(l == 1 && c == 2) {
				return true
			}
			break
		case PAWN :
			team := piece.getTeam()
			// Opening
			if line1-line2 == -2 && col1-col2 == 0  && team && line1 == 1 {
				return !this.GetPieceAt(line2,col2).isLegalPiece() && !this.GetPieceAt(line2-1,col2).isLegalPiece() 
			}
			if line1-line2 == 2 && col1-col2 == 0  && !team && line1 == 6 {
				return !this.GetPieceAt(line2,col2).isLegalPiece() && !this.GetPieceAt(line2+1,col2).isLegalPiece() 
			}
			if line1-line2 == -1 && team {
				if col1-col2 == 0 {
					return !this.GetPieceAt(line2,col2).isLegalPiece()
				}
				if abs(col1-col2) == 1 {
					p := this.GetPieceAt(line2,col2)
					return p.isLegalPiece() //en passant
				}
			}
			if line1-line2 == 1 && !team {
				if col1-col2 == 0 {
					return !this.GetPieceAt(line2,col2).isLegalPiece()
				}
				if abs(col1-col2) == 1 {
					p := this.GetPieceAt(line2,col2)
					return p.isLegalPiece() //en passant
				}
			}
			break
	}
	return false
}

func (this *Chessboard) Make_move(t Team, line1 int, col1 int, line2 int, col2 int) (error){
	if this.IsLegalMove(t,line1,col1,line2,col2) {
		b,l1,c1,l2,c2 := this.IsALegalRook(t,line1,col1,line2,col2)
		if b {
			this.move(line1,col1,line2,col2)
			this.move(l1,c1,l2,c2)
			return nil
		}
		if this.IsEnPassant(t,line1,col1,line2,col2) {
			this.move(line1,col1,line2,col2)
			return this.delete(line1,col2)
		}
		return this.move(line1,col1,line2,col2)
	}
	
	return ChessError{"Not a legal move"}
}

func (this *Chessboard) move(line1 int, col1 int, line2 int, col2 int) error {
	if line2 < 0 || line2 >= 8 || col2 < 0 || col2 >= 8 || line1 < 0 || line1 >= 8 || col1 < 0 || col1 >= 8 {
		return ChessError{"Bad coordinates"}
	}
	p := this[line1*8+col1]
	if p.Type() == PAWN {
		if p.getTeam() == WHITE_TEAM && line2 == 0 {
			this[line2*8+col2] = WHITE_QUEEN
			this[line1*8+col1] = 0
			this[line2*8+col2].moveit()
			return nil
		}
		if p.getTeam() == BLACK_TEAM && line2 == 7 {
			this[line2*8+col2] = BLACK_QUEEN
			this[line1*8+col1] = 0
			this[line2*8+col2].moveit()
			return nil
		}
	}
	this[line2*8+col2] = this[line1*8+col1]
	this[line1*8+col1] = 0
	this[line2*8+col2].moveit()
	return nil
}

func (this *Chessboard) delete(l int, c int) error {
	if l < 0 || l >= 8 || c < 0 || c >= 8 {
		return ChessError{"Bad coordinates"}
	}
	this[l*8+c] = 0
	return nil
}

func (e ChessError) Error() string{
	return e.msg
}

func (this Chessboard) GetAllPossiblePlaysFromDigest(t Team, l int, c int) []string {
	ret := []string{}
	for i := 0 ; i<8; i++ {
		for j := 0 ; j<8 ; j++ {
			b,_,_,_,_ := this.IsALegalRook(t,l,c,i,j)
			if this.IsLegalMove(t,l,c,i,j) || b || this.IsEnPassant(t,l,c,i,j){
				ret = append(ret, CordToPos(i,j))
			}
		}
	}
	return ret
}

func (this Chessboard) GetAllPlaysDigest(t Team) []string{
	ret := []string{}
	for i := 0 ; i<8; i++ {
		for j := 0 ; j<8 ; j++ {
			for _,item := range this.GetAllPossiblePlaysFromDigest(t,i,j) {
				ret = append(ret,CordToPos(i,j) + "->" + item)
			}
		}
	}
	return ret
}

func (this Chessboard) CheckForChecksAt(t Team,ci int,cj int) bool {
	for i := 0 ; i<8 ; i++ {
		for j := 0 ; j<8 ; j++ {
			a := this.GetPieceAt(i,j)
			if a.getTeam() != t {
				if this.IsLegalMove_NoCheck(!t,i,j,ci,cj) {
					return true
				}
			}
		}
	}
	return false
}

func (this Chessboard) CheckForChecks(t Team) bool {
	// find the king
	ci := -1
	cj := -1
	for i := 0 ; i<8 ; i++ {
		for j := 0 ; j<8 ; j++ {
			a := this.GetPieceAt(i,j)
			if a.Type() == KING && a.getTeam() == t {
				ci = i
				cj = j
				break
			}
		}
	}
	if ci == -1 || cj == -1 {
		panic("no king !")
	}
	return this.CheckForChecksAt(t,ci,cj)
}


func (this Chessboard) IsALegalRook(t Team, line1 int, col1 int, line2 int, col2 int) (bool,int,int,int,int) {	
	if t && line1 == 0 && line2 == 0 && col1 == 4 {
		if col2 == 2 {
			a := this.GetPieceAt(0,4)
			b := this.GetPieceAt(0,0)
			if a == BLACK_KING && !a.hasAlreadyMoved() && b == BLACK_TOWER && !b.hasAlreadyMoved() {
				for i := 1 ; i < 4 ; i++ {
					if this.GetPieceAt(0,i).isLegalPiece() || this.CheckForChecksAt(t,0,i){
						return false,-1,-1,-1,-1
					}
				}
				return true,0,0,0,3
			}
		}
		if col2 == 6 {
			a := this.GetPieceAt(0,4)
			b := this.GetPieceAt(0,7)
			if a == BLACK_KING && !a.hasAlreadyMoved() && b == BLACK_TOWER && !b.hasAlreadyMoved() {
				for i := 5 ; i < 7 ; i++ {
					if this.GetPieceAt(0,i).isLegalPiece() || this.CheckForChecksAt(t,0,i){
						return false,-1,-1,-1,-1
					}
				}
				return true,0,7,0,5
			}
		}
	} else if line1 == 7 && line2 == 7 && col1 == 4 {
		if col2 == 2 {
			a := this.GetPieceAt(7,4)
			b := this.GetPieceAt(7,0)
			if a == WHITE_KING && !a.hasAlreadyMoved() && b == WHITE_TOWER && !b.hasAlreadyMoved() {
				for i := 1 ; i < 4 ; i++ {
					if this.GetPieceAt(7,i).isLegalPiece() || this.CheckForChecksAt(t,7,i){
						return false,-1,-1,-1,-1
					}
				}
				return true,7,0,7,3
			}
		}
		if col2 == 6 {
			a := this.GetPieceAt(7,4)
			b := this.GetPieceAt(7,7)
			if a == WHITE_KING && !a.hasAlreadyMoved() && b == WHITE_TOWER && !b.hasAlreadyMoved() {
				for i := 5 ; i < 7 ; i++ {
					if this.GetPieceAt(7,i).isLegalPiece() || this.CheckForChecksAt(t,7,i){
						return false,-1,-1,-1,-1
					}
				}
				return true,7,7,7,5
			}
		}
	}
	return false,-1,-1,-1,-1
}

func (this Chessboard) IsEnPassant(t Team, l1 int, c1 int, l2 int, c2 int) bool{
	p := this.GetPieceAt(l1,c1)
	if p.Type() != PAWN || p.getTeam() != t {
		return false
	}
	if l1-l2 == -1 && t { //noir
		if abs(c1-c2) == 1 {
			p := this.GetPieceAt(l1,c2)
			return p.isLegalPiece() && p.getTeam() != t //en passant
		}
	}
	if l1-l2 == 1 && !t { //bkanc
		if abs(c1-c2) == 1 {
			p := this.GetPieceAt(l1,c2)
			return p.isLegalPiece() && p.getTeam() != t  //en passant
		}
	}
	return false
}

func (this Chessboard) CheckMate(t Team) bool {
	for i := 0 ; i<8; i++ {
		for j := 0 ; j<8 ; j++ {
			if len(this.GetAllPossiblePlaysFrom(t,i,j))>0 {
				return false
			}
		}
	}
	return true
}

func (this Chessboard) GetAllPossiblePlaysFrom(t Team, l int, c int) [][2]int {
	ret := [][2]int{}
	for i := 0 ; i<8; i++ {
		for j := 0 ; j<8 ; j++ {
			b,_,_,_,_ := this.IsALegalRook(t,l,c,i,j)
			if this.IsLegalMove(t,l,c,i,j) || b || this.IsEnPassant(t,l,c,i,j){
				ret = append(ret, [2]int{i,j})
			}
		}
	}
	return ret
}

func (this Chessboard) GetAllPlays(t Team) [][2][2]int{
	ret := [][2][2]int{}
	for i := 0 ; i<8; i++ {
		for j := 0 ; j<8 ; j++ {
			for _,item := range this.GetAllPossiblePlaysFrom(t,i,j) {
				ret = append(ret,[2][2]int{{i,j},{item[0],item[1]}})
			}
		}
	}
	return ret
}