package parser

import "github.com/DeRuneLabs/jane/ast"

type defmap struct {
	Funcs   []*function
	Globals []ast.Var
	Types   []ast.Type
}

func (dm *defmap) findTypeById(id string) int {
	for i, t := range dm.Types {
		if t.Id == id {
			return i
		}
	}
	return -1
}

func (dm *defmap) typeById(id string) *ast.Type {
	i := dm.findTypeById(id)
	if i == -1 {
		return nil
	}
	return &dm.Types[i]
}

func (dm *defmap) findFuncById(id string) int {
	for i, f := range dm.Funcs {
		if f.Ast.Id == id {
			return i
		}
	}
	return -1
}

func (dm *defmap) funcById(id string) *function {
	i := dm.findFuncById(id)
	if i == -1 {
		return nil
	}
	return dm.Funcs[i]
}

func (dm *defmap) findGlobalById(id string) int {
	for i, v := range dm.Globals {
		if v.Id == id {
			return i
		}
	}
	return -1
}

func (dm *defmap) globalById(id string) *ast.Var {
	i := dm.findGlobalById(id)
	if i == -1 {
		return nil
	}
	return &dm.Globals[i]
}

func (dm *defmap) defById(id string) (int, byte) {
	var i int
	i = dm.findGlobalById(id)
	if i != -1 {
		return i, 'g'
	}
	i = dm.findFuncById(id)
	if i != -1 {
		return i, 'f'
	}
	return -1, ' '
}
