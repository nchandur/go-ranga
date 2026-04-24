package main

const BOARD_SQ_NUM = 120
const MAX_GAMES_MOVES = 4096

var Sq120To64 [BOARD_SQ_NUM]Square
var Sq64To120 [64]Square

var FileBoard [BOARD_SQ_NUM]File
var RankBoard [BOARD_SQ_NUM]Rank

var SetMask [64]uint64
var ClearMask [64]uint64

var PieceKeys [13][120]uint64
var SideKey uint64
var CastleKey [16]uint64

const PieceChar = ".PNBRQKpnbrqk"
const SideChar = "wb-"
const RankChar = "12345678"
const FileChar = "abcdefgh"

var PieceBig = [13]bool{false, false, true, true, true, true, true, false, true, true, true, true, true}
var PieceMajor = [13]bool{false, false, false, false, true, true, true, false, false, false, true, true, true}
var PieceMinor = [13]bool{false, false, true, true, false, false, false, false, true, true, false, false, false}
var PieceValue = [13]int{0, 100, 325, 325, 550, 1000, 50000, 100, 325, 325, 550, 1000, 500000}
var PieceColor = [13]Color{Both, White, White, White, White, White, White, Black, Black, Black, Black, Black, Black}

var IsPawn = [13]bool{false, true, false, false, false, false, false, true, false, false, false, false, false}
var IsKnight = [13]bool{false, false, true, false, false, false, false, false, true, false, false, false, false}
var IsBishop = [13]bool{false, false, false, true, false, false, false, false, false, true, false, false, false}
var IsRook = [13]bool{false, false, false, false, true, false, false, false, false, false, true, false, false}
var IsQueen = [13]bool{false, false, false, false, false, true, false, false, false, false, false, true, false}
var IsKing = [13]bool{false, false, false, false, false, false, true, false, false, false, false, false, true}

var KnightDirections = [8]int{-21, -19, -12, -8, 8, 12, 19, 21}
var BishopDirections = [4]int{-11, -9, 9, 11}
var RookDirections = [4]int{-10, -1, 1, 10}
var KingDirections = [8]int{-11, -10, -9, -1, 1, 9, 10, 11} // also queen directions