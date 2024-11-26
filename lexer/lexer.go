package lexer

import "http1/token"

type StateType int

const (
        stateRequestLine StateType = iota
        stateHeaders
        stateEOF
)

type Lexer struct {
        input string
        pos   int
        start int
        Ch    byte
}

func New(input string) *Lexer {
        l := &Lexer{input: input}
        l.readChar()
        return l
}

func (l *Lexer) readChar() {
        if l.pos >= len(l.input) {
                l.Ch = 0
        } 

        l.start = l.pos
        l.Ch = l.input[l.start]
        l.pos++
}

func (l *Lexer) backupChar() {
        l.pos--
        l.start--
        l.Ch = l.input[l.start]
}

func (l *Lexer) peek() byte {
        if l.pos >= len(l.input) {
                return 0
        } 

        return l.input[l.pos]
}

func (l *Lexer) readWord() string {
        start := l.start

        for !isSeparator(l.Ch) && l.Ch != 0 {
                l.readChar()
        }
        l.backupChar()

        return l.input[start:l.pos]
}

func isSeparator(c byte) bool {

        switch c {
        case ' ', '/', '\n':
                return true
        default:
                return false
        }
}

func (l *Lexer) mkToken(tp token.TokenType, lit string) *token.Token {
        return &token.Token{Typ: tp, Lit: lit}
}

func (l *Lexer) requestLine() ([]*token.Token, StateType) {
        var tks []*token.Token
        var prev token.TokenType

        for l.Ch != '\n' && l.Ch != 0 {
                if l.Ch == ' ' {
                        tks = append(tks, l.mkToken(token.TpSP, string(l.Ch)))
                        prev = token.TpSP
                } else if l.Ch == '/' {
                        if l.peek() >= '0' && l.peek() <= '9' && prev == token.TpProtocol {
                                l.readChar()
                                s := l.readWord()
                                tks = append(tks, l.mkToken(token.TpVersion, s))
                                prev = token.TpVersion

                        } else if prev == token.TpSP {
                                l.readChar()
                                s := l.readWord()
                                tks = append(tks, l.mkToken(token.TpUri, s))
                                prev = token.TpUri
                        }

                } else {
                        s := l.readWord()

                        switch s {
                        case "PUT", "GET", "DELETE", "POST":
                                tks = append(tks, l.mkToken(token.TpMethod, s))
                                prev = token.TpMethod
                        case "HTTP":
                                tks = append(tks, l.mkToken(token.TpProtocol, s))
                                prev = token.TpProtocol
                        default:
                                /* lanzar error */
                        }
                }
                l.readChar()
        }
        return tks, stateEOF
}

func (l *Lexer) RunLexer() []*token.Token {
        var tokens []*token.Token

        for state := stateRequestLine; state != stateEOF; {
                switch state {
                case stateRequestLine:
                        tk, nextState := l.requestLine()
                        tokens = append(tokens, tk...)
                        state = nextState

                }
        }
        return tokens
}
