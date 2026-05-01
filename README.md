# Activity Heatmap

Add GPX Files in `/activities` folder

## Usage

### Locally

```
MAPLIBRE_STYLE="https://api.protomaps.com/styles/v5/light/en.json?key=MY_KEY" go run .
```

### Garmin

```sh
brew install --cask openmtp
```

Copy `.fit` files from `/GARMIN/ACTIVITY` to your `activities` folder

### Convert `.fit` to `.gpx`

```bash
brew install gpsbabel
```

```bash
for f in *.fit; do gpsbabel -i garmin_fit -f "$f" -o gpx -F "${f%.fit}.gpx"; done
```
