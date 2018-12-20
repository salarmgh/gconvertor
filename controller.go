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
	size := videoSize(conf.Path + "/" + body.Name)
	scaler(conf.Path + "/" + body.Name, size)
}

func scaler(name string, size int) {
	log.Println("name: ", name)
	log.Println("size: ", size)
	sizes := map[int]string{
    720: "960x720",
	480: "640x480",
    360: "480x360",
    240: "320x240",
	144: "256x144",
    }
    log.Println("ffmpeg, -i, " + name + ", -s, " + sizes[size] + ", " + strings.Split(name, ".")[0] + "_" + strconv.Itoa(size) + "p." + strings.Split(name, ".")[1])
	cmd := exec.Command("ffmpeg", "-i", name, "-s", sizes[size], strings.Split(name, ".")[0] + "_" + strconv.Itoa(size) + "p." + strings.Split(name, ".")[1])
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
	size = videoSize(strings.Split(name, ".")[0] + "_" + strconv.Itoa(size) + "p." + strings.Split(name, ".")[1])
	if size == -1 {
		return
	}
	scaler(strings.Split(name, ".")[0] + "_" + strconv.Itoa(size) + "p." + strings.Split(name, ".")[1], size)
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
	height, err := strconv.Atoi(strings.Replace(strings.Split(out.String(), "x")[1], "\n", "", -1))
	if err != nil {
		log.Fatal(err)
	}
	if height >= 1080 {
		return 720
	} else if height >= 720  {
		return 480
	} else if height > 480  {
		return 360
	} else if height > 360  {
		return 240
	} else if height > 240  {
		return 144
	} else if height > 144  {
		return -1
	}
	return -1
}
