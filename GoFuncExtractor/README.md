# GoFuncExtractor

A go function extractor

## Pre-requisites
* Go 1.21 or higher

## Dependencies
From the root of the project run:

```bash
go mod tidy && go mod vendor
```

## Install
From the root of the project run:

```bash
go build -o gfunc
```

## Run

```bash
./gfunc --project-path <path_to_project> [--single-line]
```

## Example

Given the following [file](func_extractor/test_assets/file.go)

Run the following command:
```bash
./gfunc --project-path func_extractor/test_assets --single-line
```

The output will be:
```
{"code":"func (p *Person) SayHello() string {\n\treturn \"Hello \" + p.Name\n}","id":"89d4bfd0-cbef-4065-a7f8-34d9d94053ca"}
{"code":"func (p *Person) SayAge() string {\n\treturn \"I am \" + strconv.Itoa(p.Age)\n}","id":"80eaf7fb-933a-4e14-913c-3bab6b6f122a"}
{"code":"func NewPerson(name string, age int) *Person {\n\treturn &Person{Name: name, Age: age}\n}","id":"226477d4-5537-4dfd-a1bc-f8a1c82c683d"}
```
