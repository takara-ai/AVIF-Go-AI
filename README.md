# AVIF-Go-AI

A native Go AVIF Encoder/Decoder written entirely by AI.

## Installation

To install the package, use:

```
go get github.com/takara-ai/AVIF-Go-AI
```

## Usage

### As a library

```go
import "github.com/takara-ai/AVIF-Go-AI/pkg/avif"

// Decoding
file, _ := os.Open("image.avif")
img, err := avif.Decode(file)

// Encoding
outFile, _ := os.Create("output.avif")
err := avif.Encode(outFile, img)
```

### Command-line tool

To encode a PNG to AVIF:

```
avif-tool encode -input input.png -output output.avif
```

To decode an AVIF to PNG:

```
avif-tool decode -input input.avif -output output.png
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
