package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/takara-ai/AVIF-Go-AI"
)

func main() {
	encodeCmd := flag.NewFlagSet("encode", flag.ExitOnError)
	encodeInput := encodeCmd.String("input", "", "Input PNG file")
	encodeOutput := encodeCmd.String("output", "", "Output AVIF file")

	decodeCmd := flag.NewFlagSet("decode", flag.ExitOnError)
	decodeInput := decodeCmd.String("input", "", "Input AVIF file")
	decodeOutput := decodeCmd.String("output", "", "Output PNG file")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'encode' or 'decode' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "encode":
		encodeCmd.Parse(os.Args[2:])
		if *encodeInput == "" || *encodeOutput == "" {
			encodeCmd.PrintDefaults()
			os.Exit(1)
		}
		encodeFile(*encodeInput, *encodeOutput)
	case "decode":
		decodeCmd.Parse(os.Args[2:])
		if *decodeInput == "" || *decodeOutput == "" {
			decodeCmd.PrintDefaults()
			os.Exit(1)
		}
		decodeFile(*decodeInput, *decodeOutput)
	default:
		fmt.Println("Expected 'encode' or 'decode' subcommands")
		os.Exit(1)
	}
}

func encodeFile(input, output string) {
	// Open input file
	f, err := os.Open(input)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// Decode PNG
	img, err := png.Decode(f)
	if err != nil {
		fmt.Printf("Error decoding PNG: %v\n", err)
		os.Exit(1)
	}

	// Create output file
	out, err := os.Create(output)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	// Encode AVIF
	if err := avif.Encode(out, img); err != nil {
		fmt.Printf("Error encoding AVIF: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("AVIF file created successfully")
}

func decodeFile(input, output string) {
	// Open input file
	f, err := os.Open(input)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// Decode AVIF
	img, err := avif.Decode(f)
	if err != nil {
		fmt.Printf("Error decoding AVIF: %v\n", err)
		os.Exit(1)
	}

	// Create output file
	out, err := os.Create(output)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	// Encode PNG
	if err := png.Encode(out, img); err != nil {
		fmt.Printf("Error encoding PNG: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("PNG file created successfully")
}