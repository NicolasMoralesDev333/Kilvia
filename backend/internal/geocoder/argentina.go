package geocoder

import (
	"strings"
	"unicode"
)

type Location struct {
	City      string
	Region    string
	Latitude  float64
	Longitude float64
}

var cities = []Location{
	{"Buenos Aires", "CABA", -34.6037, -58.3816},
	{"La Plata", "Buenos Aires", -34.9215, -57.9545},
	{"Mar del Plata", "Buenos Aires", -38.0055, -57.5426},
	{"Bahía Blanca", "Buenos Aires", -38.7196, -62.2724},
	{"Rosario", "Santa Fe", -32.9442, -60.6505},
	{"Santa Fe", "Santa Fe", -31.6333, -60.7000},
	{"Córdoba", "Córdoba", -31.4201, -64.1888},
	{"Mendoza", "Mendoza", -32.8895, -68.8458},
	{"San Luis", "San Luis", -33.3017, -66.3378},
	{"San Juan", "San Juan", -31.5375, -68.5364},
	{"San Miguel de Tucumán", "Tucumán", -26.8241, -65.2226},
	{"Salta", "Salta", -24.7821, -65.4232},
	{"Jujuy", "Jujuy", -24.1858, -65.2995},
	{"Posadas", "Misiones", -27.3621, -55.9009},
	{"Corrientes", "Corrientes", -27.4692, -58.8306},
	{"Resistencia", "Chaco", -27.4514, -58.9867},
	{"Paraná", "Entre Ríos", -31.7413, -60.5115},
	{"Concordia", "Entre Ríos", -31.3929, -58.0209},
	{"Neuquén", "Neuquén", -38.9516, -68.0591},
	{"San Carlos de Bariloche", "Río Negro", -41.1335, -71.3103},
	{"Trelew", "Chubut", -43.2489, -65.3051},
	{"Comodoro Rivadavia", "Chubut", -45.8641, -67.4966},
	{"Río Gallegos", "Santa Cruz", -51.6230, -69.2168},
	{"Ushuaia", "Tierra del Fuego", -54.8019, -68.3030},
	{"San Salvador de Jujuy", "Jujuy", -24.1858, -65.2995},
	{"Santiago del Estero", "Santiago del Estero", -27.7834, -64.2642},
	{"Formosa", "Formosa", -26.1849, -58.1731},
	{"Catamarca", "Catamarca", -28.4696, -65.7795},
	{"La Rioja", "La Rioja", -29.4131, -66.8558},
	{"Rawson", "Chubut", -43.3002, -65.1023},
	{"Viedma", "Río Negro", -40.8135, -62.9967},
	{"Santa Rosa", "La Pampa", -36.6203, -64.2900},
	{"San Fernando del Valle de Catamarca", "Catamarca", -28.4696, -65.7795},
}

func Resolve(city string) (Location, bool) {
	key := normalize(city)
	for _, location := range cities {
		if normalize(location.City) == key {
			return location, true
		}
	}
	return Location{}, false
}

func normalize(value string) string {
	var out strings.Builder
	for _, r := range strings.ToLower(strings.TrimSpace(value)) {
		switch r {
		case 'á':
			r = 'a'
		case 'é':
			r = 'e'
		case 'í':
			r = 'i'
		case 'ó':
			r = 'o'
		case 'ú', 'ü':
			r = 'u'
		case 'ñ':
			r = 'n'
		}
		if unicode.IsSpace(r) || unicode.Is(unicode.Mn, r) {
			continue
		}
		out.WriteRune(r)
	}
	return out.String()
}
