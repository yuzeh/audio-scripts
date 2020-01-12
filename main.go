package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	inPtr := flag.String("in", "", "Input file name")
	outPtr := flag.String("out", "", "Output file")

	flag.Parse()

	if *inPtr == "" {
		fmt.Println("Input file required")
		os.Exit(1)
	}

	if *outPtr == "" {
		fmt.Println("Output file required")
		os.Exit(1)
	}

	cmd := exec.Command("ffmpeg", "-y", "-i", *inPtr, "-codec:a", "libmp3lame", "-filter:a", "loudnorm=I=-16", *outPtr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

	cmd.Wait()
}
