# GoExtractor
A go AST path extractor

## Pre-requisites

* Go 1.21 or higher 

## Install
From the root of the project run:

```bash
go build -o gextr
```

## Run

```bash
./gextr --file <path_to_file> [--hash --max-length <max_length> --max-width <max_width> --clean-print]
```

## Tests

```bash
go test ./...
```

## Example

Given the following go file:

```go
package test_assets

func add(x, y int) int {
	var z int
	z = (x + y) + 2
	return z
}
```

Run the following command:
```bash
./gextr --file test.go --clean-print --max-length 5
```

The output will be:
```
add
add,(Ident)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(Ident),z
add,(Ident)^(FuncDecl)_(BlockStmt)_(ReturnStmt)_(Ident),z
x,(Ident)^(Field)_(Ident),y
x,(Ident)^(Field)_(Ident),int
y,(Ident)^(Field)_(Ident),int
z,(Ident)^(ValueSpec)_(Ident),int
z,(Ident)^(AssignStmt:=)_(BinaryExpr:+)_(BasicLit),2
z,(Ident)^(AssignStmt:=)^(BlockStmt)_(ReturnStmt)_(Ident),z
x,(Ident)^(BinaryExpr:+)_(Ident),y
x,(Ident)^(BinaryExpr:+)^(ParenExpr)^(BinaryExpr:+)_(BasicLit),2
y,(Ident)^(BinaryExpr:+)^(ParenExpr)^(BinaryExpr:+)_(BasicLit),2
```

## Author
* Julián Benítez Gutiérrez - [julian.benitez@mail.escuelaing.edu.co](mailto:julian.benitez@mail.escuelaing.edu.co)