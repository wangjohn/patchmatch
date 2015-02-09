package patchmatch

import (
  "image"
)

type NNFOptions struct {
  PatchSize int
}

func NearestNeighborField(source, target image.Image, options NNFOptions) (error) {
  
}

func nearestNeighborSearch(source, target image.Image) (error) {
  sourceBounds := source.Bounds()
  targetBounds := target.Bounds()

  for sy := sourceBounds.Min.Y; sy < sourceBounds.Max.Y; sy++ {
    for sx := sourceBounds.Min.X; sy < sourceBounds.Max.X, sx++ {

    }
  }
}
