package parse

import "github.com/awalterschulze/gographviz/ast"

type sortedStmtList []ast.Stmt

func (l sortedStmtList) Len() int {
	return len(l)
}

func (l sortedStmtList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l sortedStmtList) Less(i, j int) bool {
	var iIndex, jIndex int

	iIndex, _ = getOperationsIndex(l[i])
	jIndex, _ = getOperationsIndex(l[j])

	return iIndex < jIndex
}
