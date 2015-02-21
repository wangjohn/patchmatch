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
