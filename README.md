# go-perm
Go implementation of QuickPerm.

## Install

```sh
go get github.com/ktnyt/go-perm
```

## Usage

### Code

```go
package main

import (
	"fmt"

	"github.com/ktnyt/go-perm"
)

func main() {
	a := []string{"foo", "bar", "baz"}
	p := perm.Permute(a)
	for p.Next() {
		fmt.Println(a)
	}
}
```

### Output

```
[foo bar baz]
[bar foo baz]
[baz foo bar]
[foo baz bar]
[bar baz foo]
[baz bar foo]
```

## Algorithm

This package implements the [*Counting* QuickPerm algorithm](https://www.quickperm.org/quickperm.php).
The reason for adopting the *counting* version is simply due to the fact that in Go, initializing a slice containing zeros is easier than one containing its own index.
For an in-depth look into how QuickPerm works, check out the [QuickPerm website](https://www.quickperm.org/).
