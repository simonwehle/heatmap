package style

import (
	"os"
)

func GetMapStyle() string {
	mapStyle := os.Getenv("MAPLIBRE_STYLE")
	if mapStyle == "" {
		mapStyle = "https://demotiles.maplibre.org/style.json"
	}
	return mapStyle
}
