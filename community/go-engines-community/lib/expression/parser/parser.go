// parser contains expression parser for search parameter.
package parser

import (
	"fmt"
	"strings"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"go.mongodb.org/mongo-driver/bson"
)

// Parser parses expression.
type Parser interface {
	Parse(string) (MongoQuery, error)
}

type MongoQuery interface {
	Query() bson.M
	GetFields() []string
}

// NewParser creates new parser.
func NewParser() Parser {
	l := lexer.Must(lexer.Regexp(`(\s+)` +
		`|(?P<Keyword>(?i)TRUE|FALSE|NULL|NOT|AND|OR|LIKE|CONTAINS)` +
		`|(?P<Ident>[a-zA-Z_][a-zA-Z0-9_\.]*)` +
		`|(?P<Float>[-+]?\d*\.\d+([eE][-+]?\d+)?)` +
		`|(?P<Int>[-+]?\d+)` +
		`|(?P<String>"[^"]*")` +
		`|(?P<Operators>!=|<=|>=|<|>|=)`,
	))

	return &parser{
		baseParser: participle.MustBuild(
			&Expression{},
			participle.Lexer(l),
			participle.Unquote("String"),
			participle.CaseInsensitive("Keyword"),
		),
	}
}

// parser implements following syntax:
// expr:
//    expr OR expr
//  | expr AND expr
//  | NOT expr
//  | boolean_expr
//
// boolean_expr:
//    boolean_expr comparison_operator predicate
//  | predicate
//
// comparison_operator: = | >= | > | <= | < | !=
//
// predicate:
//    simple_expr [NOT] CONTAINS simple_expr
//  | simple_expr [NOT] LIKE simple_expr
//  | simple_expr
//
// simple_expr:
//    variable
//  | float
//  | int
//  | "string"
//  | TRUE
//  | FALSE
//  | NULL
type parser struct {
	baseParser *participle.Parser
}

func (p *parser) hasComparison(str string) bool {
	if strings.ContainsAny(str, "=><") {
		return true
	}
	s := strings.ToUpper(str)
	return strings.Contains(s, "LIKE") || strings.Contains(s, "CONTAINS")
}

func (p *parser) Parse(str string) (MongoQuery, error) {
	if !p.hasComparison(str) {
		return nil, fmt.Errorf("comparison not found")
	}
	expr := &Expression{}
	err := p.baseParser.ParseString(str, expr)
	if err != nil {
		return nil, err
	}

	return expr, nil
}
