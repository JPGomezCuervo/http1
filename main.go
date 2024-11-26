package main

import( "http1/lexer"
        "fmt"
)

func main() {
        /* no maneja subrutas, mejorar*/
        httpLex := lexer.New("GET /shop HTTP/1.1 \n")

        tks := httpLex.RunLexer()

        for _, tk := range tks {
                fmt.Printf("%v\n", tk)
        }
}
