package seamcarving

import (
  "image"

  "github.com/wangjohn/quickselect"
)

type EnergyFunction int

const (
  Energy1 EnergyFunction = iota
  Energy2 EnergyFunction = iota
)

type Seam struct {
  Points []image.Point
}

func Resize(source image.Image, targetHeight, targetWidth int) (image.Image, error) {
  rec := source.Bounds()
  energies := initializeEnergies(source, Energy1) // TODO: don't just use Energy1 function
}

func initializeEnergies(img image.Image, funcType EnergyFunction) ([][]float64) {
  var energies [rec.Max.X - rec.Min.X][rec.Max.Y - rec.Min.X]float64

  for i := rec.Min.X; i < rec.Max.X; i++ {
    for j := rec.Min.Y; i < rec.Max.Y; j++ {
      xIndex := i - rec.Min.X
      yIndex := y - rec.Min.Y
      energies[xIndex][yIndex] = energyFunction(source, i, j, funcType)
    }
  }

  return energies
}

func energyFunction(img image.Image, i, j int, funcType EnergyFunction) (float64) {
  return 1.0
}

func computeSeams(energies [][]float64, numSeams int) ([]Seam) {
  height := len(energies)
  width := len(energies[0])

  var seamTable [width][height]float64
  var parentTable [width][height]int

  for i := 0; i < height; i++ {
    for j := 0; j < width; j++ {
      candidate := 0
      parent := 0

      if inMatrix(i-1, j-1, width, height) {
        parent = -1
        candidate = math.Min(candidate, seamTable[i-1][j-1])
      }
      if inMatrix(i-1, j, width, height) {
        parent = 0
        candidate = math.Min(candidate, seamTable[i-1][j])
      }
      if inMatrix(i-1, j+1, width, height) {
        parent = 1
        candidate = math.Min(candidate, seamTable[i-1][j+1])
      }

      parentTable[i][j] = parent
      seamTable[i][j] = candidate + energies[i][j]
    }
  }

  lastRowClone := copy(seamTable[height-1])
  quickselect.QuickSelect(quickselect.Float64Slice(lastRowClone), numSeams)
  thresholdEnergy := lastRowClone[numSeams-1]

  computedSeams := make([]Seam, numSeams)
  seamNum := 0
  for j := 0; j < width; j++ {
    if seamTable[height-1][j] <= thresholdEnergy {
      points := make([]image.Point, height)
      currentY := j

      for i := height-1; i >= 0; i-- {
        points[i] = image.Point{i, currentY}
        currentY = currentY + parentTable[i][j]
      }

      computedSeams[seamNum] = Seam{points}
      seamNum++
    }
  }

  return computedSeams
}

func inMatrix(i, j, width, height int) (bool) {
  return 0 <= i && i < height && 0 <= j && j < width
}
