# colorhash

Map arbitrary strings and byte streams to deterministic colors from a given palette.

## Features

- **Deterministic hashing** — same input always produces the same color
- **FNV-64 based** — fast, well-distributed hash function
- **OKLCH palette generation** — perceptually uniform color palettes
- **ANSI terminal colors** — built-in escape code wrappers for terminal output
- **True color support** — 24-bit RGB terminal color output

## Install

```bash
go get github.com/taigrr/colorhash
```

## Usage

### Map a string to a color

```go
import (
    "github.com/taigrr/colorhash"
    "github.com/taigrr/simplecolorpalettes/flatui"
)

// Pick a color from a palette based on a string
c := colorhash.StringToColor(flatui.Palette, "alice")
```

### Hash a byte stream

```go
hash, err := colorhash.HashReader(reader)
if err != nil {
    return err
}
fmt.Println(hash)
```

### Generate an OKLCH palette

```go
// 8 evenly-spaced hues at lightness 0.7, chroma 0.15
palette := colorhash.GenerateOKLCHPalette(8, 0.7, 0.15)
```

### Terminal color output

```go
fmt.Println(colorhash.Red("error"))
fmt.Println(colorhash.Green("success"))
fmt.Println(colorhash.Info("info message"))
```

## License

0BSD
