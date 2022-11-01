# onemonth

A program that iterate the day of the month. I created it to generate the date text for the diary I am generating every month.

## Installation

```sh
go install github.com/tsuen4/onemonth@latest
```

## Usage

### Output the date of the current month

```sh
onemonth
```

#### Output:

Partially omitted.

```
2022/11/01: 
2022/11/02: 
2022/11/03: 
2022/11/04: 
2022/11/05: 
...
2022/11/26: 
2022/11/27: 
2022/11/28: 
2022/11/29: 
2022/11/30: 
```

### Output the date of the specified month:

```sh
onemonth -month 1
# Shorthand
onemonth -m 1
```

### Output the date of the current month of the specified year:

```sh
onemonth -year 2023
# Shorthand
onemonth -y 2023
```

### Specify output format:


Conform to Golang time format.

see: https://pkg.go.dev/time#pkg-constants

```sh
onemonth -layout 'Jan _2'
# shorthand
onemonth -l 'Jan _2'
```

#### Output: 

Partially omitted.

```
Nov  1
Nov  2
Nov  3
Nov  4
Nov  5
...
Nov 26
Nov 27
Nov 28
Nov 29
Nov 30
```

## Sample code

```golang
func main() {
	month, err := onemonth.New(2022, 1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	month.Iterate(func(day time.Time) {
        // some operation
	})
}
```
