package main

import (
	"log"
	"os/exec"
	"net/http"
	"bytes"
	"fmt"
)

func scaleHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("ffmpeg", "-i", "test.mp4", "-s", "960x720", "test_720p.mp4")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
	fmt.Printf("in all caps: %q\n", stderr.String())
}
