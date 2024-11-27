package test

import (
	"http1/token"
	"http1/lexer"
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.TpMethod, "GET"},
		{token.TpUri, "/shop"},
		{token.TpProtocol, "HTTP"},
		{token.TpVersion, "/1.1"},
	}

	httpTest := lexer.New("GET /shop HTTP/1.1 \n")

        httpTest.RunLexer()

	tks := httpTest.Tks

	for i, tk := range tks {
		if i >= len(tests) {
			t.Errorf("Extra token found: got %v", tk)
			continue
		}

		expected := tests[i]
		if tk.Typ != expected.expectedType || tk.Lit != expected.expectedLiteral {
			t.Errorf("Mismatch at token %d: got %v, expected %v", i, tk, expected)
		}
	}

	if len(tks) < len(tests) {
		t.Errorf("Missing tokens: got %v, expected %v", tks, tests)
	}
}
