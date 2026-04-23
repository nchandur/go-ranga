package main

const BOARD_SQ_NUM = 120
const MAX_GAMES_MOVES = 4096

var Sq120To64 [BOARD_SQ_NUM]Square
var Sq64To120 [64]Square

var SetMask[64]uint64
var ClearMask[64]uint64