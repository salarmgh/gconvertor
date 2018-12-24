package main

import (
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
	"os/exec"
	"bytes"
	"strings"
	"strconv"
	"path/filepath"
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
	if _, err := os.Stat("/data/" + body.Name); os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	scaler("/data/" + body.Name)
}

func getSize(name string) int {
	height := -1
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "csv=s=x:p=0", name)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		height, err = strconv.Atoi(strings.Replace(strings.Split(out.String(), "x")[1], "\n", "", -1))
		if err != nil {
			log.Fatal(err)
		}
	}
	switch height {
	case 1080:
		return 0
	case 720:
		return 1
	case 480:
		return 2
	case 360:
		return 3
	case 240:
		return 4
	case 144:
		return -1
	}
	return -1
}

func scaler(srcName string) {
	sizes := []string{"1920x1080", "960x720", "640x480", "480x360", "320x240", "256x144"}


	srcSizeIndex := getSize(srcName)
	if srcSizeIndex == -1 {
		return
	}

	srcSize := sizes[srcSizeIndex]
	srcHeight := strings.Split(srcSize, "x")[1]

	dstSize := sizes[srcSizeIndex + 1]
	dstHeight := strings.Split(dstSize, "x")[1]

	contain := false
	for _, size := range sizes {
		if strings.Contains(srcName, strings.Split(size, "x")[1]) {
			contain = true
			break
		}
	}

	if !contain {
    	nameHeight := strings.Replace(srcName, filepath.Ext(srcName), "", -1) + "_" + strings.Split(srcSize, "x")[1] + "p" + filepath.Ext(srcName)
    	os.Rename(srcName, nameHeight)
    	srcName = nameHeight
	}

	dstName := strings.Replace(srcName, srcHeight, dstHeight, -1)

	
	if _, err := os.Stat(dstName); os.IsNotExist(err) {
		cmd := exec.Command("ffmpeg", "-i", srcName, "-s", dstSize, dstName)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
    }
	scaler(dstName)
}
