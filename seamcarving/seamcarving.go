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
  energies := initializeEnergies(rec, Energy1) // TODO: don't just use Energy1 function
  return nil, nil
}

func initializeEnergies(rec image.Rectangle, funcType EnergyFunction) ([][]float64) {
  height := rec.Max.X - rec.Min.X
  width := rec.Max.Y - rec.Min.Y

  energies := float64Matrix(height, width)
  for i := rec.Min.X; i < rec.Max.X; i++ {
    for j := rec.Min.Y; i < rec.Max.Y; j++ {
      xIndex := i - rec.Min.X
      yIndex := j - rec.Min.Y
      energies[xIndex][yIndex] = energyFunction(source, i, j, funcType)
    }
  }

  return energies
}

func float64Matrix(height, width int) ([][]float64) {
  matrix := make([][]float64, height)
  for i, _ := range matrix {
    matrix[i] = make([]float64, width)
  }

  return matrix
}

func intMatrix(height, width int) ([][]int) {
  matrix := make([][]int, height)
  for i, _ := range matrix {
    matrix[i] = make([]int, width)
  }

  return matrix
}

func energyFunction(img image.Image, i, j int, funcType EnergyFunction) (float64) {
  return 1.0
}

func shouldReadjust(i, j int, candidate float64, matrix [][]float64, adjusted bool) (bool) {
  height := len(matrix)
  width := len(matrix[0])

  return inMatrix(i, j, width, height) && (!adjusted || matrix[i][j] < candidate)
}

func computeSeams(energies [][]float64, numSeams int) ([]Seam) {
  height := len(energies)
  width := len(energies[0])

  seamTable := float64Matrix(height, width)
  parentTable := intMatrix(height, width)

  for i := 0; i < height; i++ {
    for j := 0; j < width; j++ {
      var candidate float64
      adjusted := false
      parent := 0

      if shouldReadjust(i-1, j-1, candidate, seamTable, adjusted) {
        parent = -1
        adjusted = true
        candidate = seamTable[i-1][j-1]
      }
      if shouldReadjust(i-1, j, candidate, seamTable, adjusted) {
        parent = 0
        adjusted = true
        candidate = seamTable[i-1][j]
      }
      if shouldReadjust(i-1, j+1, candidate, seamTable, adjusted) {
        parent = 1
        adjusted = true
        candidate = seamTable[i-1][j+1]
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
