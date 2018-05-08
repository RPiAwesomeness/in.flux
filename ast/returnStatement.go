package ast

import "github.com/RPiAwesomeness/in.flux/token"

type ReturnStatement struct {
	Token       token.Token // token.LET token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
