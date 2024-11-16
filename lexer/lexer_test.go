package lexer

import (
	"testing"

	"github.com/Nearrivers/combo-parser/token"
)

type Case struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func assertInput(t testing.TB, input string, cases []Case) {
	t.Helper()

	l := NewLexer(input)

	for i, tt := range cases {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("cases[%d] - tokentype of '%s' wrong, expected %q, got %q", i, tok.Literal, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("cases[%d] - literal wrong, expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	t.Run("combo basique", func(t *testing.T) {
		input := `2LP, 5LK~LK > 623HK`

		cases := []Case{
			{token.INPUT, "2LP"},
			{token.COMMA, ","},
			{token.INPUT, "5LK"},
			{token.TILDA, "~"},
			{token.INPUT, "LK"},
			{token.GT, ">"},
			{token.INPUT, "623HK"},
		}

		assertInput(t, input, cases)
	})

	t.Run("combo plus complexe", func(t *testing.T) {
		input := `5MP , 5LP~LK~MP > 236MP~MK~MK`

		cases := []Case{
			{token.INPUT, "5MP"},
			{token.COMMA, ","},
			{token.INPUT, "5LP"},
			{token.TILDA, "~"},
			{token.INPUT, "LK"},
			{token.TILDA, "~"},
			{token.INPUT, "MP"},
			{token.GT, ">"},
			{token.INPUT, "236MP"},
			{token.TILDA, "~"},
			{token.INPUT, "MK"},
			{token.TILDA, "~"},
			{token.INPUT, "MK"},
		}

		assertInput(t, input, cases)
	})

	t.Run("air combo", func(t *testing.T) {
		input := `2KK > dl.j.MP > j.214HK`

		cases := []Case{
			{token.INPUT, "2KK"},
			{token.GT, ">"},
			{token.INPUT, "dl.j.MP"},
			{token.GT, ">"},
			{token.INPUT, "j.214HK"},
		}

		assertInput(t, input, cases)
	})

	t.Run("punish counter drink level 2 hard combo", func(t *testing.T) {
		input := `DL2 [PC] DR~5HP, 2HP > 236KK > 214214P, 2HP > DR~5MP, 2HP > DR~5MP, 2HP > 214HP~6P`

		cases := []Case{
			{token.UNKOWN, "DL2"},
			{token.PUNISH_COUNTER, "[PC]"},
			{token.DRIVE_RUSH, "DR"},
			{token.TILDA, "~"},
			{token.INPUT, "5HP"},
			{token.COMMA, ","},
			{token.INPUT, "2HP"},
			{token.GT, ">"},
			{token.INPUT, "236KK"},
			{token.GT, ">"},
			{token.INPUT, "214214P"},
			{token.COMMA, ","},
			{token.INPUT, "2HP"},
			{token.GT, ">"},
			{token.DRIVE_RUSH, "DR"},
			{token.TILDA, "~"},
			{token.INPUT, "5MP"},
			{token.COMMA, ","},
			{token.INPUT, "2HP"},
			{token.GT, ">"},
			{token.DRIVE_RUSH, "DR"},
			{token.TILDA, "~"},
			{token.INPUT, "5MP"},
			{token.COMMA, ","},
			{token.INPUT, "2HP"},
			{token.GT, ">"},
			{token.INPUT, "214HP"},
			{token.TILDA, "~"},
			{token.INPUT, "6P"},
		}

		assertInput(t, input, cases)
	})

	t.Run("super extension with Drive impact", func(t *testing.T) {
		input := `... > 236236P, DR~2HP, 2LP > DI`

		cases := []Case{
			{token.EXTENSION, "..."},
			{token.GT, ">"},
			{token.INPUT, "236236P"},
			{token.COMMA, ","},
			{token.DRIVE_RUSH, "DR"},
			{token.TILDA, "~"},
			{token.INPUT, "2HP"},
			{token.COMMA, ","},
			{token.INPUT, "2LP"},
			{token.GT, ">"},
			{token.DRIVE_IMPACT, "DI"},
		}

		assertInput(t, input, cases)
	})

	t.Run("combo with a dash", func(t *testing.T) {
		input := `5HK, Dash, 2HP > 236236P, DR~5MP, 5MK, 63214K`

		cases := []Case{
			{token.INPUT, "5HK"},
			{token.COMMA, ","},
			{token.UNKOWN, "Dash"},
			{token.COMMA, ","},
			{token.INPUT, "2HP"},
			{token.GT, ">"},
			{token.INPUT, "236236P"},
			{token.COMMA, ","},
			{token.DRIVE_RUSH, "DR"},
			{token.TILDA, "~"},
			{token.INPUT, "5MP"},
			{token.COMMA, ","},
			{token.INPUT, "5MK"},
			{token.COMMA, ","},
			{token.INPUT, "63214K"},
		}

		assertInput(t, input, cases)
	})
}
