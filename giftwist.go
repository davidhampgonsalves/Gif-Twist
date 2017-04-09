package main

import (
  "fmt"
  "giftwist/utils"
  "image/color"
  "image/draw"
  "image"
  "math"
  "sync"
)

// take all frames from input/
// divide image
const OUTPUT_DIR = "./out"
const FRAME_COUNT = 360 / 2
func main() {
  var inputDir = "./inputs/"
	fmt.Printf("3,2,1 lets jam\n")

  inputFrames := utils.LoadFrames(inputDir)
  fmt.Printf("Loaded %v frames\n", len(inputFrames))

  frameAngle := float64(360 / len(inputFrames))
  fmt.Printf("Using frame angle of %v, based on %v input frames.\n", frameAngle, len(inputFrames))
  offset := 0.0

  var wg sync.WaitGroup
  for i := 0 ; i < FRAME_COUNT ; i++ {
    offset = float64(i * (360 / FRAME_COUNT))
    wg.Add(1)

    go generateFrame(inputFrames, offset, frameAngle, i, &wg)
  }

  fmt.Printf("%v frames enqueued\n", FRAME_COUNT)
  wg.Wait()
  fmt.Printf("%v frames created, you are welcome.\n", FRAME_COUNT)
}

func generateFrame(inputFrames []*image.Image, offset float64, frameAngle float64, i int, wg *sync.WaitGroup) {
  var first image.Image = *(inputFrames[0])
  bounds := first.Bounds()
  size := bounds.Size()
  center := image.Point{size.X/2, size.Y/2}
  dst := image.NewRGBA(bounds)

  for i, inputFrame := range inputFrames {
    mask := NewAngleMask(center, bounds, offset, frameAngle, i)
    draw.DrawMask(dst, bounds, *inputFrame, image.ZP, mask, bounds.Min, draw.Over)
  }

  utils.SaveFrame(OUTPUT_DIR, dst, i)
  defer wg.Done()
  fmt.Printf("Saved frame %v / %v\n", i, FRAME_COUNT)
}

func NewAngleMask(center image.Point, bounds image.Rectangle, offset float64, frameAngle float64, index int) *AngleMask {
  start := offset + (float64(index) * frameAngle)
  end := start + frameAngle

  if end > 360 { end -= 360 }
  if start > 360 { start -= 360 }

  return &AngleMask{center, bounds, index, offset, frameAngle, start, end}
}

type AngleMask struct {
  center image.Point
  bounds image.Rectangle
  index int
  offset float64
  frameAngle float64
  start float64
  end float64
}

func (c *AngleMask) ColorModel() color.Model {
  return color.AlphaModel
}

func (c *AngleMask) Bounds() image.Rectangle {
  return c.bounds
}

func (c *AngleMask) At(x, y int) color.Color {
  pX := c.center.X
  pY := float64(c.center.Y) - math.Sqrt(math.Abs(float64(x - c.center.X)) * math.Abs(float64(x - c.center.X)) + math.Abs(float64(y - c.center.Y)) * math.Abs(float64(y - c.center.Y)))

  angle := math.Atan2(float64(y) - pY, float64(x - pX)) * (360 / math.Pi)

  if (angle >= c.start && angle <= c.end) || (c.end < c.start && (angle >= c.start || angle <= c.end)) {
    return color.Alpha{255}
  }
  return color.Alpha{0}
}

