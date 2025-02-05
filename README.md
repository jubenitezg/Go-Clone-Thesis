# Code cloning in Go open source projects: A prevalence and structure analysis on GitHub

Escuela Colombiana de Ingeniería Master Thesis on Computer Science - Bogotá, Colombia

## Example of Types of Clones

#### Initial Function
```go
func factorial(n int) int {
	j := 1
	for i := 1; i <= n; i++ {
		j *= i
	}
	return j
}
```

#### Type 1 - Exact Clone
```go
func factorial(n int) int {
	j := 1
	for i := 1; i <= n; i++ {
		j *= i
	}
	return j
}
```

#### Type 2 - Renamed clone
```go
func getFactorial(number int) int {
	acc := 1
	for curr := 1; curr <= number; curr++ {
		acc *= curr
	}
	return acc
}
```

#### Type 3 - Near miss clone
```go
func factorialAdd(n int) (int, int) {
	var accFactorial int
	var accSum int
	for curr := 1; curr <= n; curr++ {
		accFactorial *= curr
		accSum += curr
	}
	return accFactorial, accSum
}
```

#### Type 4 - Semantic clone
```go
func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)
}
```

## Available Go repositories on GitHub

Total Go repositories on GitHub (excluding forks): `1,332,711`

Total Go repositories with more than 100 stars: `15,863`

## Author

- Julián Benítez Gutiérrez - [julian.benitez@mail.escuelaing.edu.co](mailto:julian.benitez@mail.escuelaing.edu.co)

## Director

- Héctor Fabio Cadavid Rengifo

## Co-director

- Wilmer Edicson Garzón Alfonso
