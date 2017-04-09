package utils

import (
  "image"
  "image/png"
  "image/gif"
  "os"
  "sort"
  "strings"
  "fmt"
)

func CountFrames(inputDir string) int {
  dir, err := os.Open(inputDir)
  if err != nil {
    panic(err)
  }

  fileNames, err := dir.Readdirnames(0)
  if err != nil {
    panic(err)
  }

  var count = 0
  for _, fileName := range fileNames {
    if strings.HasSuffix(fileName, ".png") { count ++ }
  }
  return count
}

func LoadFrames(inputDir string) []*image.Image {
  dir, err := os.Open(inputDir)
  fileNames, err := dir.Readdirnames(0)
  if err != nil {
    panic(err)
  }

  frames := []*image.Image{}
  sort.Strings(fileNames)
  for _, fileName := range fileNames {
    if !strings.HasSuffix(fileName, ".png") { continue }

    file, err := os.Open(inputDir + fileName)
    if err != nil {
      panic(err)
    }
    frame, err := png.Decode(file)
    frames = append(frames, &frame)
  }

  return frames
}

func SaveFrame(dir string, output image.Image, index int) {
  f, err := os.Create(fmt.Sprintf("%v/%04d.gif", dir, index))
  if err != nil { panic(err) }

  defer f.Close()

  options := gif.Options{255, nil, nil}
  err = gif.Encode(f, output, &options)
  if err != nil { panic(err) }
}
