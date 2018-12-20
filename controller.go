package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"net/http"
	"bytes"
	"encoding/json"
	"strings"
	"strconv"
)

type Body struct {
	Name string
}

func scaleHandler(w http.ResponseWriter, r *http.Request) {
	var body Body
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
	    panic(err)
	}
	json.Unmarshal([]byte(bodyByte), &body)
	scaler(conf.Path + "/" + body.Name, videoSize(conf.Path + "/" + body.Name))
}

func scaler(name string, size string) {
    log.Println("ffmpeg, -i, " + name + ", -s, " + size + ", " + strings.Split(name, ".")[0] + "_720p." + strings.Split(name, ".")[1])
	cmd := exec.Command("ffmpeg", "-i", name, "-s", size, strings.Split(name, ".")[0] + "_720p." + strings.Split(name, ".")[1])
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("in all caps: %q\n", out.String())
	log.Printf("in all caps: %q\n", stderr.String())
}

func videoSize(name string) int {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "csv=s=x:p=0", name)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("in all caps: %q\n", out.String())
	height, err := strconv.Atoi(strings.Split(out.String(), "x")[1])
	if err != nil {
		log.Fatal(err)
	}
	if height >= 1024 {
		return 1024
	} else if height >=720  {
		return 720
	} else if height >= 480  {
		return 480
	} else if height >= 320  {
		return 320
	} else if height >= 280  {
		return 280
	} else if height >= 144  {
		return 144
	}
	return -1
}
