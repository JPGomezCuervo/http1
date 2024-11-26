package token
import "fmt"

type TokenType int

type Token struct {
        Typ TokenType
        Lit string
}

const (
        TpIllegal TokenType = iota
        TpEOF
        TpBody
        TpComment
        TpHeader
        TpMethod 
        TpProtocol
        TpSP
        TpVersion
        TpUri
)

var tokenTypeNames = map[TokenType]string{
        TpIllegal:  "TpIllegal",
        TpEOF:      "TpEOF",
        TpBody:     "TpBody",
        TpComment:  "TpComment",
        TpHeader:   "TpHeader",
        TpMethod:   "TpMethod",
        TpProtocol: "TpProtocol",
        TpSP:       "TpSP",
        TpVersion:  "TpVersion",
        TpUri:      "TpUri",
}

func (t TokenType) String() string {
        if name, ok := tokenTypeNames[t]; ok {
                return name
        }
        return fmt.Sprintf("UnknownTokenType(%d)", t)
}

func (t Token) String() string {
	return fmt.Sprintf("&{%s %q}", t.Typ, t.Lit)
}
