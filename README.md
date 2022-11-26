# randrange
A CLI tool to generate random ranges

## Building
`go build .`
Then move the executable into your `$PATH` and run it whenever you want!

## Usage
randrange has two modes: Standard mode, and math mode (interval notation)

Your range also needs to be in parentheses. This is because otherwise your shell would try to parse it.

## Standard Mode
`randrange -s "0-20"`
This is pretty much the range that you expect. The above command will generate a random number from 0-20, including 0 and 20

## Math Mode
`randrange -m "[0,30)"`
This mode uses interval notation. For a quick summary, [] means inclusive, and () means exclusive, so
[0, 20] Would include 0 and 20, and (0, 20) wouldn't include either. (0, 20) is the same as [1, 19]

The program does not support negative ranges or decimal points. The following are invalid:
 - `randrange -m "[-10, -3]`
 - `randrange -s "0.5-0.7"`
