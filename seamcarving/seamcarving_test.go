package seamcarving

import (
  "testing"
  "image"
)

func TestInitializatingEnergies(t *testing.T) {
  fixtures := []struct {
    Image image.Image
  }{
    {image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{20, 20}})},
  }

  for _, fixture := range fixtures {
    energies := initializeEnergies(fixture.Image, Energy1)
    if len(energies) != fixture.Image.Bounds().Dy() {
      t.Errorf("Energy rectangle bounds are invalid: expected %v, but got %v",
        fixture.Image.Bounds().Dy(), len(energies))
    }

    if len(energies[0]) != fixture.Image.Bounds().Dx() {
      t.Errorf("Energy rectangle bounds are invalid: expected %v, but got %v",
        fixture.Image.Bounds().Dx(), len(energies[0]))
    }

    for i, _ := range energies {
      for j, _ := range energies[i] {
        if energies[i][j] != 0.0 {
          t.Errorf("Expected energy %v, but got %v", 0.0, energies[i][j])
        }
      }
    }
  }
}

func TestComputeSeams(t *testing.T) {
  fixtures := []struct {
    Energies [][]float64
    NumSeams int
    ExpectedSeams []Seam
  }{
  }

  for _, fixture := range fixtures {
    resultSeams := computeSeams(fixture.Energies, fixture.NumSeams)
    for _, expectedSeam := range fixture.ExpectedSeams {
      var hasExpectedSeam bool
      for _, resultSeam := range resultSeams {
        if isMatchingSeam(expectedSeam, resultSeam) {
          hasExpectedSeam = true
          break
        }
      }

      if !hasExpectedSeam {
        t.Errorf("Did not find the expected seam: %v", expectedSeam)
      }
    }
  }
}

func isMatchingSeam(seam1, seam2 Seam) (bool) {
  if len(seam1.Points) != len(seam2.Points) {
    return false
  }

  for i, s1point := range seam1.Points {
    s2point := seam2.Points[i]
    if s1point.X != s2point.X || s1point.Y != s2point.Y {
      return false
    }
  }

  return true
}
