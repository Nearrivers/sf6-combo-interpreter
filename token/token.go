package token

const (
	UNKOWN = "UNKNOWN"
	EOF    = "EOF"

	// Identifiants + input
	INPUT  = "INPUT"
	BUTTON = "BUTTON"

	// DÉLIMITEURS / Liens
	COMMA = ","
	TILDA = "~"
	GT    = ">"
	SLASH = "/"

	LPAREN = "("
	// RPAREN   = ")"
	RBRACKET = "["
	// LBRACKET = "]"
	MINUS = "-"

	// Mots-clés
	DASH           = "dash"
	DRIVE_RUSH     = "DR"
	DRIVE_IMPACT   = "DI"
	EXTENSION      = "..."
	PUNISH_COUNTER = "[PC]"
	COUNTER_HIT    = "[CH]"
	CRITICAL_ART   = "(CA)"
	TIGER_KNEE     = "tk"

	// Boutons
	// MEDIUM = "M"
	// LIGHT  = "L"
	// HEAVY  = "H"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"DR":   DRIVE_RUSH,
	"dash": DASH,
	"DI":   DRIVE_IMPACT,
	"...":  EXTENSION,
	"[CH]": COUNTER_HIT,
	"[PC]": PUNISH_COUNTER,
	"(CA)": CRITICAL_ART,
	"tk":   TIGER_KNEE,
}

var delimiter = map[byte]TokenType{
	',': COMMA,
	'>': GT,
	'~': TILDA,
	'/': SLASH,
	'-': MINUS,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return INPUT
}

func IsIdent(ident string) bool {
	_, ok := keywords[ident]
	return ok
}

func IsDelimiter(delim byte) bool {
	_, ok := delimiter[delim]
	return ok
}

// type Color string
//
// var colors = map[string]Color{
// 	LIGHT:          "#7dffff",
// 	MEDIUM:         "#ffff01",
// 	HEAVY:          "#ff9899",
// 	COUNTER_HIT:    "#b70c0b",
// 	PUNISH_COUNTER: "#b70c0b",
// }
