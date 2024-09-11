package testing

import (
	"fmt"
	"testing"
)

type Package struct {
	width, height, length, mass int
}

func NewPackage(width, height, length, mass int) (*Package, error) {
	if width <= 0 || height <= 0 || length <= 0 || mass <= 0 {
		return nil, fmt.Errorf("invalid package")
	}
	return &Package{width, height, length, mass}, nil
}

// Sort returns the type of package based on its dimensions and mass
func Sort(width, height, length, mass int) string {
	p, err := NewPackage(width, height, length, mass)
	if err != nil {
		return "REJECTED"
	}
	if p.isBulky() && p.isHeavy() {
		return "REJECTED"
	}
	if p.isBulky() || p.isHeavy() {
		return "SPECIAL"
	}
	return "STANDARD"
}

func (p *Package) isBulky() bool {
	if p.width*p.height*p.length >= 1000000 {
		return true
	}
	if p.width >= 150 || p.height >= 150 || p.length >= 150 {
		return true
	}
	return false
}

func (p *Package) isHeavy() bool {
	return p.mass >= 20
}

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		height   int
		length   int
		mass     int
		expected string
	}{
		{"validPackage", 1, 1, 1, 1, "STANDARD"},
		{"Volume is too big", 150, 150, 150, 1, "SPECIAL"},
		{"Side is too long", 1, 1, 150, 1, "SPECIAL"},
		{"isBulky and Heavy", 150, 150, 20, 50, "REJECTED"},
		{"invalidPackage 0x1x1x1", 0, 1, 1, 1, "REJECTED"},
		{"invalidPackage 1x0x1x1", 1, 0, 1, 1, "REJECTED"},
		{"invalidPackage 1x1x0x1", 1, 1, 0, 1, "REJECTED"},
		{"invalidPackage 1x1x1x0", 1, 1, 1, 0, "REJECTED"},
	}
	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			if Sort(tests[i].width, tests[i].height, tests[i].length, tests[i].mass) != tests[i].expected {
				t.Errorf("expected %v, got %v", tests[i].expected, Sort(tests[i].width, tests[i].height, tests[i].length, tests[i].mass))
			}
		})
	}
}

func TestNewPackage(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		height   int
		length   int
		mass     int
		expected *Package
	}{
		{"validPackage 1x1x1x1", 1, 1, 1, 1, &Package{1, 1, 1, 1}},
		{"invalidPackage 0x1x1x1", 0, 1, 1, 1, nil},
		{"invalidPackage 1x0x1x1", 1, 0, 1, 1, nil},
		{"invalidPackage 1x1x0x1", 1, 1, 0, 1, nil},
		{"invalidPackage 1x1x1x0", 1, 1, 1, 0, nil},
	}
	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			p, err := NewPackage(tests[i].width, tests[i].height, tests[i].length, tests[i].mass)
			if tests[i].expected == nil {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected nil, got %v", err)
				}
				if p.width != tests[i].expected.width || p.height != tests[i].expected.height || p.length != tests[i].expected.length || p.mass != tests[i].expected.mass {
					t.Errorf("expected %v, got %v", tests[i].expected, p)
				}
			}
		})
	}
}

func TestIsBulky(t *testing.T) {
	tests := []struct {
		name     string
		p        *Package
		expected bool
	}{
		{"notBulky", &Package{1, 1, 1, 1}, false},
		{"Volume is too big", &Package{150, 150, 150, 1}, true},
		{"Size is too long", &Package{1, 1, 150, 1}, true},
	}
	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			if tests[i].p.isBulky() != tests[i].expected {

				t.Errorf("expected %v, got %v", tests[i].expected, tests[i].p.isBulky())
			}
		})
	}
}

func TestIsHeavy(t *testing.T) {
	tests := []struct {
		name     string
		p        *Package
		expected bool
	}{
		{"notHeavy", &Package{1, 1, 1, 1}, false},
		{"isHeavy", &Package{1, 1, 1, 20}, true},
	}
	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			if tests[i].p.isHeavy() != tests[i].expected {
				t.Errorf("expected %v, got %v", tests[i].expected, tests[i].p.isBulky())
			}
		})
	}
}
