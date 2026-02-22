# colorhash

Deterministic color assignment from arbitrary input. Feed it a string or byte stream and get back a consistent color every time.

## Features

- **Deterministic** — same input always produces the same color
- **Pluggable palettes** — bring your own `ColorSet` or use [simplecolorpalettes](https://github.com/taigrr/simplecolorpalettes)
- **Multiple output formats** — `color.Color`, ANSI escape codes, true-color terminal strings
- **String and byte input** — hash strings directly or stream bytes via `io.Reader`

## Install

```bash
go get github.com/taigrr/colorhash
```

## Usage

### Hash a string to a color

```go
import (
    "github.com/taigrr/colorhash"
    "github.com/taigrr/simplecolorpalettes/palettes/html"
)

palette := html.GetPalette() // or any ColorSet
c := colorhash.StringToColor(palette, "username")
// c is a deterministic color.Color
```

### ANSI terminal colors

```go
fmt.Println(colorhash.Red("error message"))
fmt.Println(colorhash.Green("success"))
fmt.Println(colorhash.BIYellow("bold high-intensity yellow"))
```

### Automatic color from a palette

```go
sp := colorhash.CreateStringerPalette(false, false, palette)
colored := sp.GetString("username") // wraps "username" in its assigned color
```

## License

0BSD
