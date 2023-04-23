package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// This is a tests
type Kurtis interface {
	// +kubebuilder:validation:Maximum=100
	FindByID()
}

func main() {
	path := "/home/kurtis/wrkspc/github.com/angelokurtis/go-rhonline"
	directories, err := listDirectories(path)
	dieIfErr(err)

	fset := token.NewFileSet()
	filter := func(info fs.FileInfo) bool { return true }

	for _, directory := range directories {
		pkgs, err := parser.ParseDir(fset, directory, filter, parser.ParseComments)
		dieIfErr(errors.WithStack(err))
		_ = pkgs
		//
		// 	for _, pkg := range pkgs {
		// 		slog.Info("", slog.String("directory", directory), slog.String("pkg", pkg.Name))
		//
		// 		for _, file := range pkg.Files {
		// 			ast.Inspect(file, InspectNode)
		// 		}
		// 	}
	}
}

func InspectNode(n ast.Node) bool {
	if file, ok := n.(*ast.File); ok {
		for _, i := range file.Imports {
			path := i.Path.Value
			fmt.Println(path)
		}
	}
	if decl, ok := n.(*ast.GenDecl); ok {
		InspectGenDecl(decl)
	}
	if decl, ok := n.(*ast.FuncDecl); ok {
		funcName := decl.Name.String()
		if f, _ := first(split(funcName)); f == "new" {
			if decl.Type.Results != nil {
				for _, result := range decl.Type.Results.List {
					if sexpr, isStarExpr := result.Type.(*ast.StarExpr); isStarExpr {
						_ = sexpr
					}
					_ = result
					// returnTypes = append(returnTypes, result.Type.(*ast.Ident).Name)
				}
			}
		}
	}
	return true
}

func InspectTypeSpec(ts *ast.TypeSpec) bool {
	switch expr := ts.Type.(type) {
	case *ast.InterfaceType:
		return InspectInterfaceType(ts.Name.String(), expr)
	case *ast.StructType:
		return InspectStructType(ts.Name.String(), expr)
	default:
		return false
	}
}

func InspectInterfaceType(name string, expr *ast.InterfaceType) bool {
	return false
}

func InspectStructType(name string, expr *ast.StructType) bool {
	return false
}

func InspectGenDecl(decl *ast.GenDecl) bool {
	if decl.Tok != token.TYPE {
		return false
	}

	for _, spec := range decl.Specs {
		if ts, isTypeSpec := spec.(*ast.TypeSpec); isTypeSpec {
			InspectTypeSpec(ts)
		}
	}

	return false
}

func listDirectories(path string) ([]string, error) {
	var directories []string

	fn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			directories = append(directories, path)
		}

		return nil
	}
	if err := filepath.Walk(path, fn); err != nil {
		return nil, errors.WithStack(err)
	}

	return directories, nil
}

func split(s string) []string {
	r := regexp.MustCompile(`[A-Z][^A-Z]*`)
	words := r.FindAllString(s, -1)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}

func first(s []string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("slice is empty")
	}
	return s[0], nil
}
