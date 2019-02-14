package lib

import (
	"github.com/spf13/cast"
	"go/ast"
	"go/token"
)

type funcArg struct {
	node *RunNode
	dep  int
	val  string
}

type expr struct {
	args     []interface{}
	t        int
	method   string
	instance interface{}
}

func CompileExpr(x ast.Expr, r *RunNode) (ret interface{}) {
	noret := false
	defer func() {
		if ret != nil || noret {
			Debug("compile_expr -> %#v , ret -> %+v ", x, ret)
		} else {
			Error("compile_expr -> %#v , ret -> %+v ", x, ret)
		}
	}()
	switch x := x.(type) {
	case *ast.BasicLit:
		switch x.Kind {
		case token.INT:
			ret = cast.ToInt(x.Value)
		case token.FLOAT:
			//TODO
		case token.STRING:
			ret = x.Value
		default:
			Error("there is should happend_1 %#v \n", x)
		}
	case *ast.CompositeLit:
		//TODO
		Error("there is should happend_2 %#v \n", x)
	case *ast.FuncLit:
		//TODO
		Error("there is should happend_3 %#v \n", x)
	case *ast.ArrayType:
		//TODO
		Error("there is should happend_4 %#v \n", x)
	case *ast.ChanType:
		Error("there is should happend_4.1 %#v \n", x)
	case *ast.Ellipsis:
		Error("there is should happend_4.2 %#v \n", x)
	case *ast.FuncType:
		Error("there is should happend_4.3 %#v \n", x)
	case *ast.InterfaceType:
		Error("there is should happend_4.4 %#v \n", x)
	case *ast.MapType:
		return make(map[interface{}]interface{}, 0)
	case *ast.BadExpr:
		//TODO
		Error("there is should happend_5 %#v \n", x)
	case *ast.BinaryExpr:
		ret = CompileBinary(x, r)
	case *ast.CallExpr:
		noret = true
		switch f := x.Fun.(type) {
		case *ast.Ident:
			e := expr{
				method: f.Name,
				args:   CompileArgs(x.Args, r),
				//TODO t
			}
			return Invock(e, r)
		default:
			Error("there is should not happd")
		}

	case *ast.Ident:
		ret = r.VarMap[x.Name]
	case *ast.IndexExpr:
		l := CompileExpr(x.X, r)
		switch l := l.(type) {
		case map[interface{}]interface{}:
			return l[CompileExpr(x.Index, r)]
		default:
			Error("o it dont  ha %#v", l)
		}
	case *ast.SliceExpr:
		//TODO
		Error("there is should happend_8 %#v \n", x)
	case *ast.KeyValueExpr:
		//TODO
		Error("there is should happend_9 %#v \n", x)
	case *ast.ParenExpr:
		//TODO
		Error("there is should happend_10 %#v \n", x)
	case *ast.SelectorExpr:
		//TODO
		Error("there is should happend_11 %#v \n", x)
	case *ast.StarExpr:
		//TODO
		Error("there is should happend_12 %#v \n", x)
	case *ast.StructType:
		//TODO
		Error("there is should happend_13 %#v \n", x)
	case *ast.TypeAssertExpr:
		//TODO
		Error("there is should happend_14 %#v \n", x)
	case *ast.UnaryExpr:
		//TODO
		Error("there is should happend_15 %#v \n", x)
	default:
		//TODO
		Error("there is should happend_16 %#v \n", x)
	}
	return
}