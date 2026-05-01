package ultrafeeder

// LastPosition holds the last valid position when current data is stale.
type LastPosition struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	// NIC is the Navigation Integrity Category (2.2.3.2.7.2.6).
	NIC int `json:"nic"`
	// RadiusOfContainment is a measure of position integrity, in meters, derived from NIC and supplementary bits (2.2.3.2.7.2.6, Table 2-69).
	RadiusOfContainment int `json:"rc"`
	// Seen is how long ago (in seconds before "now") this was received.
	Seen float64 `json:"seen_pos"`
}
