package lexer

import "http1/token"

type StateType int

const (
        stateRequestLine StateType = iota
        stateHeaders
        stateEOF
)

type Lexer struct {
        input   string
        pos     int
        start   int
        Ch      byte
        Tks     []*token.Token
}

func New(input string) *Lexer {
        l := &Lexer{input: input}
        l.readChar()
        return l
}

func (l *Lexer) readChar() byte {
        if l.pos >= len(l.input) {
                l.Ch = 0
        } 

        l.start = l.pos
        l.Ch = l.input[l.start]
        l.pos++

        return l.Ch
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

func (l *Lexer) readPath() string {
        start := l.start

        for l.Ch != ' ' && l.Ch != '\n' && l.Ch != 0 {
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

func isLetter(c byte) bool {
        if c >= 'a' && c <= 'z' { return true }
        if c >= 'A' && c <= 'Z' { return true }
        return false
}

func isEndLine (c byte) bool {
        if c == '\n' || c == '\r' { return true }
        return false
}

func (l *Lexer) appendToken(tp token.TokenType, lit string) *token.Token {
        t := &token.Token{Typ: tp, Lit: lit}
        l.Tks = append(l.Tks, t)
        return t
}

func (l *Lexer) mkToken(c byte) bool {
        if isSeparator(c) {
                switch c {
                case '/':
                        s := l.readPath()
                        prev := l.Tks[len(l.Tks)-1].Typ

                        if prev == token.TpProtocol { 
                                l.appendToken(token.TpVersion, s)
                        } else if prev == token.TpMethod {
                                l.appendToken(token.TpUri, s)
                        }
                        /* check the las if */
                case '\n':
                        return false
                }
        }

        if isLetter(c) {
                s := l.readWord()
                switch s {
                case "PUT", "GET", "DELETE", "POST":
                        l.appendToken(token.TpMethod, s)
                case "HTTP":
                        l.appendToken(token.TpProtocol, s)
                default:
                        return false
                }
        }
        return true
}

/* rethink this function */
func (l *Lexer) requestLine() (StateType) {

        for c := l.Ch; c != 0 && !isEndLine(c); c = l.readChar() {
                if c == ' ' { continue }
                l.mkToken(c)
        }

        return stateEOF
}

func (l *Lexer) RunLexer() []*token.Token {
        var tokens []*token.Token

        for state := stateRequestLine; state != stateEOF; {
                switch state {
                case stateRequestLine:
                        nextState := l.requestLine()
                        state = nextState
                }
        }
        return tokens
}
