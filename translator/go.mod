module RomanNumeralParser/translator

go 1.17

replace part1/scanner => ../scanner

require (
	part1/logger v0.0.0-00010101000000-000000000000
	part1/roman v0.0.0-00010101000000-000000000000
	part1/scanner v0.0.0-00010101000000-000000000000
)

replace part1/roman => ../roman

replace part1/logger => ../logger
