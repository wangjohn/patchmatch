package main

import (
  "image"
  "image/png"

  "github.com/wangjohn/patchmatch/seamcarving"
)

func main() {
  inputSource := ""
  testSeamCarving(inputSource)
}

func testSeamCarving(filename string) (error) {
  source, _, err := decodeImage(filename)
  if err != nil {
    return err
  }

  result, err := Resize(source, source.Bounds().Dy(), source.Bounds().Dx())
  if err != nil {
    return err
  }

  png.Encode(w, result)
}

func decodeImage(filename string) (image.Image, string, error) {
  f, err := os.Open(filename)
  if err != nil {
    return nil, "", err
  }
  defer f.Close()
  return image.Decode(bufio.NewReader(f))
}
