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

- [Quick Start](#quick-start)
- [Licensing](#licensing)

---

### Why?

Working with the Go Standard Library's `sync.Map` can be frustrating at times
and produce ugly code, as type casting is required. `sync.Map` also
uses `unsafe` which isn't necessarily a problem by itself, however avoiding its
usage is always a win.

---

### Quick Start

Store is both safe and incredibly easy to use, the below example shows how to
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
	store := new(store.Store[string, int])

	store.Set("two", 2)

	if two, ok := store.Get("two"); ok {
		fmt.Println(two) // 2

		store.Set("four", two+2) // stores 4 under "four" as generics remove the need for type casting
	}
}
```

Store only accepts comparable types for keys, removing the worry of using
impossible key types. Stores are incomparable, and cannot be declared using
`Store{}`.

---

### Licensing

Store is licensed under the `BSD-3-Clause` license found [here](/LICENSE).