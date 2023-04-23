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
	"golang.org/x/exp/slog"
	"golang.org/x/mod/modfile"
)

func main() {
	path := "/home/kurtis/wrkspc/github.com/angelokurtis/go-rhonline"
	module, err := GetModuleFromPath(path)
	dieIfErr(err)
	_ = module

	abs, err := filepath.Abs(path)
	dieIfErr(err)

	directories, err := listDirectories(abs)
	dieIfErr(err)

	fset := token.NewFileSet()

	for _, directory := range directories {
		filter := func(info fs.FileInfo) bool { return true }
		pkgs, err := parser.ParseDir(fset, directory, filter, parser.ParseComments)
		dieIfErr(errors.WithStack(err))

		for _, pkg := range pkgs {
			dir := strings.Replace(directory, abs, module, 1)
			last := GetLastPackage(dir)
			if last != pkg.Name {
				slog.Info("", slog.String("pkg", dir), slog.String("alias", pkg.Name))
				// slog.Info("", slog.String("directory", directory), slog.String("pkg", dir), slog.String("alias", pkg.Name))
			} else {
				slog.Info("", slog.String("pkg", dir))
				// slog.Info("", slog.String("directory", directory), slog.String("pkg", dir))
			}

			for _, file := range pkg.Files {
				_ = file
				// file.Name
				// ast.Inspect(file, InspectNode)
			}
		}
	}
}

func GetLastPackage(pkg string) string {
	// Split the string into substrings based on "/"
	parts := strings.Split(pkg, "/")

	// Get the last substring from the array
	return parts[len(parts)-1]
}

func GetModuleFromPath(path string) (string, error) {
	modContent, err := os.ReadFile(filepath.Join(path, "go.mod"))
	if err != nil {
		return "", err
	}

	modFile, err := modfile.Parse("go.mod", modContent, nil)
	if err != nil {
		return "", err
	}

	return modFile.Module.Mod.Path, nil
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
