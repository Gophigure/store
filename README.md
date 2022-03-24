<!--suppress HtmlDeprecatedAttribute -->
<img align="center" alt="store banner" src="/.github/assets/GophigureBannerStore.svg">

---

<h4 align="center"><i>
Provides an alternative to the Go Standard Library's <code>sync.Map</code> with
added type
safety using <code>1.18</code> generics.
</i></h4>

---

### Table of Contents

- [Why?](#why)
- [Features](#features)
- [Quick Start](#quick-start)
- [Licensing](#licensing)

---

### Why?

Working with the Go Standard Library's `sync.Map` can be frustrating at times
and produce ugly code, as type casting is required. `sync.Map` also
uses `unsafe` which isn't necessarily a problem by itself, however avoiding its
usage is always a win.

---

### Features

- Supports concurrency
- Type-safe
- No usage of `unsafe.Pointer`
- `Store.Set`  
  *Creates or overrides a key and stores a value associated with it.*
- `Store.Get`  
  *Attempts to retrieve the value under a key in the store.*
- `Store.GetOrSet`  
  *Attempts to retrieve the value under a key in the store, if it fails it will
  store the provided value instead.*
- `Store.Pluck`  
  *Removes a key and its associated value from the store, but returns the
  removed value.*
- `Store.Delete`  
  *Removes a key and its associated value from the store.*
- `Store.ForEach`  
  *Takes a function as a parameter and iterates over the Store's values, calling
  the function for each value.*

---

### Quick Start

Store is incredibly easy to use, the below example shows how to
both get and set keys. You'll also notice how no type casting is needed when
retrieving values thanks to generics.

If you'd like to store multiple types, use interfaces and *then* type-check &
cast.

```go
package main

import (
	"fmt"
	"github.com/Gophigure/store"
)

func main() {
	myStore := new(store.Store[string, int]) // you must initialize a Store using new.

	myStore.Set("two", 2)

	if two, ok := myStore.Get("two"); ok {
		fmt.Println(two) // 2

		myStore.Set("four", two+2) // stores 4 under "four" as generics remove the need for type casting
	}
}
```

Store only accepts comparable types for keys, removing the worry of using
impossible key types. Stores are incomparable, and cannot be declared using
`Store{}`.

---

### Licensing

Store is licensed under the `BSD-3-Clause` license found [here](/LICENSE).