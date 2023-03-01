package board

type CastleRights struct {
	KingSide bool
	QueenSide bool
}

type ChessBoard struct {
	BoardState []int
	PiecesBlack Pieces
	PiecesWhite Pieces

	ActiveColor Color

	CastleRightsBlack CastleRights
	CastleRightsWhite CastleRights

	EnPassantTarget BoardLocation

	HalfMoveClock int
	FullMoveNumber int
}
