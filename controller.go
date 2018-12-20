package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"net/http"
	"bytes"
	"encoding/json"
	"strings"
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
	scaler(conf.Path + "/" + body.Name, "960x720")
}

func scaler(name string, size string) {
    log.Println("ffmpeg, -i, " + name + ", -s, " + size + ", " + strings.Split(name, ".")[0] + "_720p." + strings.Split(name, ".")[1])
	cmd := exec.Command("ffmpeg", "-i", name, "-s", "960x720", strings.Split(name, ".")[0] + "_720p." + strings.Split(name, ".")[1])
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
