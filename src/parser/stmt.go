package parser

import (
	"github.com/divakaivan/lang-parser-go/src/ast"
	"github.com/divakaivan/lang-parser-go/src/lexer"
)

// 10 + 5;

func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists {
		return stmt_fn(p)
	}

	expression := parse_expr(p, defalt_bp)
	p.expect(lexer.SEMI_COLON)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}
