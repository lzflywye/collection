package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

func generateUniqueOutputPath(inputPath string) (string, error) {
	ext := filepath.Ext(inputPath)
	baseName := strings.TrimSuffix(inputPath, ext)
	dir := filepath.Dir(inputPath)
	name := filepath.Base(baseName)

	i := 1
	for {
		outputPath := filepath.Join(dir, fmt.Sprintf("%s_%d.jpeg", name, i))
		if _, err := os.Stat(outputPath); os.IsNotExist(err) {
			return outputPath, nil
		} else if err != nil {
			return "", fmt.Errorf("failed to stat output file '%s': %w", outputPath, err)
		}
		i++
		if i > 1000 {
			return "", fmt.Errorf("could not generate a unique output path for '%s'", inputPath)
		}
	}
}

func convertImageToJpeg(inputPath string, outputPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file '%s': %w", inputPath, err)
	}
	defer inputFile.Close()

	config, format, err := image.DecodeConfig(inputFile)
	if err != nil {
		return fmt.Errorf("failed to decode image config from '%s': %w", inputPath, err)
	}
	fmt.Println("Width:", config.Width, "Height:", config.Height, "Format:", format)

	_, err = inputFile.Seek(0, 0)

	img, _, err := image.Decode(inputFile)
	if err != nil {
		return fmt.Errorf("failed to decode image from '%s': %w", inputPath, err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file '%s': %w", outputPath, err)
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, img, &jpeg.Options{Quality: 90})
	if err != nil {
		return fmt.Errorf("failed to encode image to JPEG '%s': %w", outputPath, err)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run main.go <input_file> [<output_jpeg_file>]")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	var outputPath string

	if len(os.Args) == 3 {
		outputPath = os.Args[2]
		if _, err := os.Stat(outputPath); !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: Output file '%s' already exists.\n", outputPath)
			os.Exit(1)
		}
	} else {
		var err error
		outputPath, err = generateUniqueOutputPath(inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating output path: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Using automatically generated output path: %s\n", outputPath)
	}

	if err := convertImageToJpeg(inputPath, outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error during conversion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully converted '%s' to '%s'\n", inputPath, outputPath)
}
