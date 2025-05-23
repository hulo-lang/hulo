// Code generated from huloLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type huloLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var HuloLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hulolexerLexerInit() {
	staticData := &HuloLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'mod'", "'use'", "'package'", "'import'", "'from'", "'type'", "'typeof'",
		"'as'", "'if'", "'else'", "'match'", "'do'", "'loop'", "'in'", "'range'",
		"'continue'", "'break'", "'let'", "'var'", "'const'", "'static'", "'final'",
		"'pub'", "'required'", "'factory'", "'try'", "'catch'", "'finally'",
		"'throw'", "'throws'", "'fn'", "'operator'", "'return'", "'enum'", "'class'",
		"'trait'", "'impl'", "'for'", "'this'", "'super'", "'extend'", "'declare'",
		"'defer'", "'comptime'", "'unsafe'", "'.'", "'..'", "','", "'('", "')'",
		"'['", "']'", "'{'", "'}'", "'**'", "'*'", "'%'", "'/'", "'+'", "'-'",
		"'#'", "'@'", "'?'", "'$'", "'\\'", "'_'", "'='", "'+='", "'-='", "'*='",
		"'**='", "'/='", "'%='", "'&&='", "'<'", "'>'", "'<='", "'>='", "'!='",
		"'=='", "'<<'", "'>>'", "'->'", "'<-'", "'=>'", "'...'", "':'", "'::'",
		"';'", "'&&'", "'||'", "'!'", "'++'", "'--'", "'&'", "'|'", "'^'", "'''",
		"'\"'", "'\"\"\"'", "'new'", "'delete'", "'null'", "'num'", "'str'",
		"'bool'", "'any'", "", "'[['", "']]'",
	}
	staticData.SymbolicNames = []string{
		"", "MOD_LIT", "USE", "PACKAGE", "IMPORT", "FROM", "TYPE", "TYPEOF",
		"AS", "IF", "ELSE", "MATCH", "DO", "LOOP", "IN", "RANGE", "CONTINUE",
		"BREAK", "LET", "VAR", "CONST", "STATIC", "FINAL", "PUB", "REQUIRED",
		"FACTORY", "TRY", "CATCH", "FINALLY", "THROW", "THROWS", "FN", "OPERATOR",
		"RETURN", "ENUM", "CLASS", "TRAIT", "IMPL", "FOR", "THIS", "SUPER",
		"EXTEND", "DECLARE", "DEFER", "COMPTIME", "UNSAFE", "DOT", "DOUBLE_DOT",
		"COMMA", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE",
		"EXP", "MUL", "MOD", "DIV", "ADD", "SUB", "HASH", "AT", "QUEST", "DOLLAR",
		"BACKSLASH", "WILDCARD", "ASSIGN", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN",
		"EXP_ASSIGN", "DIV_ASSIGN", "MOD_ASSIGN", "AND_ASSIGN", "LT", "GT",
		"LE", "GE", "NEQ", "EQ", "SHL", "SHR", "ARROW", "BACKARROW", "DOUBLE_ARROW",
		"ELLIPSIS", "COLON", "DOUBLE_COLON", "SEMI", "AND", "OR", "NOT", "INC",
		"DEC", "BITAND", "BITOR", "BITXOR", "SINGLE_QUOTE", "QUOTE", "TRIPLE_QUOTE",
		"NEW", "DELETE", "NULL", "NUM", "STR", "BOOL", "ANY", "UnsafeLiteral",
		"DOUBLE_LBRACK", "DOUBLE_RBRACK", "NumberLiteral", "StringLiteral",
		"BoolLiteral", "RawStringLiteral", "FileStringLiteral", "CommandStringLiteral",
		"ESC", "LineComment", "BlockComment", "Identifier", "WS",
	}
	staticData.RuleNames = []string{
		"MOD_LIT", "USE", "PACKAGE", "IMPORT", "FROM", "TYPE", "TYPEOF", "AS",
		"IF", "ELSE", "MATCH", "DO", "LOOP", "IN", "RANGE", "CONTINUE", "BREAK",
		"LET", "VAR", "CONST", "STATIC", "FINAL", "PUB", "REQUIRED", "FACTORY",
		"TRY", "CATCH", "FINALLY", "THROW", "THROWS", "FN", "OPERATOR", "RETURN",
		"ENUM", "CLASS", "TRAIT", "IMPL", "FOR", "THIS", "SUPER", "EXTEND",
		"DECLARE", "DEFER", "COMPTIME", "UNSAFE", "DOT", "DOUBLE_DOT", "COMMA",
		"LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "EXP", "MUL",
		"MOD", "DIV", "ADD", "SUB", "HASH", "AT", "QUEST", "DOLLAR", "BACKSLASH",
		"WILDCARD", "ASSIGN", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN", "EXP_ASSIGN",
		"DIV_ASSIGN", "MOD_ASSIGN", "AND_ASSIGN", "LT", "GT", "LE", "GE", "NEQ",
		"EQ", "SHL", "SHR", "ARROW", "BACKARROW", "DOUBLE_ARROW", "ELLIPSIS",
		"COLON", "DOUBLE_COLON", "SEMI", "AND", "OR", "NOT", "INC", "DEC", "BITAND",
		"BITOR", "BITXOR", "SINGLE_QUOTE", "QUOTE", "TRIPLE_QUOTE", "NEW", "DELETE",
		"NULL", "NUM", "STR", "BOOL", "ANY", "UnsafeLiteral", "DOUBLE_LBRACK",
		"DOUBLE_RBRACK", "TRUE", "FALSE", "NumberLiteral", "StringLiteral",
		"BoolLiteral", "RawStringLiteral", "FileStringLiteral", "CommandStringLiteral",
		"ESC", "LineComment", "BlockComment", "Identifier", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 121, 835, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3,
		2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9,
		2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2,
		15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20,
		7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7,
		25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30,
		2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2,
		36, 7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41,
		7, 41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7,
		46, 2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51,
		2, 52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2,
		57, 7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62,
		7, 62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7,
		67, 2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72,
		2, 73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2,
		78, 7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83,
		7, 83, 2, 84, 7, 84, 2, 85, 7, 85, 2, 86, 7, 86, 2, 87, 7, 87, 2, 88, 7,
		88, 2, 89, 7, 89, 2, 90, 7, 90, 2, 91, 7, 91, 2, 92, 7, 92, 2, 93, 7, 93,
		2, 94, 7, 94, 2, 95, 7, 95, 2, 96, 7, 96, 2, 97, 7, 97, 2, 98, 7, 98, 2,
		99, 7, 99, 2, 100, 7, 100, 2, 101, 7, 101, 2, 102, 7, 102, 2, 103, 7, 103,
		2, 104, 7, 104, 2, 105, 7, 105, 2, 106, 7, 106, 2, 107, 7, 107, 2, 108,
		7, 108, 2, 109, 7, 109, 2, 110, 7, 110, 2, 111, 7, 111, 2, 112, 7, 112,
		2, 113, 7, 113, 2, 114, 7, 114, 2, 115, 7, 115, 2, 116, 7, 116, 2, 117,
		7, 117, 2, 118, 7, 118, 2, 119, 7, 119, 2, 120, 7, 120, 2, 121, 7, 121,
		2, 122, 7, 122, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1,
		47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52,
		1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1,
		57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62,
		1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66, 1, 66, 1, 67, 1,
		67, 1, 67, 1, 68, 1, 68, 1, 68, 1, 69, 1, 69, 1, 69, 1, 70, 1, 70, 1, 70,
		1, 70, 1, 71, 1, 71, 1, 71, 1, 72, 1, 72, 1, 72, 1, 73, 1, 73, 1, 73, 1,
		73, 1, 74, 1, 74, 1, 75, 1, 75, 1, 76, 1, 76, 1, 76, 1, 77, 1, 77, 1, 77,
		1, 78, 1, 78, 1, 78, 1, 79, 1, 79, 1, 79, 1, 80, 1, 80, 1, 80, 1, 81, 1,
		81, 1, 81, 1, 82, 1, 82, 1, 82, 1, 83, 1, 83, 1, 83, 1, 84, 1, 84, 1, 84,
		1, 85, 1, 85, 1, 85, 1, 85, 1, 86, 1, 86, 1, 87, 1, 87, 1, 87, 1, 88, 1,
		88, 1, 89, 1, 89, 1, 89, 1, 90, 1, 90, 1, 90, 1, 91, 1, 91, 1, 92, 1, 92,
		1, 92, 1, 93, 1, 93, 1, 93, 1, 94, 1, 94, 1, 95, 1, 95, 1, 96, 1, 96, 1,
		97, 1, 97, 1, 98, 1, 98, 1, 99, 1, 99, 1, 99, 1, 99, 1, 100, 1, 100, 1,
		100, 1, 100, 1, 101, 1, 101, 1, 101, 1, 101, 1, 101, 1, 101, 1, 101, 1,
		102, 1, 102, 1, 102, 1, 102, 1, 102, 1, 103, 1, 103, 1, 103, 1, 103, 1,
		104, 1, 104, 1, 104, 1, 104, 1, 105, 1, 105, 1, 105, 1, 105, 1, 105, 1,
		106, 1, 106, 1, 106, 1, 106, 1, 107, 1, 107, 1, 107, 1, 107, 1, 107, 5,
		107, 686, 8, 107, 10, 107, 12, 107, 689, 9, 107, 1, 107, 1, 107, 1, 108,
		1, 108, 1, 108, 1, 109, 1, 109, 1, 109, 1, 110, 1, 110, 1, 110, 1, 110,
		1, 110, 1, 111, 1, 111, 1, 111, 1, 111, 1, 111, 1, 111, 1, 112, 3, 112,
		711, 8, 112, 1, 112, 4, 112, 714, 8, 112, 11, 112, 12, 112, 715, 1, 112,
		1, 112, 4, 112, 720, 8, 112, 11, 112, 12, 112, 721, 3, 112, 724, 8, 112,
		1, 113, 1, 113, 1, 113, 5, 113, 729, 8, 113, 10, 113, 12, 113, 732, 9,
		113, 1, 113, 1, 113, 1, 113, 1, 113, 1, 113, 5, 113, 739, 8, 113, 10, 113,
		12, 113, 742, 9, 113, 1, 113, 1, 113, 1, 113, 1, 113, 5, 113, 748, 8, 113,
		10, 113, 12, 113, 751, 9, 113, 1, 113, 1, 113, 1, 113, 3, 113, 756, 8,
		113, 1, 114, 1, 114, 3, 114, 760, 8, 114, 1, 115, 1, 115, 1, 115, 1, 115,
		1, 115, 5, 115, 767, 8, 115, 10, 115, 12, 115, 770, 9, 115, 1, 115, 1,
		115, 1, 116, 1, 116, 1, 116, 1, 116, 1, 116, 5, 116, 779, 8, 116, 10, 116,
		12, 116, 782, 9, 116, 1, 116, 1, 116, 1, 117, 1, 117, 1, 117, 1, 117, 1,
		117, 5, 117, 791, 8, 117, 10, 117, 12, 117, 794, 9, 117, 1, 117, 1, 117,
		1, 118, 1, 118, 1, 118, 1, 119, 1, 119, 1, 119, 1, 119, 5, 119, 805, 8,
		119, 10, 119, 12, 119, 808, 9, 119, 1, 120, 1, 120, 1, 120, 1, 120, 5,
		120, 814, 8, 120, 10, 120, 12, 120, 817, 9, 120, 1, 120, 1, 120, 1, 120,
		1, 121, 1, 121, 5, 121, 824, 8, 121, 10, 121, 12, 121, 827, 9, 121, 1,
		122, 4, 122, 830, 8, 122, 11, 122, 12, 122, 831, 1, 122, 1, 122, 2, 749,
		815, 0, 123, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27,
		55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36,
		73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45,
		91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107,
		54, 109, 55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 123,
		62, 125, 63, 127, 64, 129, 65, 131, 66, 133, 67, 135, 68, 137, 69, 139,
		70, 141, 71, 143, 72, 145, 73, 147, 74, 149, 75, 151, 76, 153, 77, 155,
		78, 157, 79, 159, 80, 161, 81, 163, 82, 165, 83, 167, 84, 169, 85, 171,
		86, 173, 87, 175, 88, 177, 89, 179, 90, 181, 91, 183, 92, 185, 93, 187,
		94, 189, 95, 191, 96, 193, 97, 195, 98, 197, 99, 199, 100, 201, 101, 203,
		102, 205, 103, 207, 104, 209, 105, 211, 106, 213, 107, 215, 108, 217, 109,
		219, 110, 221, 0, 223, 0, 225, 111, 227, 112, 229, 113, 231, 114, 233,
		115, 235, 116, 237, 117, 239, 118, 241, 119, 243, 120, 245, 121, 1, 0,
		9, 1, 0, 93, 93, 1, 0, 48, 57, 3, 0, 10, 10, 39, 39, 92, 92, 3, 0, 10,
		10, 34, 34, 92, 92, 8, 0, 34, 34, 39, 39, 92, 92, 98, 98, 102, 102, 110,
		110, 114, 114, 116, 116, 1, 0, 10, 10, 3, 0, 65, 90, 95, 95, 97, 122, 4,
		0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 857, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47,
		1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0,
		55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0,
		0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0,
		0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0,
		0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1,
		0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93,
		1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0,
		101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0,
		0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115,
		1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0,
		0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1,
		0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 0,
		137, 1, 0, 0, 0, 0, 139, 1, 0, 0, 0, 0, 141, 1, 0, 0, 0, 0, 143, 1, 0,
		0, 0, 0, 145, 1, 0, 0, 0, 0, 147, 1, 0, 0, 0, 0, 149, 1, 0, 0, 0, 0, 151,
		1, 0, 0, 0, 0, 153, 1, 0, 0, 0, 0, 155, 1, 0, 0, 0, 0, 157, 1, 0, 0, 0,
		0, 159, 1, 0, 0, 0, 0, 161, 1, 0, 0, 0, 0, 163, 1, 0, 0, 0, 0, 165, 1,
		0, 0, 0, 0, 167, 1, 0, 0, 0, 0, 169, 1, 0, 0, 0, 0, 171, 1, 0, 0, 0, 0,
		173, 1, 0, 0, 0, 0, 175, 1, 0, 0, 0, 0, 177, 1, 0, 0, 0, 0, 179, 1, 0,
		0, 0, 0, 181, 1, 0, 0, 0, 0, 183, 1, 0, 0, 0, 0, 185, 1, 0, 0, 0, 0, 187,
		1, 0, 0, 0, 0, 189, 1, 0, 0, 0, 0, 191, 1, 0, 0, 0, 0, 193, 1, 0, 0, 0,
		0, 195, 1, 0, 0, 0, 0, 197, 1, 0, 0, 0, 0, 199, 1, 0, 0, 0, 0, 201, 1,
		0, 0, 0, 0, 203, 1, 0, 0, 0, 0, 205, 1, 0, 0, 0, 0, 207, 1, 0, 0, 0, 0,
		209, 1, 0, 0, 0, 0, 211, 1, 0, 0, 0, 0, 213, 1, 0, 0, 0, 0, 215, 1, 0,
		0, 0, 0, 217, 1, 0, 0, 0, 0, 219, 1, 0, 0, 0, 0, 225, 1, 0, 0, 0, 0, 227,
		1, 0, 0, 0, 0, 229, 1, 0, 0, 0, 0, 231, 1, 0, 0, 0, 0, 233, 1, 0, 0, 0,
		0, 235, 1, 0, 0, 0, 0, 237, 1, 0, 0, 0, 0, 239, 1, 0, 0, 0, 0, 241, 1,
		0, 0, 0, 0, 243, 1, 0, 0, 0, 0, 245, 1, 0, 0, 0, 1, 247, 1, 0, 0, 0, 3,
		251, 1, 0, 0, 0, 5, 255, 1, 0, 0, 0, 7, 263, 1, 0, 0, 0, 9, 270, 1, 0,
		0, 0, 11, 275, 1, 0, 0, 0, 13, 280, 1, 0, 0, 0, 15, 287, 1, 0, 0, 0, 17,
		290, 1, 0, 0, 0, 19, 293, 1, 0, 0, 0, 21, 298, 1, 0, 0, 0, 23, 304, 1,
		0, 0, 0, 25, 307, 1, 0, 0, 0, 27, 312, 1, 0, 0, 0, 29, 315, 1, 0, 0, 0,
		31, 321, 1, 0, 0, 0, 33, 330, 1, 0, 0, 0, 35, 336, 1, 0, 0, 0, 37, 340,
		1, 0, 0, 0, 39, 344, 1, 0, 0, 0, 41, 350, 1, 0, 0, 0, 43, 357, 1, 0, 0,
		0, 45, 363, 1, 0, 0, 0, 47, 367, 1, 0, 0, 0, 49, 376, 1, 0, 0, 0, 51, 384,
		1, 0, 0, 0, 53, 388, 1, 0, 0, 0, 55, 394, 1, 0, 0, 0, 57, 402, 1, 0, 0,
		0, 59, 408, 1, 0, 0, 0, 61, 415, 1, 0, 0, 0, 63, 418, 1, 0, 0, 0, 65, 427,
		1, 0, 0, 0, 67, 434, 1, 0, 0, 0, 69, 439, 1, 0, 0, 0, 71, 445, 1, 0, 0,
		0, 73, 451, 1, 0, 0, 0, 75, 456, 1, 0, 0, 0, 77, 460, 1, 0, 0, 0, 79, 465,
		1, 0, 0, 0, 81, 471, 1, 0, 0, 0, 83, 478, 1, 0, 0, 0, 85, 486, 1, 0, 0,
		0, 87, 492, 1, 0, 0, 0, 89, 501, 1, 0, 0, 0, 91, 508, 1, 0, 0, 0, 93, 510,
		1, 0, 0, 0, 95, 513, 1, 0, 0, 0, 97, 515, 1, 0, 0, 0, 99, 517, 1, 0, 0,
		0, 101, 519, 1, 0, 0, 0, 103, 521, 1, 0, 0, 0, 105, 523, 1, 0, 0, 0, 107,
		525, 1, 0, 0, 0, 109, 527, 1, 0, 0, 0, 111, 530, 1, 0, 0, 0, 113, 532,
		1, 0, 0, 0, 115, 534, 1, 0, 0, 0, 117, 536, 1, 0, 0, 0, 119, 538, 1, 0,
		0, 0, 121, 540, 1, 0, 0, 0, 123, 542, 1, 0, 0, 0, 125, 544, 1, 0, 0, 0,
		127, 546, 1, 0, 0, 0, 129, 548, 1, 0, 0, 0, 131, 550, 1, 0, 0, 0, 133,
		552, 1, 0, 0, 0, 135, 554, 1, 0, 0, 0, 137, 557, 1, 0, 0, 0, 139, 560,
		1, 0, 0, 0, 141, 563, 1, 0, 0, 0, 143, 567, 1, 0, 0, 0, 145, 570, 1, 0,
		0, 0, 147, 573, 1, 0, 0, 0, 149, 577, 1, 0, 0, 0, 151, 579, 1, 0, 0, 0,
		153, 581, 1, 0, 0, 0, 155, 584, 1, 0, 0, 0, 157, 587, 1, 0, 0, 0, 159,
		590, 1, 0, 0, 0, 161, 593, 1, 0, 0, 0, 163, 596, 1, 0, 0, 0, 165, 599,
		1, 0, 0, 0, 167, 602, 1, 0, 0, 0, 169, 605, 1, 0, 0, 0, 171, 608, 1, 0,
		0, 0, 173, 612, 1, 0, 0, 0, 175, 614, 1, 0, 0, 0, 177, 617, 1, 0, 0, 0,
		179, 619, 1, 0, 0, 0, 181, 622, 1, 0, 0, 0, 183, 625, 1, 0, 0, 0, 185,
		627, 1, 0, 0, 0, 187, 630, 1, 0, 0, 0, 189, 633, 1, 0, 0, 0, 191, 635,
		1, 0, 0, 0, 193, 637, 1, 0, 0, 0, 195, 639, 1, 0, 0, 0, 197, 641, 1, 0,
		0, 0, 199, 643, 1, 0, 0, 0, 201, 647, 1, 0, 0, 0, 203, 651, 1, 0, 0, 0,
		205, 658, 1, 0, 0, 0, 207, 663, 1, 0, 0, 0, 209, 667, 1, 0, 0, 0, 211,
		671, 1, 0, 0, 0, 213, 676, 1, 0, 0, 0, 215, 680, 1, 0, 0, 0, 217, 692,
		1, 0, 0, 0, 219, 695, 1, 0, 0, 0, 221, 698, 1, 0, 0, 0, 223, 703, 1, 0,
		0, 0, 225, 710, 1, 0, 0, 0, 227, 755, 1, 0, 0, 0, 229, 759, 1, 0, 0, 0,
		231, 761, 1, 0, 0, 0, 233, 773, 1, 0, 0, 0, 235, 785, 1, 0, 0, 0, 237,
		797, 1, 0, 0, 0, 239, 800, 1, 0, 0, 0, 241, 809, 1, 0, 0, 0, 243, 821,
		1, 0, 0, 0, 245, 829, 1, 0, 0, 0, 247, 248, 5, 109, 0, 0, 248, 249, 5,
		111, 0, 0, 249, 250, 5, 100, 0, 0, 250, 2, 1, 0, 0, 0, 251, 252, 5, 117,
		0, 0, 252, 253, 5, 115, 0, 0, 253, 254, 5, 101, 0, 0, 254, 4, 1, 0, 0,
		0, 255, 256, 5, 112, 0, 0, 256, 257, 5, 97, 0, 0, 257, 258, 5, 99, 0, 0,
		258, 259, 5, 107, 0, 0, 259, 260, 5, 97, 0, 0, 260, 261, 5, 103, 0, 0,
		261, 262, 5, 101, 0, 0, 262, 6, 1, 0, 0, 0, 263, 264, 5, 105, 0, 0, 264,
		265, 5, 109, 0, 0, 265, 266, 5, 112, 0, 0, 266, 267, 5, 111, 0, 0, 267,
		268, 5, 114, 0, 0, 268, 269, 5, 116, 0, 0, 269, 8, 1, 0, 0, 0, 270, 271,
		5, 102, 0, 0, 271, 272, 5, 114, 0, 0, 272, 273, 5, 111, 0, 0, 273, 274,
		5, 109, 0, 0, 274, 10, 1, 0, 0, 0, 275, 276, 5, 116, 0, 0, 276, 277, 5,
		121, 0, 0, 277, 278, 5, 112, 0, 0, 278, 279, 5, 101, 0, 0, 279, 12, 1,
		0, 0, 0, 280, 281, 5, 116, 0, 0, 281, 282, 5, 121, 0, 0, 282, 283, 5, 112,
		0, 0, 283, 284, 5, 101, 0, 0, 284, 285, 5, 111, 0, 0, 285, 286, 5, 102,
		0, 0, 286, 14, 1, 0, 0, 0, 287, 288, 5, 97, 0, 0, 288, 289, 5, 115, 0,
		0, 289, 16, 1, 0, 0, 0, 290, 291, 5, 105, 0, 0, 291, 292, 5, 102, 0, 0,
		292, 18, 1, 0, 0, 0, 293, 294, 5, 101, 0, 0, 294, 295, 5, 108, 0, 0, 295,
		296, 5, 115, 0, 0, 296, 297, 5, 101, 0, 0, 297, 20, 1, 0, 0, 0, 298, 299,
		5, 109, 0, 0, 299, 300, 5, 97, 0, 0, 300, 301, 5, 116, 0, 0, 301, 302,
		5, 99, 0, 0, 302, 303, 5, 104, 0, 0, 303, 22, 1, 0, 0, 0, 304, 305, 5,
		100, 0, 0, 305, 306, 5, 111, 0, 0, 306, 24, 1, 0, 0, 0, 307, 308, 5, 108,
		0, 0, 308, 309, 5, 111, 0, 0, 309, 310, 5, 111, 0, 0, 310, 311, 5, 112,
		0, 0, 311, 26, 1, 0, 0, 0, 312, 313, 5, 105, 0, 0, 313, 314, 5, 110, 0,
		0, 314, 28, 1, 0, 0, 0, 315, 316, 5, 114, 0, 0, 316, 317, 5, 97, 0, 0,
		317, 318, 5, 110, 0, 0, 318, 319, 5, 103, 0, 0, 319, 320, 5, 101, 0, 0,
		320, 30, 1, 0, 0, 0, 321, 322, 5, 99, 0, 0, 322, 323, 5, 111, 0, 0, 323,
		324, 5, 110, 0, 0, 324, 325, 5, 116, 0, 0, 325, 326, 5, 105, 0, 0, 326,
		327, 5, 110, 0, 0, 327, 328, 5, 117, 0, 0, 328, 329, 5, 101, 0, 0, 329,
		32, 1, 0, 0, 0, 330, 331, 5, 98, 0, 0, 331, 332, 5, 114, 0, 0, 332, 333,
		5, 101, 0, 0, 333, 334, 5, 97, 0, 0, 334, 335, 5, 107, 0, 0, 335, 34, 1,
		0, 0, 0, 336, 337, 5, 108, 0, 0, 337, 338, 5, 101, 0, 0, 338, 339, 5, 116,
		0, 0, 339, 36, 1, 0, 0, 0, 340, 341, 5, 118, 0, 0, 341, 342, 5, 97, 0,
		0, 342, 343, 5, 114, 0, 0, 343, 38, 1, 0, 0, 0, 344, 345, 5, 99, 0, 0,
		345, 346, 5, 111, 0, 0, 346, 347, 5, 110, 0, 0, 347, 348, 5, 115, 0, 0,
		348, 349, 5, 116, 0, 0, 349, 40, 1, 0, 0, 0, 350, 351, 5, 115, 0, 0, 351,
		352, 5, 116, 0, 0, 352, 353, 5, 97, 0, 0, 353, 354, 5, 116, 0, 0, 354,
		355, 5, 105, 0, 0, 355, 356, 5, 99, 0, 0, 356, 42, 1, 0, 0, 0, 357, 358,
		5, 102, 0, 0, 358, 359, 5, 105, 0, 0, 359, 360, 5, 110, 0, 0, 360, 361,
		5, 97, 0, 0, 361, 362, 5, 108, 0, 0, 362, 44, 1, 0, 0, 0, 363, 364, 5,
		112, 0, 0, 364, 365, 5, 117, 0, 0, 365, 366, 5, 98, 0, 0, 366, 46, 1, 0,
		0, 0, 367, 368, 5, 114, 0, 0, 368, 369, 5, 101, 0, 0, 369, 370, 5, 113,
		0, 0, 370, 371, 5, 117, 0, 0, 371, 372, 5, 105, 0, 0, 372, 373, 5, 114,
		0, 0, 373, 374, 5, 101, 0, 0, 374, 375, 5, 100, 0, 0, 375, 48, 1, 0, 0,
		0, 376, 377, 5, 102, 0, 0, 377, 378, 5, 97, 0, 0, 378, 379, 5, 99, 0, 0,
		379, 380, 5, 116, 0, 0, 380, 381, 5, 111, 0, 0, 381, 382, 5, 114, 0, 0,
		382, 383, 5, 121, 0, 0, 383, 50, 1, 0, 0, 0, 384, 385, 5, 116, 0, 0, 385,
		386, 5, 114, 0, 0, 386, 387, 5, 121, 0, 0, 387, 52, 1, 0, 0, 0, 388, 389,
		5, 99, 0, 0, 389, 390, 5, 97, 0, 0, 390, 391, 5, 116, 0, 0, 391, 392, 5,
		99, 0, 0, 392, 393, 5, 104, 0, 0, 393, 54, 1, 0, 0, 0, 394, 395, 5, 102,
		0, 0, 395, 396, 5, 105, 0, 0, 396, 397, 5, 110, 0, 0, 397, 398, 5, 97,
		0, 0, 398, 399, 5, 108, 0, 0, 399, 400, 5, 108, 0, 0, 400, 401, 5, 121,
		0, 0, 401, 56, 1, 0, 0, 0, 402, 403, 5, 116, 0, 0, 403, 404, 5, 104, 0,
		0, 404, 405, 5, 114, 0, 0, 405, 406, 5, 111, 0, 0, 406, 407, 5, 119, 0,
		0, 407, 58, 1, 0, 0, 0, 408, 409, 5, 116, 0, 0, 409, 410, 5, 104, 0, 0,
		410, 411, 5, 114, 0, 0, 411, 412, 5, 111, 0, 0, 412, 413, 5, 119, 0, 0,
		413, 414, 5, 115, 0, 0, 414, 60, 1, 0, 0, 0, 415, 416, 5, 102, 0, 0, 416,
		417, 5, 110, 0, 0, 417, 62, 1, 0, 0, 0, 418, 419, 5, 111, 0, 0, 419, 420,
		5, 112, 0, 0, 420, 421, 5, 101, 0, 0, 421, 422, 5, 114, 0, 0, 422, 423,
		5, 97, 0, 0, 423, 424, 5, 116, 0, 0, 424, 425, 5, 111, 0, 0, 425, 426,
		5, 114, 0, 0, 426, 64, 1, 0, 0, 0, 427, 428, 5, 114, 0, 0, 428, 429, 5,
		101, 0, 0, 429, 430, 5, 116, 0, 0, 430, 431, 5, 117, 0, 0, 431, 432, 5,
		114, 0, 0, 432, 433, 5, 110, 0, 0, 433, 66, 1, 0, 0, 0, 434, 435, 5, 101,
		0, 0, 435, 436, 5, 110, 0, 0, 436, 437, 5, 117, 0, 0, 437, 438, 5, 109,
		0, 0, 438, 68, 1, 0, 0, 0, 439, 440, 5, 99, 0, 0, 440, 441, 5, 108, 0,
		0, 441, 442, 5, 97, 0, 0, 442, 443, 5, 115, 0, 0, 443, 444, 5, 115, 0,
		0, 444, 70, 1, 0, 0, 0, 445, 446, 5, 116, 0, 0, 446, 447, 5, 114, 0, 0,
		447, 448, 5, 97, 0, 0, 448, 449, 5, 105, 0, 0, 449, 450, 5, 116, 0, 0,
		450, 72, 1, 0, 0, 0, 451, 452, 5, 105, 0, 0, 452, 453, 5, 109, 0, 0, 453,
		454, 5, 112, 0, 0, 454, 455, 5, 108, 0, 0, 455, 74, 1, 0, 0, 0, 456, 457,
		5, 102, 0, 0, 457, 458, 5, 111, 0, 0, 458, 459, 5, 114, 0, 0, 459, 76,
		1, 0, 0, 0, 460, 461, 5, 116, 0, 0, 461, 462, 5, 104, 0, 0, 462, 463, 5,
		105, 0, 0, 463, 464, 5, 115, 0, 0, 464, 78, 1, 0, 0, 0, 465, 466, 5, 115,
		0, 0, 466, 467, 5, 117, 0, 0, 467, 468, 5, 112, 0, 0, 468, 469, 5, 101,
		0, 0, 469, 470, 5, 114, 0, 0, 470, 80, 1, 0, 0, 0, 471, 472, 5, 101, 0,
		0, 472, 473, 5, 120, 0, 0, 473, 474, 5, 116, 0, 0, 474, 475, 5, 101, 0,
		0, 475, 476, 5, 110, 0, 0, 476, 477, 5, 100, 0, 0, 477, 82, 1, 0, 0, 0,
		478, 479, 5, 100, 0, 0, 479, 480, 5, 101, 0, 0, 480, 481, 5, 99, 0, 0,
		481, 482, 5, 108, 0, 0, 482, 483, 5, 97, 0, 0, 483, 484, 5, 114, 0, 0,
		484, 485, 5, 101, 0, 0, 485, 84, 1, 0, 0, 0, 486, 487, 5, 100, 0, 0, 487,
		488, 5, 101, 0, 0, 488, 489, 5, 102, 0, 0, 489, 490, 5, 101, 0, 0, 490,
		491, 5, 114, 0, 0, 491, 86, 1, 0, 0, 0, 492, 493, 5, 99, 0, 0, 493, 494,
		5, 111, 0, 0, 494, 495, 5, 109, 0, 0, 495, 496, 5, 112, 0, 0, 496, 497,
		5, 116, 0, 0, 497, 498, 5, 105, 0, 0, 498, 499, 5, 109, 0, 0, 499, 500,
		5, 101, 0, 0, 500, 88, 1, 0, 0, 0, 501, 502, 5, 117, 0, 0, 502, 503, 5,
		110, 0, 0, 503, 504, 5, 115, 0, 0, 504, 505, 5, 97, 0, 0, 505, 506, 5,
		102, 0, 0, 506, 507, 5, 101, 0, 0, 507, 90, 1, 0, 0, 0, 508, 509, 5, 46,
		0, 0, 509, 92, 1, 0, 0, 0, 510, 511, 5, 46, 0, 0, 511, 512, 5, 46, 0, 0,
		512, 94, 1, 0, 0, 0, 513, 514, 5, 44, 0, 0, 514, 96, 1, 0, 0, 0, 515, 516,
		5, 40, 0, 0, 516, 98, 1, 0, 0, 0, 517, 518, 5, 41, 0, 0, 518, 100, 1, 0,
		0, 0, 519, 520, 5, 91, 0, 0, 520, 102, 1, 0, 0, 0, 521, 522, 5, 93, 0,
		0, 522, 104, 1, 0, 0, 0, 523, 524, 5, 123, 0, 0, 524, 106, 1, 0, 0, 0,
		525, 526, 5, 125, 0, 0, 526, 108, 1, 0, 0, 0, 527, 528, 5, 42, 0, 0, 528,
		529, 5, 42, 0, 0, 529, 110, 1, 0, 0, 0, 530, 531, 5, 42, 0, 0, 531, 112,
		1, 0, 0, 0, 532, 533, 5, 37, 0, 0, 533, 114, 1, 0, 0, 0, 534, 535, 5, 47,
		0, 0, 535, 116, 1, 0, 0, 0, 536, 537, 5, 43, 0, 0, 537, 118, 1, 0, 0, 0,
		538, 539, 5, 45, 0, 0, 539, 120, 1, 0, 0, 0, 540, 541, 5, 35, 0, 0, 541,
		122, 1, 0, 0, 0, 542, 543, 5, 64, 0, 0, 543, 124, 1, 0, 0, 0, 544, 545,
		5, 63, 0, 0, 545, 126, 1, 0, 0, 0, 546, 547, 5, 36, 0, 0, 547, 128, 1,
		0, 0, 0, 548, 549, 5, 92, 0, 0, 549, 130, 1, 0, 0, 0, 550, 551, 5, 95,
		0, 0, 551, 132, 1, 0, 0, 0, 552, 553, 5, 61, 0, 0, 553, 134, 1, 0, 0, 0,
		554, 555, 5, 43, 0, 0, 555, 556, 5, 61, 0, 0, 556, 136, 1, 0, 0, 0, 557,
		558, 5, 45, 0, 0, 558, 559, 5, 61, 0, 0, 559, 138, 1, 0, 0, 0, 560, 561,
		5, 42, 0, 0, 561, 562, 5, 61, 0, 0, 562, 140, 1, 0, 0, 0, 563, 564, 5,
		42, 0, 0, 564, 565, 5, 42, 0, 0, 565, 566, 5, 61, 0, 0, 566, 142, 1, 0,
		0, 0, 567, 568, 5, 47, 0, 0, 568, 569, 5, 61, 0, 0, 569, 144, 1, 0, 0,
		0, 570, 571, 5, 37, 0, 0, 571, 572, 5, 61, 0, 0, 572, 146, 1, 0, 0, 0,
		573, 574, 5, 38, 0, 0, 574, 575, 5, 38, 0, 0, 575, 576, 5, 61, 0, 0, 576,
		148, 1, 0, 0, 0, 577, 578, 5, 60, 0, 0, 578, 150, 1, 0, 0, 0, 579, 580,
		5, 62, 0, 0, 580, 152, 1, 0, 0, 0, 581, 582, 5, 60, 0, 0, 582, 583, 5,
		61, 0, 0, 583, 154, 1, 0, 0, 0, 584, 585, 5, 62, 0, 0, 585, 586, 5, 61,
		0, 0, 586, 156, 1, 0, 0, 0, 587, 588, 5, 33, 0, 0, 588, 589, 5, 61, 0,
		0, 589, 158, 1, 0, 0, 0, 590, 591, 5, 61, 0, 0, 591, 592, 5, 61, 0, 0,
		592, 160, 1, 0, 0, 0, 593, 594, 5, 60, 0, 0, 594, 595, 5, 60, 0, 0, 595,
		162, 1, 0, 0, 0, 596, 597, 5, 62, 0, 0, 597, 598, 5, 62, 0, 0, 598, 164,
		1, 0, 0, 0, 599, 600, 5, 45, 0, 0, 600, 601, 5, 62, 0, 0, 601, 166, 1,
		0, 0, 0, 602, 603, 5, 60, 0, 0, 603, 604, 5, 45, 0, 0, 604, 168, 1, 0,
		0, 0, 605, 606, 5, 61, 0, 0, 606, 607, 5, 62, 0, 0, 607, 170, 1, 0, 0,
		0, 608, 609, 5, 46, 0, 0, 609, 610, 5, 46, 0, 0, 610, 611, 5, 46, 0, 0,
		611, 172, 1, 0, 0, 0, 612, 613, 5, 58, 0, 0, 613, 174, 1, 0, 0, 0, 614,
		615, 5, 58, 0, 0, 615, 616, 5, 58, 0, 0, 616, 176, 1, 0, 0, 0, 617, 618,
		5, 59, 0, 0, 618, 178, 1, 0, 0, 0, 619, 620, 5, 38, 0, 0, 620, 621, 5,
		38, 0, 0, 621, 180, 1, 0, 0, 0, 622, 623, 5, 124, 0, 0, 623, 624, 5, 124,
		0, 0, 624, 182, 1, 0, 0, 0, 625, 626, 5, 33, 0, 0, 626, 184, 1, 0, 0, 0,
		627, 628, 5, 43, 0, 0, 628, 629, 5, 43, 0, 0, 629, 186, 1, 0, 0, 0, 630,
		631, 5, 45, 0, 0, 631, 632, 5, 45, 0, 0, 632, 188, 1, 0, 0, 0, 633, 634,
		5, 38, 0, 0, 634, 190, 1, 0, 0, 0, 635, 636, 5, 124, 0, 0, 636, 192, 1,
		0, 0, 0, 637, 638, 5, 94, 0, 0, 638, 194, 1, 0, 0, 0, 639, 640, 5, 39,
		0, 0, 640, 196, 1, 0, 0, 0, 641, 642, 5, 34, 0, 0, 642, 198, 1, 0, 0, 0,
		643, 644, 5, 34, 0, 0, 644, 645, 5, 34, 0, 0, 645, 646, 5, 34, 0, 0, 646,
		200, 1, 0, 0, 0, 647, 648, 5, 110, 0, 0, 648, 649, 5, 101, 0, 0, 649, 650,
		5, 119, 0, 0, 650, 202, 1, 0, 0, 0, 651, 652, 5, 100, 0, 0, 652, 653, 5,
		101, 0, 0, 653, 654, 5, 108, 0, 0, 654, 655, 5, 101, 0, 0, 655, 656, 5,
		116, 0, 0, 656, 657, 5, 101, 0, 0, 657, 204, 1, 0, 0, 0, 658, 659, 5, 110,
		0, 0, 659, 660, 5, 117, 0, 0, 660, 661, 5, 108, 0, 0, 661, 662, 5, 108,
		0, 0, 662, 206, 1, 0, 0, 0, 663, 664, 5, 110, 0, 0, 664, 665, 5, 117, 0,
		0, 665, 666, 5, 109, 0, 0, 666, 208, 1, 0, 0, 0, 667, 668, 5, 115, 0, 0,
		668, 669, 5, 116, 0, 0, 669, 670, 5, 114, 0, 0, 670, 210, 1, 0, 0, 0, 671,
		672, 5, 98, 0, 0, 672, 673, 5, 111, 0, 0, 673, 674, 5, 111, 0, 0, 674,
		675, 5, 108, 0, 0, 675, 212, 1, 0, 0, 0, 676, 677, 5, 97, 0, 0, 677, 678,
		5, 110, 0, 0, 678, 679, 5, 121, 0, 0, 679, 214, 1, 0, 0, 0, 680, 687, 3,
		217, 108, 0, 681, 682, 3, 129, 64, 0, 682, 683, 3, 103, 51, 0, 683, 686,
		1, 0, 0, 0, 684, 686, 8, 0, 0, 0, 685, 681, 1, 0, 0, 0, 685, 684, 1, 0,
		0, 0, 686, 689, 1, 0, 0, 0, 687, 685, 1, 0, 0, 0, 687, 688, 1, 0, 0, 0,
		688, 690, 1, 0, 0, 0, 689, 687, 1, 0, 0, 0, 690, 691, 3, 219, 109, 0, 691,
		216, 1, 0, 0, 0, 692, 693, 5, 91, 0, 0, 693, 694, 5, 91, 0, 0, 694, 218,
		1, 0, 0, 0, 695, 696, 5, 93, 0, 0, 696, 697, 5, 93, 0, 0, 697, 220, 1,
		0, 0, 0, 698, 699, 5, 116, 0, 0, 699, 700, 5, 114, 0, 0, 700, 701, 5, 117,
		0, 0, 701, 702, 5, 101, 0, 0, 702, 222, 1, 0, 0, 0, 703, 704, 5, 102, 0,
		0, 704, 705, 5, 97, 0, 0, 705, 706, 5, 108, 0, 0, 706, 707, 5, 115, 0,
		0, 707, 708, 5, 101, 0, 0, 708, 224, 1, 0, 0, 0, 709, 711, 5, 45, 0, 0,
		710, 709, 1, 0, 0, 0, 710, 711, 1, 0, 0, 0, 711, 713, 1, 0, 0, 0, 712,
		714, 7, 1, 0, 0, 713, 712, 1, 0, 0, 0, 714, 715, 1, 0, 0, 0, 715, 713,
		1, 0, 0, 0, 715, 716, 1, 0, 0, 0, 716, 723, 1, 0, 0, 0, 717, 719, 5, 46,
		0, 0, 718, 720, 7, 1, 0, 0, 719, 718, 1, 0, 0, 0, 720, 721, 1, 0, 0, 0,
		721, 719, 1, 0, 0, 0, 721, 722, 1, 0, 0, 0, 722, 724, 1, 0, 0, 0, 723,
		717, 1, 0, 0, 0, 723, 724, 1, 0, 0, 0, 724, 226, 1, 0, 0, 0, 725, 730,
		3, 195, 97, 0, 726, 729, 3, 237, 118, 0, 727, 729, 8, 2, 0, 0, 728, 726,
		1, 0, 0, 0, 728, 727, 1, 0, 0, 0, 729, 732, 1, 0, 0, 0, 730, 728, 1, 0,
		0, 0, 730, 731, 1, 0, 0, 0, 731, 733, 1, 0, 0, 0, 732, 730, 1, 0, 0, 0,
		733, 734, 3, 195, 97, 0, 734, 756, 1, 0, 0, 0, 735, 740, 3, 197, 98, 0,
		736, 739, 3, 237, 118, 0, 737, 739, 8, 3, 0, 0, 738, 736, 1, 0, 0, 0, 738,
		737, 1, 0, 0, 0, 739, 742, 1, 0, 0, 0, 740, 738, 1, 0, 0, 0, 740, 741,
		1, 0, 0, 0, 741, 743, 1, 0, 0, 0, 742, 740, 1, 0, 0, 0, 743, 744, 3, 197,
		98, 0, 744, 756, 1, 0, 0, 0, 745, 749, 3, 199, 99, 0, 746, 748, 9, 0, 0,
		0, 747, 746, 1, 0, 0, 0, 748, 751, 1, 0, 0, 0, 749, 750, 1, 0, 0, 0, 749,
		747, 1, 0, 0, 0, 750, 752, 1, 0, 0, 0, 751, 749, 1, 0, 0, 0, 752, 753,
		3, 199, 99, 0, 753, 756, 1, 0, 0, 0, 754, 756, 3, 231, 115, 0, 755, 725,
		1, 0, 0, 0, 755, 735, 1, 0, 0, 0, 755, 745, 1, 0, 0, 0, 755, 754, 1, 0,
		0, 0, 756, 228, 1, 0, 0, 0, 757, 760, 3, 221, 110, 0, 758, 760, 3, 223,
		111, 0, 759, 757, 1, 0, 0, 0, 759, 758, 1, 0, 0, 0, 760, 230, 1, 0, 0,
		0, 761, 762, 5, 114, 0, 0, 762, 763, 5, 34, 0, 0, 763, 768, 1, 0, 0, 0,
		764, 767, 3, 237, 118, 0, 765, 767, 8, 3, 0, 0, 766, 764, 1, 0, 0, 0, 766,
		765, 1, 0, 0, 0, 767, 770, 1, 0, 0, 0, 768, 766, 1, 0, 0, 0, 768, 769,
		1, 0, 0, 0, 769, 771, 1, 0, 0, 0, 770, 768, 1, 0, 0, 0, 771, 772, 3, 197,
		98, 0, 772, 232, 1, 0, 0, 0, 773, 774, 5, 102, 0, 0, 774, 775, 5, 34, 0,
		0, 775, 780, 1, 0, 0, 0, 776, 779, 3, 237, 118, 0, 777, 779, 8, 3, 0, 0,
		778, 776, 1, 0, 0, 0, 778, 777, 1, 0, 0, 0, 779, 782, 1, 0, 0, 0, 780,
		778, 1, 0, 0, 0, 780, 781, 1, 0, 0, 0, 781, 783, 1, 0, 0, 0, 782, 780,
		1, 0, 0, 0, 783, 784, 3, 197, 98, 0, 784, 234, 1, 0, 0, 0, 785, 786, 5,
		99, 0, 0, 786, 787, 5, 34, 0, 0, 787, 792, 1, 0, 0, 0, 788, 791, 3, 237,
		118, 0, 789, 791, 8, 3, 0, 0, 790, 788, 1, 0, 0, 0, 790, 789, 1, 0, 0,
		0, 791, 794, 1, 0, 0, 0, 792, 790, 1, 0, 0, 0, 792, 793, 1, 0, 0, 0, 793,
		795, 1, 0, 0, 0, 794, 792, 1, 0, 0, 0, 795, 796, 3, 197, 98, 0, 796, 236,
		1, 0, 0, 0, 797, 798, 5, 92, 0, 0, 798, 799, 7, 4, 0, 0, 799, 238, 1, 0,
		0, 0, 800, 801, 5, 47, 0, 0, 801, 802, 5, 47, 0, 0, 802, 806, 1, 0, 0,
		0, 803, 805, 8, 5, 0, 0, 804, 803, 1, 0, 0, 0, 805, 808, 1, 0, 0, 0, 806,
		804, 1, 0, 0, 0, 806, 807, 1, 0, 0, 0, 807, 240, 1, 0, 0, 0, 808, 806,
		1, 0, 0, 0, 809, 810, 5, 47, 0, 0, 810, 811, 5, 42, 0, 0, 811, 815, 1,
		0, 0, 0, 812, 814, 9, 0, 0, 0, 813, 812, 1, 0, 0, 0, 814, 817, 1, 0, 0,
		0, 815, 816, 1, 0, 0, 0, 815, 813, 1, 0, 0, 0, 816, 818, 1, 0, 0, 0, 817,
		815, 1, 0, 0, 0, 818, 819, 5, 42, 0, 0, 819, 820, 5, 47, 0, 0, 820, 242,
		1, 0, 0, 0, 821, 825, 7, 6, 0, 0, 822, 824, 7, 7, 0, 0, 823, 822, 1, 0,
		0, 0, 824, 827, 1, 0, 0, 0, 825, 823, 1, 0, 0, 0, 825, 826, 1, 0, 0, 0,
		826, 244, 1, 0, 0, 0, 827, 825, 1, 0, 0, 0, 828, 830, 7, 8, 0, 0, 829,
		828, 1, 0, 0, 0, 830, 831, 1, 0, 0, 0, 831, 829, 1, 0, 0, 0, 831, 832,
		1, 0, 0, 0, 832, 833, 1, 0, 0, 0, 833, 834, 6, 122, 0, 0, 834, 246, 1,
		0, 0, 0, 24, 0, 685, 687, 710, 715, 721, 723, 728, 730, 738, 740, 749,
		755, 759, 766, 768, 778, 780, 790, 792, 806, 815, 825, 831, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// huloLexerInit initializes any static state used to implement huloLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewhuloLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func HuloLexerInit() {
	staticData := &HuloLexerLexerStaticData
	staticData.once.Do(hulolexerLexerInit)
}

// NewhuloLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewhuloLexer(input antlr.CharStream) *huloLexer {
	HuloLexerInit()
	l := new(huloLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &HuloLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "huloLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// huloLexer tokens.
const (
	huloLexerMOD_LIT              = 1
	huloLexerUSE                  = 2
	huloLexerPACKAGE              = 3
	huloLexerIMPORT               = 4
	huloLexerFROM                 = 5
	huloLexerTYPE                 = 6
	huloLexerTYPEOF               = 7
	huloLexerAS                   = 8
	huloLexerIF                   = 9
	huloLexerELSE                 = 10
	huloLexerMATCH                = 11
	huloLexerDO                   = 12
	huloLexerLOOP                 = 13
	huloLexerIN                   = 14
	huloLexerRANGE                = 15
	huloLexerCONTINUE             = 16
	huloLexerBREAK                = 17
	huloLexerLET                  = 18
	huloLexerVAR                  = 19
	huloLexerCONST                = 20
	huloLexerSTATIC               = 21
	huloLexerFINAL                = 22
	huloLexerPUB                  = 23
	huloLexerREQUIRED             = 24
	huloLexerFACTORY              = 25
	huloLexerTRY                  = 26
	huloLexerCATCH                = 27
	huloLexerFINALLY              = 28
	huloLexerTHROW                = 29
	huloLexerTHROWS               = 30
	huloLexerFN                   = 31
	huloLexerOPERATOR             = 32
	huloLexerRETURN               = 33
	huloLexerENUM                 = 34
	huloLexerCLASS                = 35
	huloLexerTRAIT                = 36
	huloLexerIMPL                 = 37
	huloLexerFOR                  = 38
	huloLexerTHIS                 = 39
	huloLexerSUPER                = 40
	huloLexerEXTEND               = 41
	huloLexerDECLARE              = 42
	huloLexerDEFER                = 43
	huloLexerCOMPTIME             = 44
	huloLexerUNSAFE               = 45
	huloLexerDOT                  = 46
	huloLexerDOUBLE_DOT           = 47
	huloLexerCOMMA                = 48
	huloLexerLPAREN               = 49
	huloLexerRPAREN               = 50
	huloLexerLBRACK               = 51
	huloLexerRBRACK               = 52
	huloLexerLBRACE               = 53
	huloLexerRBRACE               = 54
	huloLexerEXP                  = 55
	huloLexerMUL                  = 56
	huloLexerMOD                  = 57
	huloLexerDIV                  = 58
	huloLexerADD                  = 59
	huloLexerSUB                  = 60
	huloLexerHASH                 = 61
	huloLexerAT                   = 62
	huloLexerQUEST                = 63
	huloLexerDOLLAR               = 64
	huloLexerBACKSLASH            = 65
	huloLexerWILDCARD             = 66
	huloLexerASSIGN               = 67
	huloLexerADD_ASSIGN           = 68
	huloLexerSUB_ASSIGN           = 69
	huloLexerMUL_ASSIGN           = 70
	huloLexerEXP_ASSIGN           = 71
	huloLexerDIV_ASSIGN           = 72
	huloLexerMOD_ASSIGN           = 73
	huloLexerAND_ASSIGN           = 74
	huloLexerLT                   = 75
	huloLexerGT                   = 76
	huloLexerLE                   = 77
	huloLexerGE                   = 78
	huloLexerNEQ                  = 79
	huloLexerEQ                   = 80
	huloLexerSHL                  = 81
	huloLexerSHR                  = 82
	huloLexerARROW                = 83
	huloLexerBACKARROW            = 84
	huloLexerDOUBLE_ARROW         = 85
	huloLexerELLIPSIS             = 86
	huloLexerCOLON                = 87
	huloLexerDOUBLE_COLON         = 88
	huloLexerSEMI                 = 89
	huloLexerAND                  = 90
	huloLexerOR                   = 91
	huloLexerNOT                  = 92
	huloLexerINC                  = 93
	huloLexerDEC                  = 94
	huloLexerBITAND               = 95
	huloLexerBITOR                = 96
	huloLexerBITXOR               = 97
	huloLexerSINGLE_QUOTE         = 98
	huloLexerQUOTE                = 99
	huloLexerTRIPLE_QUOTE         = 100
	huloLexerNEW                  = 101
	huloLexerDELETE               = 102
	huloLexerNULL                 = 103
	huloLexerNUM                  = 104
	huloLexerSTR                  = 105
	huloLexerBOOL                 = 106
	huloLexerANY                  = 107
	huloLexerUnsafeLiteral        = 108
	huloLexerDOUBLE_LBRACK        = 109
	huloLexerDOUBLE_RBRACK        = 110
	huloLexerNumberLiteral        = 111
	huloLexerStringLiteral        = 112
	huloLexerBoolLiteral          = 113
	huloLexerRawStringLiteral     = 114
	huloLexerFileStringLiteral    = 115
	huloLexerCommandStringLiteral = 116
	huloLexerESC                  = 117
	huloLexerLineComment          = 118
	huloLexerBlockComment         = 119
	huloLexerIdentifier           = 120
	huloLexerWS                   = 121
)
