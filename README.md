#funsafe

Experimental package for unsetting the read-only flags from `reflect.Values`, bypassing [package export guarantees](https://golang.org/ref/spec#Exported_identifiers).

Example:

```go
//
// package private
//
// type Example struct {
//     a int
// }

func main() {

	v := &private.Example{}

	fmt.Printf("%#v\n", v) // &private.Example{a:0}

	av := reflect.ValueOf(v).Elem().FieldByName("a")

	funsafe.MakeSettable(&av)

	av.SetInt(1)

	fmt.Printf("%#v\n", v) // &private.Example{a:1}

}
```

*For research only. Violating package scoping guarantees is the wrong approach for production code.  However it can produce some interesting opportunities for toying with the go stdlib and other imported packages.  Use at your own risk!*

LICENSE: [MIT](https://raw.githubusercontent.com/jasonmoo/funsafe/master/LICENSE)