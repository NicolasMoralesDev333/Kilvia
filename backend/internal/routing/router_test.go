package routing

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouteFallsBackWhenOSRMIsUnavailable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "offline", http.StatusServiceUnavailable)
	}))
	defer server.Close()

	router := New(server.URL)
	points := []Point{
		{Latitude: -34.6037, Longitude: -58.3816},
		{Latitude: -33.3017, Longitude: -66.3378},
		{Latitude: -32.8895, Longitude: -68.8458},
	}
	first := router.Route(context.Background(), points)
	second := router.Route(context.Background(), points)
	if first.DistanceKM <= 0 || first.DistanceKM != second.DistanceKM {
		t.Fatalf("fallback distance should be positive and deterministic: %.2f, %.2f", first.DistanceKM, second.DistanceKM)
	}
	if len(first.LegsKM) != 2 || len(first.Coordinates) != len(points) {
		t.Fatalf("expected one route leg per waypoint pair, got %+v", first)
	}
}
