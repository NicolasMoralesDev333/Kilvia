package routing

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	Latitude  float64
	Longitude float64
}

type Route struct {
	DistanceKM  float64
	Duration    time.Duration
	LegsKM      []float64
	Coordinates [][]float64
	Source      string
}

type Router struct {
	baseURL string
	client  *http.Client
}

func New(baseURL string) *Router {
	if strings.TrimSpace(baseURL) == "" {
		baseURL = "https://router.project-osrm.org"
	}
	return &Router{
		baseURL: strings.TrimRight(baseURL, "/"),
		client:  &http.Client{Timeout: 3 * time.Second},
	}
}

func (r *Router) Route(ctx context.Context, points []Point) Route {
	if len(points) < 2 {
		return Route{}
	}
	if result, err := r.osrm(ctx, points); err == nil {
		return result
	}
	return fallback(points)
}

type osrmResponse struct {
	Code   string `json:"code"`
	Routes []struct {
		Distance float64 `json:"distance"`
		Duration float64 `json:"duration"`
		Geometry struct {
			Coordinates [][]float64 `json:"coordinates"`
		} `json:"geometry"`
		Legs []struct {
			Distance float64 `json:"distance"`
		} `json:"legs"`
	} `json:"routes"`
}

func (r *Router) osrm(ctx context.Context, points []Point) (Route, error) {
	coordinates := make([]string, len(points))
	for i, point := range points {
		coordinates[i] = strconv.FormatFloat(point.Longitude, 'f', 6, 64) + "," +
			strconv.FormatFloat(point.Latitude, 'f', 6, 64)
	}
	endpoint := r.baseURL + "/route/v1/driving/" + strings.Join(coordinates, ";")
	query := url.Values{
		"overview":     {"full"},
		"geometries":   {"geojson"},
		"steps":        {"false"},
		"alternatives": {"false"},
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint+"?"+query.Encode(), nil)
	if err != nil {
		return Route{}, err
	}
	response, err := r.client.Do(request)
	if err != nil {
		return Route{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return Route{}, fmt.Errorf("OSRM returned %s", response.Status)
	}
	var decoded osrmResponse
	if err := json.NewDecoder(response.Body).Decode(&decoded); err != nil {
		return Route{}, err
	}
	if decoded.Code != "Ok" || len(decoded.Routes) == 0 {
		return Route{}, fmt.Errorf("OSRM route unavailable")
	}
	selected := decoded.Routes[0]
	legs := make([]float64, len(selected.Legs))
	for i, leg := range selected.Legs {
		legs[i] = leg.Distance / 1000
	}
	return Route{
		DistanceKM:  selected.Distance / 1000,
		Duration:    time.Duration(selected.Duration * float64(time.Second)),
		LegsKM:      legs,
		Coordinates: selected.Geometry.Coordinates,
		Source:      "OSRM",
	}, nil
}

func fallback(points []Point) Route {
	result := Route{
		LegsKM:      make([]float64, 0, len(points)-1),
		Coordinates: make([][]float64, len(points)),
		Source:      "FALLBACK",
	}
	for i, point := range points {
		result.Coordinates[i] = []float64{point.Longitude, point.Latitude}
	}
	for i := 1; i < len(points); i++ {
		km := haversine(points[i-1], points[i]) * 1.22
		result.LegsKM = append(result.LegsKM, km)
		result.DistanceKM += km
	}
	result.Duration = time.Duration(result.DistanceKM / 75 * float64(time.Hour))
	return result
}

func haversine(a, b Point) float64 {
	const earthRadiusKM = 6371.0088
	radians := math.Pi / 180
	lat1, lat2 := a.Latitude*radians, b.Latitude*radians
	dLat := (b.Latitude - a.Latitude) * radians
	dLng := (b.Longitude - a.Longitude) * radians
	value := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLng/2)*math.Sin(dLng/2)
	return 2 * earthRadiusKM * math.Asin(math.Sqrt(value))
}
