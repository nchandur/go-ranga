package main

type File int

const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
	FileNone
)

func (f *File) Valid() bool {
	return *f >= FileA && *f <= FileH

}
