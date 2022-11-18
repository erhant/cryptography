package constants

// as per the ETAOIN SHRDLU joke on the question.
// see also: https://en.wikipedia.org/wiki/Etaoin_shrdlu
// values from: https://blogs.sas.com/content/iml/2014/09/19/frequency-of-letters.html
var EnglishLetterFreqs = map[byte]float32{
	byte('Z'): 0.90,
	byte('z'): 0.90,
	byte('Q'): 0.12,
	byte('q'): 0.12,
	byte('J'): 0.16,
	byte('j'): 0.16,
	byte('X'): 0.23,
	byte('x'): 0.23,
	byte('K'): 0.54,
	byte('k'): 0.54,
	byte('V'): 1.05,
	byte('v'): 1.05,
	byte('B'): 1.48,
	byte('b'): 1.48,
	byte('Y'): 1.66,
	byte('y'): 1.66,
	byte('W'): 1.68,
	byte('w'): 1.68,
	byte('G'): 1.87,
	byte('g'): 1.87,
	byte('P'): 2.14,
	byte('p'): 2.14,
	byte('F'): 2.40,
	byte('f'): 2.40,
	byte('M'): 2.51,
	byte('m'): 2.51,
	byte('U'): 2.73,
	byte('u'): 2.73,
	byte('C'): 3.34,
	byte('c'): 3.34,
	byte('D'): 3.82,
	byte('d'): 3.82,
	byte('L'): 4.07,
	byte('l'): 4.07,
	byte('H'): 5.05,
	byte('h'): 5.05,
	byte('R'): 6.28,
	byte('r'): 6.28,
	byte('S'): 6.51,
	byte('s'): 6.51,
	byte('N'): 7.23,
	byte('n'): 7.23,
	byte('I'): 7.57,
	byte('i'): 7.57,
	byte('O'): 7.64,
	byte('o'): 7.64,
	byte('A'): 8.04,
	byte('a'): 8.04,
	byte('T'): 9.28,
	byte('t'): 9.28,
	byte('E'): 12.49,
	byte('e'): 12.49,
}
