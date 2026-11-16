package geocoder

import "testing"

func TestResolveIgnoresCaseAndAccents(t *testing.T) {
	for _, city := range []string{"CORDOBA", "Córdoba", " córdoba "} {
		if _, ok := Resolve(city); !ok {
			t.Errorf("Resolve(%q) did not find the city", city)
		}
	}
}
