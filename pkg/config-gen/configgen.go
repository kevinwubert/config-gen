package configgen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strconv"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
)

type Data struct {
	PackageName   string
	Fields        []Field
	HasFieldTypes HasFieldTypes
}

type Field struct {
	Name        Name
	Type        string
	Description string
	IsSecret    bool
}

type Name struct {
	PascalCase         string
	KebabCase          string
	ScreamingSnakeCase string
}

type HasFieldTypes struct {
	HasString   bool
	HasInt      bool
	HasBool     bool
	HasDuration bool
}

func Generate(filename string, prefix string) error {
	data, err := Parse(filename, prefix)
	if err != nil {
		return errors.Wrap(err, "parsing file")
	}

	bytes := Render(data)

	dir := filepath.Dir(filename)
	err = ioutil.WriteFile(filepath.Join(dir, "config.gen.go"), bytes, 0644)
	if err != nil {
		return errors.Wrap(err, "writing file")
	}

	return nil
}

func Parse(filename string, prefix string) (Data, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return Data{}, err
	}

	d := Data{
		PackageName: "main",
		Fields:      []Field{},
		HasFieldTypes: HasFieldTypes{
			HasString:   false,
			HasInt:      false,
			HasBool:     false,
			HasDuration: false,
		},
	}

	ast.Inspect(file, func(x ast.Node) bool {
		// Try switching to *ast.TypeSpec to check name of struct first
		s, ok := x.(*ast.StructType)
		if !ok {
			return true
		}

		for _, field := range s.Fields.List {
			// this fails if the struct doesn't have a struct tag
			tag := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])
			name := field.Names[0].Name
			fieldType := "unknown"
			switch fmt.Sprintf("%s", field.Type) {
			case "string":
				fieldType = "string"
				d.HasFieldTypes.HasString = true
			case "int":
				fieldType = "int"
				d.HasFieldTypes.HasInt = true
			case "bool":
				fieldType = "bool"
				d.HasFieldTypes.HasBool = true
			case "&{time Duration}":
				fieldType = "duration"
				d.HasFieldTypes.HasDuration = true
			}

			isSecret, err := strconv.ParseBool(tag.Get("secret"))
			if err != nil {
				isSecret = false
			}

			d.Fields = append(d.Fields, Field{
				Name: Name{
					PascalCase:         name,
					KebabCase:          strcase.ToKebab(name),
					ScreamingSnakeCase: strcase.ToScreamingSnake(fmt.Sprintf("%s_%s", prefix, name)),
				},
				Type:        fieldType,
				Description: tag.Get("description"),
				IsSecret:    isSecret,
			})
		}
		return false
	})

	return d, nil
}

func Render(data Data) []byte {
	tmpl := template.Must(template.ParseFiles("pkg/config-gen/config.gen.go.template"))

	buf := bytes.NewBufferString("")
	tmpl.Execute(buf, data)

	return buf.Bytes()
}
