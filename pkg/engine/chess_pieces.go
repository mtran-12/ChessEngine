package engine

// File is traditionally a Char between A and H, but computationally, it is easier to represent it as an int
type BoardLocation struct {
	Rank int
	File int
}

// King is a single location because there can only be one in a given game
type Pieces struct {
	King BoardLocation
	Paws []BoardLocation
	Bishops []BoardLocation
	Knight []BoardLocation
	Rooks []BoardLocation
	Queens []BoardLocation
}
