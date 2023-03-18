package main

import (
    "fmt"
    "io/ioutil"
    "math/rand"
    "os"
    "os/exec"
    "path/filepath"
    "strconv"
    "time"
)

const (
    defaultDuration = 0 // in seconds
)

func main() {
    var (
        imagesDir string
        duration  int
    )

    args := os.Args[1:]

    switch len(args) {
    case 1:
        imagesDir = args[0]
    case 2:
        imagesDir = args[0]
        duration, _ = strconv.Atoi(args[1])
    default:
        fmt.Println("Usage: setwallpaper <directory> [rotation_duration]")
        return
    }

    // check if feh is installed
    _, err := exec.LookPath("feh")
    if err != nil {
        fmt.Println("Warning: feh not found. Please install it to set wallpapers.")
        return
    }

    for {
        imageFiles, err := ioutil.ReadDir(imagesDir)
        if err != nil {
            fmt.Println("Error: ", err)
            return
        }

        // get random image file from directory
        rand.Seed(time.Now().UnixNano())
        randomImage := imageFiles[rand.Intn(len(imageFiles))].Name()

        // set the image file as wallpaper
        imageFilePath := filepath.Join(imagesDir, randomImage)
        cmd := exec.Command("feh", "--bg-scale", imageFilePath)
        if err := cmd.Run(); err != nil {
            fmt.Println("Error: ", err)
            return
        }

        if duration == 0 {
            // if no rotation duration specified, exit after setting wallpaper once
            break
        } else {
            time.Sleep(time.Duration(duration) * time.Second)
        }
    }
}

