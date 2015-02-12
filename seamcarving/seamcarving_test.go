package seamcarving

import (
  "testing"
  "image"
)

func TestInitializatingEnergies(t *testing.T) {
  fixtures := []struct {
    Bounds image.Rectangle
  }{
    {image.Rectangle{image.Point{0, 0}, image.Point{20, 20}}},
  }

  for _, fixture := range fixtures {
    energies := initializeEnergies(fixture.Bounds, Energy1)
  }
}
