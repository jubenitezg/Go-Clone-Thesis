package extractor

import (
	"strings"
	"testing"
)

const (
	TestFile = "test_assets/test.go"
	AstPath  = `add add,(Ident)^(FuncDecl)_(FuncType)_(FieldList)_(Field)_(Ident),x add,(Ident)^(FuncDecl)_(FuncType)_(FieldList)_(Field)_(Ident),y add,(Ident)^(FuncDecl)_(FuncType)_(FieldList)_(Field)_(Ident),int add,(Ident)^(FuncDecl)_(FuncType)_(FieldList)_(Field)_(Ident),int add,(Ident)^(FuncDecl)_(BlockStmt)_(DeclStmt)_(GenDecl)_(ValueSpec)_(Ident),z add,(Ident)^(FuncDecl)_(BlockStmt)_(DeclStmt)_(GenDecl)_(ValueSpec)_(Ident),int add,(Ident)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(Ident),z add,(Ident)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),x add,(Ident)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),y add,(Ident)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(BasicLit),2 add,(Ident)^(FuncDecl)_(BlockStmt)_(ReturnStmt)_(Ident),z x,(Ident)^(Field)_(Ident),y x,(Ident)^(Field)_(Ident),int x,(Ident)^(Field)^(FieldList)^(FuncType)_(FieldList)_(Field)_(Ident),int x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(DeclStmt)_(GenDecl)_(ValueSpec)_(Ident),z x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(DeclStmt)_(GenDecl)_(ValueSpec)_(Ident),int x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(Ident),z x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),x x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),y x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(BasicLit),2 x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ReturnStmt)_(Ident),z y,(Ident)^(Field)_(Ident),int y,(Ident)^(Field)^(FieldList)^(FuncType)_(FieldList)_(Field)_(Ident),int y,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(DeclStmt)_(GenDecl)_(ValueSpec)_(Ident),z y,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(DeclStmt)_(GenDecl)_(ValueSpec)_(Ident),int y,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(Ident),z y,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),x y,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),y y,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(BasicLit),2 y,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ReturnStmt)_(Ident),z int,(Ident)^(Field)^(FieldList)^(FuncType)_(FieldList)_(Field)_(Ident),int int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(DeclStmt)_(GenDecl)_(ValueSpec)_(Ident),z int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(DeclStmt)_(GenDecl)_(ValueSpec)_(Ident),int int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(Ident),z int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),x int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),y int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(BasicLit),2 int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ReturnStmt)_(Ident),z int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(DeclStmt)_(GenDecl)_(ValueSpec)_(Ident),z int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(DeclStmt)_(GenDecl)_(ValueSpec)_(Ident),int int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(Ident),z int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),x int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),y int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(BasicLit),2 int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ReturnStmt)_(Ident),z z,(Ident)^(ValueSpec)_(Ident),int z,(Ident)^(ValueSpec)^(GenDecl)^(DeclStmt)^(BlockStmt)_(AssignStmt:=)_(Ident),z z,(Ident)^(ValueSpec)^(GenDecl)^(DeclStmt)^(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),x z,(Ident)^(ValueSpec)^(GenDecl)^(DeclStmt)^(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),y z,(Ident)^(ValueSpec)^(GenDecl)^(DeclStmt)^(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(BasicLit),2 z,(Ident)^(ValueSpec)^(GenDecl)^(DeclStmt)^(BlockStmt)_(ReturnStmt)_(Ident),z int,(Ident)^(ValueSpec)^(GenDecl)^(DeclStmt)^(BlockStmt)_(AssignStmt:=)_(Ident),z int,(Ident)^(ValueSpec)^(GenDecl)^(DeclStmt)^(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),x int,(Ident)^(ValueSpec)^(GenDecl)^(DeclStmt)^(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),y int,(Ident)^(ValueSpec)^(GenDecl)^(DeclStmt)^(BlockStmt)_(AssignStmt:=)_(BinaryExpr:+)_(BasicLit),2 int,(Ident)^(ValueSpec)^(GenDecl)^(DeclStmt)^(BlockStmt)_(ReturnStmt)_(Ident),z z,(Ident)^(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),x z,(Ident)^(AssignStmt:=)_(BinaryExpr:+)_(ParenExpr)_(BinaryExpr:+)_(Ident),y z,(Ident)^(AssignStmt:=)_(BinaryExpr:+)_(BasicLit),2 z,(Ident)^(AssignStmt:=)^(BlockStmt)_(ReturnStmt)_(Ident),z x,(Ident)^(BinaryExpr:+)_(Ident),y x,(Ident)^(BinaryExpr:+)^(ParenExpr)^(BinaryExpr:+)_(BasicLit),2 x,(Ident)^(BinaryExpr:+)^(ParenExpr)^(BinaryExpr:+)^(AssignStmt:=)^(BlockStmt)_(ReturnStmt)_(Ident),z y,(Ident)^(BinaryExpr:+)^(ParenExpr)^(BinaryExpr:+)_(BasicLit),2 y,(Ident)^(BinaryExpr:+)^(ParenExpr)^(BinaryExpr:+)^(AssignStmt:=)^(BlockStmt)_(ReturnStmt)_(Ident),z 2,(BasicLit)^(BinaryExpr:+)^(AssignStmt:=)^(BlockStmt)_(ReturnStmt)_(Ident),z 
searchLinear searchLinear,(Ident)^(FuncDecl)_(FuncType)_(FieldList)_(Field)_(Ident),arr searchLinear,(Ident)^(FuncDecl)_(FuncType)_(FieldList)_(Field)_(ArrayType)_(Ident),int searchLinear,(Ident)^(FuncDecl)_(FuncType)_(FieldList)_(Field)_(Ident),x searchLinear,(Ident)^(FuncDecl)_(FuncType)_(FieldList)_(Field)_(Ident),int searchLinear,(Ident)^(FuncDecl)_(FuncType)_(FieldList)_(Field)_(Ident),int searchLinear,(Ident)^(FuncDecl)_(BlockStmt)_(ForStmt)_(AssignStmt::=)_(Ident),i searchLinear,(Ident)^(FuncDecl)_(BlockStmt)_(ForStmt)_(AssignStmt::=)_(BasicLit),0 searchLinear,(Ident)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(Ident),i searchLinear,(Ident)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),len searchLinear,(Ident)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),arr searchLinear,(Ident)^(FuncDecl)_(BlockStmt)_(ForStmt)_(IncDecStmt)_(Ident),i searchLinear,(Ident)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),arr searchLinear,(Ident)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),i searchLinear,(Ident)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(Ident),x searchLinear,(Ident)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i searchLinear,(Ident)^(FuncDecl)_(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 arr,(Ident)^(Field)_(ArrayType)_(Ident),int arr,(Ident)^(Field)^(FieldList)_(Field)_(Ident),x arr,(Ident)^(Field)^(FieldList)_(Field)_(Ident),int arr,(Ident)^(Field)^(FieldList)^(FuncType)_(FieldList)_(Field)_(Ident),int arr,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(AssignStmt::=)_(Ident),i arr,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(AssignStmt::=)_(BasicLit),0 arr,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(Ident),i arr,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),len arr,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),arr arr,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(IncDecStmt)_(Ident),i arr,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),arr arr,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),i arr,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(Ident),x arr,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i arr,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 int,(Ident)^(ArrayType)^(Field)^(FieldList)_(Field)_(Ident),x int,(Ident)^(ArrayType)^(Field)^(FieldList)_(Field)_(Ident),int int,(Ident)^(ArrayType)^(Field)^(FieldList)^(FuncType)_(FieldList)_(Field)_(Ident),int int,(Ident)^(ArrayType)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(AssignStmt::=)_(Ident),i int,(Ident)^(ArrayType)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(AssignStmt::=)_(BasicLit),0 int,(Ident)^(ArrayType)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(Ident),i int,(Ident)^(ArrayType)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),len int,(Ident)^(ArrayType)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),arr int,(Ident)^(ArrayType)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(IncDecStmt)_(Ident),i int,(Ident)^(ArrayType)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),arr int,(Ident)^(ArrayType)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),i int,(Ident)^(ArrayType)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(Ident),x int,(Ident)^(ArrayType)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i int,(Ident)^(ArrayType)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 x,(Ident)^(Field)_(Ident),int x,(Ident)^(Field)^(FieldList)^(FuncType)_(FieldList)_(Field)_(Ident),int x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(AssignStmt::=)_(Ident),i x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(AssignStmt::=)_(BasicLit),0 x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(Ident),i x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),len x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),arr x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(IncDecStmt)_(Ident),i x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),arr x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),i x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(Ident),x x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i x,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 int,(Ident)^(Field)^(FieldList)^(FuncType)_(FieldList)_(Field)_(Ident),int int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(AssignStmt::=)_(Ident),i int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(AssignStmt::=)_(BasicLit),0 int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(Ident),i int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),len int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),arr int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(IncDecStmt)_(Ident),i int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),arr int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),i int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(Ident),x int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(AssignStmt::=)_(Ident),i int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(AssignStmt::=)_(BasicLit),0 int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(Ident),i int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),len int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),arr int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(IncDecStmt)_(Ident),i int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),arr int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),i int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(Ident),x int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ForStmt)_(BlockStmt)_(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i int,(Ident)^(Field)^(FieldList)^(FuncType)^(FuncDecl)_(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 i,(Ident)^(AssignStmt::=)_(BasicLit),0 i,(Ident)^(AssignStmt::=)^(ForStmt)_(BinaryExpr:<)_(Ident),i i,(Ident)^(AssignStmt::=)^(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),len i,(Ident)^(AssignStmt::=)^(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),arr i,(Ident)^(AssignStmt::=)^(ForStmt)_(IncDecStmt)_(Ident),i i,(Ident)^(AssignStmt::=)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),arr i,(Ident)^(AssignStmt::=)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),i i,(Ident)^(AssignStmt::=)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(Ident),x i,(Ident)^(AssignStmt::=)^(ForStmt)_(BlockStmt)_(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i i,(Ident)^(AssignStmt::=)^(ForStmt)^(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 0,(BasicLit)^(AssignStmt::=)^(ForStmt)_(BinaryExpr:<)_(Ident),i 0,(BasicLit)^(AssignStmt::=)^(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),len 0,(BasicLit)^(AssignStmt::=)^(ForStmt)_(BinaryExpr:<)_(CallExpr)_(Ident),arr 0,(BasicLit)^(AssignStmt::=)^(ForStmt)_(IncDecStmt)_(Ident),i 0,(BasicLit)^(AssignStmt::=)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),arr 0,(BasicLit)^(AssignStmt::=)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),i 0,(BasicLit)^(AssignStmt::=)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(Ident),x 0,(BasicLit)^(AssignStmt::=)^(ForStmt)_(BlockStmt)_(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i 0,(BasicLit)^(AssignStmt::=)^(ForStmt)^(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 i,(Ident)^(BinaryExpr:<)_(CallExpr)_(Ident),len i,(Ident)^(BinaryExpr:<)_(CallExpr)_(Ident),arr i,(Ident)^(BinaryExpr:<)^(ForStmt)_(IncDecStmt)_(Ident),i i,(Ident)^(BinaryExpr:<)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),arr i,(Ident)^(BinaryExpr:<)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),i i,(Ident)^(BinaryExpr:<)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(Ident),x i,(Ident)^(BinaryExpr:<)^(ForStmt)_(BlockStmt)_(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i i,(Ident)^(BinaryExpr:<)^(ForStmt)^(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 len,(Ident)^(CallExpr)_(Ident),arr len,(Ident)^(CallExpr)^(BinaryExpr:<)^(ForStmt)_(IncDecStmt)_(Ident),i len,(Ident)^(CallExpr)^(BinaryExpr:<)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),arr len,(Ident)^(CallExpr)^(BinaryExpr:<)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),i len,(Ident)^(CallExpr)^(BinaryExpr:<)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(Ident),x len,(Ident)^(CallExpr)^(BinaryExpr:<)^(ForStmt)_(BlockStmt)_(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i len,(Ident)^(CallExpr)^(BinaryExpr:<)^(ForStmt)^(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 arr,(Ident)^(CallExpr)^(BinaryExpr:<)^(ForStmt)_(IncDecStmt)_(Ident),i arr,(Ident)^(CallExpr)^(BinaryExpr:<)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),arr arr,(Ident)^(CallExpr)^(BinaryExpr:<)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),i arr,(Ident)^(CallExpr)^(BinaryExpr:<)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(Ident),x arr,(Ident)^(CallExpr)^(BinaryExpr:<)^(ForStmt)_(BlockStmt)_(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i arr,(Ident)^(CallExpr)^(BinaryExpr:<)^(ForStmt)^(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 i,(Ident)^(IncDecStmt)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),arr i,(Ident)^(IncDecStmt)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(IndexExpr)_(Ident),i i,(Ident)^(IncDecStmt)^(ForStmt)_(BlockStmt)_(IfStmt)_(BinaryExpr:==)_(Ident),x i,(Ident)^(IncDecStmt)^(ForStmt)_(BlockStmt)_(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i i,(Ident)^(IncDecStmt)^(ForStmt)^(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 arr,(Ident)^(IndexExpr)_(Ident),i arr,(Ident)^(IndexExpr)^(BinaryExpr:==)_(Ident),x arr,(Ident)^(IndexExpr)^(BinaryExpr:==)^(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i arr,(Ident)^(IndexExpr)^(BinaryExpr:==)^(IfStmt)^(BlockStmt)^(ForStmt)^(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 i,(Ident)^(IndexExpr)^(BinaryExpr:==)_(Ident),x i,(Ident)^(IndexExpr)^(BinaryExpr:==)^(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i i,(Ident)^(IndexExpr)^(BinaryExpr:==)^(IfStmt)^(BlockStmt)^(ForStmt)^(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 x,(Ident)^(BinaryExpr:==)^(IfStmt)_(BlockStmt)_(ReturnStmt)_(Ident),i x,(Ident)^(BinaryExpr:==)^(IfStmt)^(BlockStmt)^(ForStmt)^(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 i,(Ident)^(ReturnStmt)^(BlockStmt)^(IfStmt)^(BlockStmt)^(ForStmt)^(BlockStmt)_(ReturnStmt)_(UnaryExpr:-)_(BasicLit),1 `
	AstPathHash = `add add,9457696349906653245,x add,9457696349906653245,y add,9457696349906653245,int add,9457696349906653245,int add,6672165243388416080,z add,6672165243388416080,int add,13995975484808820024,z add,14909724661373915795,x add,14909724661373915795,y add,5686095246127261138,2 add,14462832031105918338,z x,16408190609135934237,y x,16408190609135934237,int x,6372329674272792047,int x,15549807665554965433,z x,15549807665554965433,int x,1327511780129170289,z x,11040555812031836362,x x,11040555812031836362,y x,3807608797982075435,2 x,11102109797395058907,z y,16408190609135934237,int y,6372329674272792047,int y,15549807665554965433,z y,15549807665554965433,int y,1327511780129170289,z y,11040555812031836362,x y,11040555812031836362,y y,3807608797982075435,2 y,11102109797395058907,z int,6372329674272792047,int int,15549807665554965433,z int,15549807665554965433,int int,1327511780129170289,z int,11040555812031836362,x int,11040555812031836362,y int,3807608797982075435,2 int,11102109797395058907,z int,15549807665554965433,z int,15549807665554965433,int int,1327511780129170289,z int,11040555812031836362,x int,11040555812031836362,y int,3807608797982075435,2 int,11102109797395058907,z z,17148070876960847283,int z,2638379050080874229,z z,10981927456071765262,x z,10981927456071765262,y z,12701333877964123063,2 z,7003335801671362471,z int,2638379050080874229,z int,10981927456071765262,x int,10981927456071765262,y int,12701333877964123063,2 int,7003335801671362471,z z,9119042879847095144,x z,9119042879847095144,y z,8549896659131320925,2 z,2083001101555384231,z x,15087356045205734748,y x,1547970451561534269,2 x,9138152563133994523,z y,1547970451561534269,2 y,9138152563133994523,z 2,9437554797313800888,z 
searchLinear searchLinear,9457696349906653245,arr searchLinear,975769635872572848,int searchLinear,9457696349906653245,x searchLinear,9457696349906653245,int searchLinear,9457696349906653245,int searchLinear,9299501546996317931,i searchLinear,1467506151180932122,0 searchLinear,715861570420353937,i searchLinear,18070668761863282130,len searchLinear,18070668761863282130,arr searchLinear,13392139858935425639,i searchLinear,16834231136120450032,arr searchLinear,16834231136120450032,i searchLinear,17539073895131998131,x searchLinear,8373707528811830752,i searchLinear,1314776747326531150,1 arr,13451387741941980112,int arr,1039595498909291230,x arr,1039595498909291230,int arr,6372329674272792047,int arr,18167791089916908044,i arr,13967126708461234535,0 arr,4595706360506324972,i arr,13464290231831765849,len arr,13464290231831765849,arr arr,8941941359826873438,i arr,811909861314682449,arr arr,811909861314682449,i arr,18007213482334779866,x arr,2137854948785079855,i arr,4592867627706957233,1 int,8240408341950714120,x int,8240408341950714120,int int,2582376111199657649,int int,12636474495107261246,i int,15298989274383758533,0 int,4647323896924806274,i int,13350661840711275671,len int,13350661840711275671,arr int,1695522206467338392,i int,15194342694752373143,arr int,15194342694752373143,i int,311723947176465684,x int,13343765195791046217,i int,12490389356345559407,1 x,16408190609135934237,int x,6372329674272792047,int x,18167791089916908044,i x,13967126708461234535,0 x,4595706360506324972,i x,13464290231831765849,len x,13464290231831765849,arr x,8941941359826873438,i x,811909861314682449,arr x,811909861314682449,i x,18007213482334779866,x x,2137854948785079855,i x,4592867627706957233,1 int,6372329674272792047,int int,18167791089916908044,i int,13967126708461234535,0 int,4595706360506324972,i int,13464290231831765849,len int,13464290231831765849,arr int,8941941359826873438,i int,811909861314682449,arr int,811909861314682449,i int,18007213482334779866,x int,2137854948785079855,i int,4592867627706957233,1 int,18167791089916908044,i int,13967126708461234535,0 int,4595706360506324972,i int,13464290231831765849,len int,13464290231831765849,arr int,8941941359826873438,i int,811909861314682449,arr int,811909861314682449,i int,18007213482334779866,x int,2137854948785079855,i int,4592867627706957233,1 i,1636894060408599206,0 i,3658841908228236485,i i,5273136842523490254,len i,5273136842523490254,arr i,1178514676072191763,i i,15545262392509369260,arr i,15545262392509369260,i i,10831948888538945031,x i,4073979834493451284,i i,8247081323652443683,1 0,14099309914766389804,i 0,12656715338222101913,len 0,12656715338222101913,arr 0,13581334232548806302,i 0,5840161503592752401,arr 0,5840161503592752401,i 0,3365115432051286426,x 0,901080167913338991,i 0,7062952598116768756,1 i,14107167364760014670,len i,14107167364760014670,arr i,4671797722808181061,i i,11949697498428939074,arr i,11949697498428939074,i i,6600683539896740737,x i,13687271608717705446,i i,4003866583962665389,1 len,1271882245012735966,arr len,3547495804110384081,i len,6824363962785613358,arr len,6824363962785613358,i len,13466080716744890389,x len,10636316600974723498,i len,136958888530746209,1 arr,3547495804110384081,i arr,6824363962785613358,arr arr,6824363962785613358,i arr,13466080716744890389,x arr,10636316600974723498,i arr,136958888530746209,1 i,16602223733131735088,arr i,16602223733131735088,i i,12952370843505420275,x i,10859832161346033056,i i,9019924669979022959,1 arr,18440238959439346254,i arr,2364426871906791767,x arr,307992080962863938,i arr,9095216703006312383,1 i,2364426871906791767,x i,307992080962863938,i i,9095216703006312383,1 x,1575516384664693006,i x,8801239681333522171,1 i,14531935199101260793,1 `
)

func TestExtractorHashedPaths(t *testing.T) {
	ex, err := NewExtractor(TestFile, true, 200, 200)
	if err != nil {
		t.Error("Error creating extractor:", err)
		return
	}
	testPathsExtraction(t, AstPathHash, ex)
}

func TestExtractorPaths(t *testing.T) {
	ex, err := NewExtractor(TestFile, false, 200, 200)
	if err != nil {
		t.Error("Error creating extractor:", err)
		return
	}
	testPathsExtraction(t, AstPath, ex)
}

func testPathsExtraction(t *testing.T, path string, ex *Extractor) {
	myPaths := ex.GenerateProgramAstPaths()
	compPaths := strings.Split(path, "\n")

	if len(myPaths) != len(compPaths) {
		t.Error("Different number of paths")
		return
	}
	for _, myPath := range myPaths {
		found := false
		for _, compPath := range compPaths {
			if myPath == compPath {
				found = true
				break
			}
		}
		if !found {
			t.Error("Different paths")
			return
		}
	}
	t.Log("Same paths")
}
