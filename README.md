# Components

Reusable, composable and testable frontend components to be used alongside the [gomponents](https://github.com/maragudk/gomponents) library.

## Quickstart / Installation

Add the library to your project with:
```shell
GOPRIVATE=github.com/geneton-ltd/* go get github.com/geneton-ltd/syder_components
```

> [!WARNING]
> If you use SSH login for Github (you can check with `ssh -T git@github.com`), you need to tell git to use SSH instead of HTTPS with `git config --global url."git@github.com:".insteadOf "https://github.com/"`.

> [!NOTE] 
> This library is supposed to be used alongside gomponents. If you do not have them installed, you can do so with:
> `go get maragu.dev/gomponents`

Get started with this simple `main.go` file:
```go
package main

import (
	"net/http"

    sc "github.com/geneton-ltd/syder_components"
	g "maragu.dev/gomponents"
)

func main() {
	mux := http.NewServeMux()
    
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := sc.Button(
            sc.WithID("helloWorldBtnID"),
            sc.WithNodes(g.Text("Hello, world!")),
        )
        page.Render(w)
	})

	http.ListenAndServe(":8000", mux)
}
```

## Components overview
TODO

## Usage examples
TODO

## Development
TODO

## Core concepts
TODO

## Customizing styles
TODO

## Edge cases
TODO

## License
Syder Components is available under the MIT license. See the LICENSE file for more info.
