package main

type Color int

const (
	White Color = iota
	Black
	Both
)

func (c *Color) Valid() bool {
	return *c == White || *c == Black
}

func (c *Color) String() string {
	if *c == White {
		return "W"
	}

	if *c == Black {
		return "B"
	}

	return "."

}
