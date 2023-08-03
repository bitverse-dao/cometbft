
// Package lexer is generated by GoGLL. Do not edit.
package lexer

import (
	// "fmt"
	"io/ioutil"
	"strings"
	"unicode"

	"github.com/cometbft/cometbft/test/e2e/pkg/grammar/clean-start/token"
)

type state int

const nullState state = -1

// Unicode categories
var (
	_Cc     = unicode.Cc     // Cc is the set of Unicode characters in category Cc (Other, control).
	_Cf     = unicode.Cf     // Cf is the set of Unicode characters in category Cf (Other, format).
	_Co     = unicode.Co     // Co is the set of Unicode characters in category Co (Other, private use).
	_Cs     = unicode.Cs     // Cs is the set of Unicode characters in category Cs (Other, surrogate).
	_Digit  = unicode.Digit  // Digit is the set of Unicode characters with the "decimal digit" property.
	_Nd     = unicode.Nd     // Nd is the set of Unicode characters in category Nd (Number, decimal digit).
	_Letter = unicode.Letter // Letter/L is the set of Unicode letters, category L.
	_L      = unicode.L
	_Lm     = unicode.Lm    // Lm is the set of Unicode characters in category Lm (Letter, modifier).
	_Lo     = unicode.Lo    // Lo is the set of Unicode characters in category Lo (Letter, other).
	_Lower  = unicode.Lower // Lower is the set of Unicode lower case letters.
	_Ll     = unicode.Ll    // Ll is the set of Unicode characters in category Ll (Letter, lowercase).
	_Mark   = unicode.Mark  // Mark/M is the set of Unicode mark characters, category M.
	_M      = unicode.M
	_Mc     = unicode.Mc     // Mc is the set of Unicode characters in category Mc (Mark, spacing combining).
	_Me     = unicode.Me     // Me is the set of Unicode characters in category Me (Mark, enclosing).
	_Mn     = unicode.Mn     // Mn is the set of Unicode characters in category Mn (Mark, nonspacing).
	_Nl     = unicode.Nl     // Nl is the set of Unicode characters in category Nl (Number, letter).
	_No     = unicode.No     // No is the set of Unicode characters in category No (Number, other).
	_Number = unicode.Number // Number/N is the set of Unicode number characters, category N.
	_N      = unicode.N
	_Other  = unicode.Other // Other/C is the set of Unicode control and special characters, category C.
	_C      = unicode.C
	_Pc     = unicode.Pc    // Pc is the set of Unicode characters in category Pc (Punctuation, connector).
	_Pd     = unicode.Pd    // Pd is the set of Unicode characters in category Pd (Punctuation, dash).
	_Pe     = unicode.Pe    // Pe is the set of Unicode characters in category Pe (Punctuation, close).
	_Pf     = unicode.Pf    // Pf is the set of Unicode characters in category Pf (Punctuation, final quote).
	_Pi     = unicode.Pi    // Pi is the set of Unicode characters in category Pi (Punctuation, initial quote).
	_Po     = unicode.Po    // Po is the set of Unicode characters in category Po (Punctuation, other).
	_Ps     = unicode.Ps    // Ps is the set of Unicode characters in category Ps (Punctuation, open).
	_Punct  = unicode.Punct // Punct/P is the set of Unicode punctuation characters, category P.
	_P      = unicode.P
	_Sc     = unicode.Sc    // Sc is the set of Unicode characters in category Sc (Symbol, currency).
	_Sk     = unicode.Sk    // Sk is the set of Unicode characters in category Sk (Symbol, modifier).
	_Sm     = unicode.Sm    // Sm is the set of Unicode characters in category Sm (Symbol, math).
	_So     = unicode.So    // So is the set of Unicode characters in category So (Symbol, other).
	_Space  = unicode.Space // Space/Z is the set of Unicode space characters, category Z.
	_Z      = unicode.Z
	_Symbol = unicode.Symbol // Symbol/S is the set of Unicode symbol characters, category S.
	_S      = unicode.S
	_Title  = unicode.Title // Title is the set of Unicode title case letters.
	_Lt     = unicode.Lt    // Lt is the set of Unicode characters in category Lt (Letter, titlecase).
	_Upper  = unicode.Upper // Upper is the set of Unicode upper case letters.
	_Lu     = unicode.Lu    // Lu is the set of Unicode characters in category Lu (Letter, uppercase).
	_Zl     = unicode.Zl    // Zl is the set of Unicode characters in category Zl (Separator, line).
	_Zp     = unicode.Zp    // Zp is the set of Unicode characters in category Zp (Separator, paragraph).
	_Zs     = unicode.Zs    // Zs is the set of Unicode characters in category Zs (Separator, space).
)

// Unicode properties
var (
	_ASCII_Hex_Digit                    = unicode.ASCII_Hex_Digit                    // ASCII_Hex_Digit is the set of Unicode characters with property ASCII_Hex_Digit.
	_Bidi_Control                       = unicode.Bidi_Control                       // Bidi_Control is the set of Unicode characters with property Bidi_Control.
	_Dash                               = unicode.Dash                               // Dash is the set of Unicode characters with property Dash.
	_Deprecated                         = unicode.Deprecated                         // Deprecated is the set of Unicode characters with property Deprecated.
	_Diacritic                          = unicode.Diacritic                          // Diacritic is the set of Unicode characters with property Diacritic.
	_Extender                           = unicode.Extender                           // Extender is the set of Unicode characters with property Extender.
	_Hex_Digit                          = unicode.Hex_Digit                          // Hex_Digit is the set of Unicode characters with property Hex_Digit.
	_Hyphen                             = unicode.Hyphen                             // Hyphen is the set of Unicode characters with property Hyphen.
	_IDS_Binary_Operator                = unicode.IDS_Binary_Operator                // IDS_Binary_Operator is the set of Unicode characters with property IDS_Binary_Operator.
	_IDS_Trinary_Operator               = unicode.IDS_Trinary_Operator               // IDS_Trinary_Operator is the set of Unicode characters with property IDS_Trinary_Operator.
	_Ideographic                        = unicode.Ideographic                        // Ideographic is the set of Unicode characters with property Ideographic.
	_Join_Control                       = unicode.Join_Control                       // Join_Control is the set of Unicode characters with property Join_Control.
	_Logical_Order_Exception            = unicode.Logical_Order_Exception            // Logical_Order_Exception is the set of Unicode characters with property Logical_Order_Exception.
	_Noncharacter_Code_Point            = unicode.Noncharacter_Code_Point            // Noncharacter_Code_Point is the set of Unicode characters with property Noncharacter_Code_Point.
	_Other_Alphabetic                   = unicode.Other_Alphabetic                   // Other_Alphabetic is the set of Unicode characters with property Other_Alphabetic.
	_Other_Default_Ignorable_Code_Point = unicode.Other_Default_Ignorable_Code_Point // Other_Default_Ignorable_Code_Point is the set of Unicode characters with property Other_Default_Ignorable_Code_Point.
	_Other_Grapheme_Extend              = unicode.Other_Grapheme_Extend              // Other_Grapheme_Extend is the set of Unicode characters with property Other_Grapheme_Extend.
	_Other_ID_Continue                  = unicode.Other_ID_Continue                  // Other_ID_Continue is the set of Unicode characters with property Other_ID_Continue.
	_Other_ID_Start                     = unicode.Other_ID_Start                     // Other_ID_Start is the set of Unicode characters with property Other_ID_Start.
	_Other_Lowercase                    = unicode.Other_Lowercase                    // Other_Lowercase is the set of Unicode characters with property Other_Lowercase.
	_Other_Math                         = unicode.Other_Math                         // Other_Math is the set of Unicode characters with property Other_Math.
	_Other_Uppercase                    = unicode.Other_Uppercase                    // Other_Uppercase is the set of Unicode characters with property Other_Uppercase.
	_Pattern_Syntax                     = unicode.Pattern_Syntax                     // Pattern_Syntax is the set of Unicode characters with property Pattern_Syntax.
	_Pattern_White_Space                = unicode.Pattern_White_Space                // Pattern_White_Space is the set of Unicode characters with property Pattern_White_Space.
	_Prepended_Concatenation_Mark       = unicode.Prepended_Concatenation_Mark       // Prepended_Concatenation_Mark is the set of Unicode characters with property Prepended_Concatenation_Mark.
	_Quotation_Mark                     = unicode.Quotation_Mark                     // Quotation_Mark is the set of Unicode characters with property Quotation_Mark.
	_Radical                            = unicode.Radical                            // Radical is the set of Unicode characters with property Radical.
	_Regional_Indicator                 = unicode.Regional_Indicator                 // Regional_Indicator is the set of Unicode characters with property Regional_Indicator.
	_STerm                              = unicode.STerm                              // STerm is an alias for Sentence_Terminal.
	_Sentence_Terminal                  = unicode.Sentence_Terminal                  // Sentence_Terminal is the set of Unicode characters with property Sentence_Terminal.
	_Soft_Dotted                        = unicode.Soft_Dotted                        // Soft_Dotted is the set of Unicode characters with property Soft_Dotted.
	_Terminal_Punctuation               = unicode.Terminal_Punctuation               // Terminal_Punctuation is the set of Unicode characters with property Terminal_Punctuation.
	_Unified_Ideograph                  = unicode.Unified_Ideograph                  // Unified_Ideograph is the set of Unicode characters with property Unified_Ideograph.
	_Variation_Selector                 = unicode.Variation_Selector                 // Variation_Selector is the set of Unicode characters with property Variation_Selector.
	_White_Space                        = unicode.White_Space                        // White_Space is the set of Unicode characters with property White_Space.
)

// Lexer contains both the input slice of runes and the slice of tokens
// parsed from the input
type Lexer struct {
	// I is the input slice of runes
	I      []rune

	// Tokens is the slice of tokens constructed by the lexer from I
	Tokens []*token.Token
}

/*
NewFile constructs a Lexer created from the input file, fname. 

If the input file is a markdown file NewFile process treats all text outside
code blocks as whitespace. All text inside code blocks are treated as input text.

If the input file is a normal text file NewFile treats all text in the inputfile
as input text.
*/
func NewFile(fname string) *Lexer {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	input := []rune(string(buf))
	if strings.HasSuffix(fname, ".md") {
		loadMd(input)
	}
	return New(input)
}

func loadMd(input []rune) {
	i := 0
	text := true
	for i < len(input) {
		if i <= len(input)-3 && input[i] == '`' && input[i+1] == '`' && input[i+2] == '`' {
			text = !text
			for j := 0; j < 3; j++ {
				input[i+j] = ' '
			}
			i += 3
		}
		if i < len(input) {
			if text {
				if input[i] == '\n' {
					input[i] = '\n'
				} else {
					input[i] = ' '
				}
			}
			i += 1
		}
	}
}

/*
New constructs a Lexer from a slice of runes. 

All contents of the input slice are treated as input text.
*/
func New(input []rune) *Lexer {
	lex := &Lexer{
		I:      input,
		Tokens: make([]*token.Token, 0, 2048),
	}
	lext := 0
	for lext < len(lex.I) {
		for lext < len(lex.I) && unicode.IsSpace(lex.I[lext]) {
			lext++
		}
		if lext < len(lex.I) {
			tok := lex.scan(lext)
			lext = tok.Rext()
			if !tok.Suppress() {
				lex.addToken(tok)
			}
		}
	}
	lex.add(token.EOF, len(input), len(input))
	return lex
}

func (l *Lexer) scan(i int) *token.Token {
	// fmt.Printf("lexer.scan(%d)\n", i)
	s, typ, rext := nullState, token.Error, i+1
	if i < len(l.I) {
		// fmt.Printf("  rext %d, i %d\n", rext, i)
		s = nextState[0](l.I[i])
	}
	for s != nullState {
		if rext >= len(l.I) {
			typ = accept[s]
			s = nullState
		} else {
			typ = accept[s]
			s = nextState[s](l.I[rext])
			if s != nullState || typ == token.Error {
				rext++
			}
		}
	}
	tok := token.New(typ, i, rext, l.I)
	// fmt.Printf("  %s\n", tok)
	return tok
}

func escape(r rune) string {
	switch r {
	case '"':
		return "\""
	case '\\':
		return "\\\\"
	case '\r':
		return "\\r"
	case '\n':
		return "\\n"
	case '\t':
		return "\\t"
	}
	return string(r)
}

// GetLineColumn returns the line and column of rune[i] in the input
func (l *Lexer) GetLineColumn(i int) (line, col int) {
	line, col = 1, 1
	for j := 0; j < i; j++ {
		switch l.I[j] {
		case '\n':
			line++
			col = 1
		case '\t':
			col += 4
		default:
			col++
		}
	}
	return
}

// GetLineColumnOfToken returns the line and column of token[i] in the imput
func (l *Lexer) GetLineColumnOfToken(i int) (line, col int) {
	return l.GetLineColumn(l.Tokens[i].Lext())
}

// GetString returns the input string from the left extent of Token[lext] to
// the right extent of Token[rext]
func (l *Lexer) GetString(lext, rext int) string {
	return string(l.I[l.Tokens[lext].Lext():l.Tokens[rext].Rext()])
}

func (l *Lexer) add(t token.Type, lext, rext int) {
	l.addToken(token.New(t, lext, rext, l.I))
}

func (l *Lexer) addToken(tok *token.Token) {
	l.Tokens = append(l.Tokens, tok)
}

func any(r rune, set []rune) bool {
	for _, r1 := range set {
		if r == r1 {
			return true
		}
	}
	return false
}

func not(r rune, set []rune) bool {
	for _, r1 := range set {
		if r == r1 {
			return false
		}
	}
	return true
}

var accept = []token.Type{ 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.T_1, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.T_3, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.T_2, 
	token.T_4, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.T_5, 
	token.T_6, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.T_0, 
}

var nextState = []func(r rune) state{ 
	// Set0
	func(r rune) state {
		switch { 
		case r == 'a':
			return 1 
		case r == 'c':
			return 2 
		case r == 'f':
			return 3 
		case r == 'i':
			return 4 
		case r == 'o':
			return 5 
		case r == 'p':
			return 6 
		}
		return nullState
	}, 
	// Set1
	func(r rune) state {
		switch { 
		case r == 'p':
			return 7 
		}
		return nullState
	}, 
	// Set2
	func(r rune) state {
		switch { 
		case r == 'o':
			return 8 
		}
		return nullState
	}, 
	// Set3
	func(r rune) state {
		switch { 
		case r == 'i':
			return 9 
		}
		return nullState
	}, 
	// Set4
	func(r rune) state {
		switch { 
		case r == 'n':
			return 10 
		}
		return nullState
	}, 
	// Set5
	func(r rune) state {
		switch { 
		case r == 'f':
			return 11 
		}
		return nullState
	}, 
	// Set6
	func(r rune) state {
		switch { 
		case r == 'r':
			return 12 
		}
		return nullState
	}, 
	// Set7
	func(r rune) state {
		switch { 
		case r == 'p':
			return 13 
		}
		return nullState
	}, 
	// Set8
	func(r rune) state {
		switch { 
		case r == 'm':
			return 14 
		}
		return nullState
	}, 
	// Set9
	func(r rune) state {
		switch { 
		case r == 'n':
			return 15 
		}
		return nullState
	}, 
	// Set10
	func(r rune) state {
		switch { 
		case r == 'i':
			return 16 
		}
		return nullState
	}, 
	// Set11
	func(r rune) state {
		switch { 
		case r == 'f':
			return 17 
		}
		return nullState
	}, 
	// Set12
	func(r rune) state {
		switch { 
		case r == 'e':
			return 18 
		case r == 'o':
			return 19 
		}
		return nullState
	}, 
	// Set13
	func(r rune) state {
		switch { 
		case r == 'l':
			return 20 
		}
		return nullState
	}, 
	// Set14
	func(r rune) state {
		switch { 
		case r == 'm':
			return 21 
		}
		return nullState
	}, 
	// Set15
	func(r rune) state {
		switch { 
		case r == 'a':
			return 22 
		}
		return nullState
	}, 
	// Set16
	func(r rune) state {
		switch { 
		case r == 't':
			return 23 
		}
		return nullState
	}, 
	// Set17
	func(r rune) state {
		switch { 
		case r == 'e':
			return 24 
		}
		return nullState
	}, 
	// Set18
	func(r rune) state {
		switch { 
		case r == 'p':
			return 25 
		}
		return nullState
	}, 
	// Set19
	func(r rune) state {
		switch { 
		case r == 'c':
			return 26 
		}
		return nullState
	}, 
	// Set20
	func(r rune) state {
		switch { 
		case r == 'y':
			return 27 
		}
		return nullState
	}, 
	// Set21
	func(r rune) state {
		switch { 
		case r == 'i':
			return 28 
		}
		return nullState
	}, 
	// Set22
	func(r rune) state {
		switch { 
		case r == 'l':
			return 29 
		}
		return nullState
	}, 
	// Set23
	func(r rune) state {
		switch { 
		case r == '_':
			return 30 
		}
		return nullState
	}, 
	// Set24
	func(r rune) state {
		switch { 
		case r == 'r':
			return 31 
		}
		return nullState
	}, 
	// Set25
	func(r rune) state {
		switch { 
		case r == 'a':
			return 32 
		}
		return nullState
	}, 
	// Set26
	func(r rune) state {
		switch { 
		case r == 'e':
			return 33 
		}
		return nullState
	}, 
	// Set27
	func(r rune) state {
		switch { 
		case r == '_':
			return 34 
		}
		return nullState
	}, 
	// Set28
	func(r rune) state {
		switch { 
		case r == 't':
			return 35 
		}
		return nullState
	}, 
	// Set29
	func(r rune) state {
		switch { 
		case r == 'i':
			return 36 
		}
		return nullState
	}, 
	// Set30
	func(r rune) state {
		switch { 
		case r == 'c':
			return 37 
		}
		return nullState
	}, 
	// Set31
	func(r rune) state {
		switch { 
		case r == '_':
			return 38 
		}
		return nullState
	}, 
	// Set32
	func(r rune) state {
		switch { 
		case r == 'r':
			return 39 
		}
		return nullState
	}, 
	// Set33
	func(r rune) state {
		switch { 
		case r == 's':
			return 40 
		}
		return nullState
	}, 
	// Set34
	func(r rune) state {
		switch { 
		case r == 's':
			return 41 
		}
		return nullState
	}, 
	// Set35
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set36
	func(r rune) state {
		switch { 
		case r == 'z':
			return 42 
		}
		return nullState
	}, 
	// Set37
	func(r rune) state {
		switch { 
		case r == 'h':
			return 43 
		}
		return nullState
	}, 
	// Set38
	func(r rune) state {
		switch { 
		case r == 's':
			return 44 
		}
		return nullState
	}, 
	// Set39
	func(r rune) state {
		switch { 
		case r == 'e':
			return 45 
		}
		return nullState
	}, 
	// Set40
	func(r rune) state {
		switch { 
		case r == 's':
			return 46 
		}
		return nullState
	}, 
	// Set41
	func(r rune) state {
		switch { 
		case r == 'n':
			return 47 
		}
		return nullState
	}, 
	// Set42
	func(r rune) state {
		switch { 
		case r == 'e':
			return 48 
		}
		return nullState
	}, 
	// Set43
	func(r rune) state {
		switch { 
		case r == 'a':
			return 49 
		}
		return nullState
	}, 
	// Set44
	func(r rune) state {
		switch { 
		case r == 'n':
			return 50 
		}
		return nullState
	}, 
	// Set45
	func(r rune) state {
		switch { 
		case r == '_':
			return 51 
		}
		return nullState
	}, 
	// Set46
	func(r rune) state {
		switch { 
		case r == '_':
			return 52 
		}
		return nullState
	}, 
	// Set47
	func(r rune) state {
		switch { 
		case r == 'a':
			return 53 
		}
		return nullState
	}, 
	// Set48
	func(r rune) state {
		switch { 
		case r == '_':
			return 54 
		}
		return nullState
	}, 
	// Set49
	func(r rune) state {
		switch { 
		case r == 'i':
			return 55 
		}
		return nullState
	}, 
	// Set50
	func(r rune) state {
		switch { 
		case r == 'a':
			return 56 
		}
		return nullState
	}, 
	// Set51
	func(r rune) state {
		switch { 
		case r == 'p':
			return 57 
		}
		return nullState
	}, 
	// Set52
	func(r rune) state {
		switch { 
		case r == 'p':
			return 58 
		}
		return nullState
	}, 
	// Set53
	func(r rune) state {
		switch { 
		case r == 'p':
			return 59 
		}
		return nullState
	}, 
	// Set54
	func(r rune) state {
		switch { 
		case r == 'b':
			return 60 
		}
		return nullState
	}, 
	// Set55
	func(r rune) state {
		switch { 
		case r == 'n':
			return 61 
		}
		return nullState
	}, 
	// Set56
	func(r rune) state {
		switch { 
		case r == 'p':
			return 62 
		}
		return nullState
	}, 
	// Set57
	func(r rune) state {
		switch { 
		case r == 'r':
			return 63 
		}
		return nullState
	}, 
	// Set58
	func(r rune) state {
		switch { 
		case r == 'r':
			return 64 
		}
		return nullState
	}, 
	// Set59
	func(r rune) state {
		switch { 
		case r == 's':
			return 65 
		}
		return nullState
	}, 
	// Set60
	func(r rune) state {
		switch { 
		case r == 'l':
			return 66 
		}
		return nullState
	}, 
	// Set61
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set62
	func(r rune) state {
		switch { 
		case r == 's':
			return 67 
		}
		return nullState
	}, 
	// Set63
	func(r rune) state {
		switch { 
		case r == 'o':
			return 68 
		}
		return nullState
	}, 
	// Set64
	func(r rune) state {
		switch { 
		case r == 'o':
			return 69 
		}
		return nullState
	}, 
	// Set65
	func(r rune) state {
		switch { 
		case r == 'h':
			return 70 
		}
		return nullState
	}, 
	// Set66
	func(r rune) state {
		switch { 
		case r == 'o':
			return 71 
		}
		return nullState
	}, 
	// Set67
	func(r rune) state {
		switch { 
		case r == 'h':
			return 72 
		}
		return nullState
	}, 
	// Set68
	func(r rune) state {
		switch { 
		case r == 'p':
			return 73 
		}
		return nullState
	}, 
	// Set69
	func(r rune) state {
		switch { 
		case r == 'p':
			return 74 
		}
		return nullState
	}, 
	// Set70
	func(r rune) state {
		switch { 
		case r == 'o':
			return 75 
		}
		return nullState
	}, 
	// Set71
	func(r rune) state {
		switch { 
		case r == 'c':
			return 76 
		}
		return nullState
	}, 
	// Set72
	func(r rune) state {
		switch { 
		case r == 'o':
			return 77 
		}
		return nullState
	}, 
	// Set73
	func(r rune) state {
		switch { 
		case r == 'o':
			return 78 
		}
		return nullState
	}, 
	// Set74
	func(r rune) state {
		switch { 
		case r == 'o':
			return 79 
		}
		return nullState
	}, 
	// Set75
	func(r rune) state {
		switch { 
		case r == 't':
			return 80 
		}
		return nullState
	}, 
	// Set76
	func(r rune) state {
		switch { 
		case r == 'k':
			return 81 
		}
		return nullState
	}, 
	// Set77
	func(r rune) state {
		switch { 
		case r == 't':
			return 82 
		}
		return nullState
	}, 
	// Set78
	func(r rune) state {
		switch { 
		case r == 's':
			return 83 
		}
		return nullState
	}, 
	// Set79
	func(r rune) state {
		switch { 
		case r == 's':
			return 84 
		}
		return nullState
	}, 
	// Set80
	func(r rune) state {
		switch { 
		case r == '_':
			return 85 
		}
		return nullState
	}, 
	// Set81
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set82
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set83
	func(r rune) state {
		switch { 
		case r == 'a':
			return 86 
		}
		return nullState
	}, 
	// Set84
	func(r rune) state {
		switch { 
		case r == 'a':
			return 87 
		}
		return nullState
	}, 
	// Set85
	func(r rune) state {
		switch { 
		case r == 'c':
			return 88 
		}
		return nullState
	}, 
	// Set86
	func(r rune) state {
		switch { 
		case r == 'l':
			return 89 
		}
		return nullState
	}, 
	// Set87
	func(r rune) state {
		switch { 
		case r == 'l':
			return 90 
		}
		return nullState
	}, 
	// Set88
	func(r rune) state {
		switch { 
		case r == 'h':
			return 91 
		}
		return nullState
	}, 
	// Set89
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set90
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set91
	func(r rune) state {
		switch { 
		case r == 'u':
			return 92 
		}
		return nullState
	}, 
	// Set92
	func(r rune) state {
		switch { 
		case r == 'n':
			return 93 
		}
		return nullState
	}, 
	// Set93
	func(r rune) state {
		switch { 
		case r == 'k':
			return 94 
		}
		return nullState
	}, 
	// Set94
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
}
