package utils

import (
  "testing"
  "giftwist/utils"
)

const DIR = "../test-inputs/"

func TestCountFrames(t *testing.T) {
  var count = utils.CountFrames(DIR)
  if count != 4 {
    t.Error("Expected 4, got ", count)
  }
}

func TestLoadFrame(t *testing.T) {
  frame := *utils.LoadFrames(DIR)[0]
  x := frame.Bounds().Max.X
  if x != 1455 {
    t.Error("Expected frame size of 1455, got", x)
  }
}

func TestLoadWhiteFrame(t *testing.T) {
  frame := *utils.LoadFrames(DIR)[0]
  r,g,b,_ := frame.At(0,0).RGBA()
  if (r != 65535 || g != 65535 || b != 65535) {
    t.Error("Expected frame color to be white, got", g)
  }
}


func TestLoadBlackFrame(t *testing.T) {
  frame := *utils.LoadFrames(DIR)[1]
  r,g,b,_ := frame.At(0,0).RGBA()
  if (r != 0 || g != 0 || b != 0) {
    t.Error("Expected frame color to be black, got", g)
  }
}
