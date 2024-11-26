package token

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
