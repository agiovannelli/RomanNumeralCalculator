module RomanNumeralParser/Part1_agiovannelli

go 1.17

replace part1/translator => ../translator

require part1/translator v0.0.0-00010101000000-000000000000

require (
	part1/logger v0.0.0-00010101000000-000000000000 // indirect
	part1/roman v0.0.0-00010101000000-000000000000 // indirect
	part1/scanner v0.0.0-00010101000000-000000000000 // indirect
)

replace part1/scanner => ../scanner

replace part1/roman => ../roman

replace part1/logger => ../logger
