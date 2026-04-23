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
