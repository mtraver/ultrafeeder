package ultrafeeder

// ACASResolutionAdvisory represents an ACAS (Airborne Collision Avoidance System)
// Resolution Advisory. This is experimental and subject to change.
// See format here: https://github.com/wiedehopf/readsb/blob/caf2c5a6dccf1e079d0ae63d2b7dd45d568de265/json_out.c#L262
type ACASResolutionAdvisory struct {
	// UTC is the time of the advisory in UTC.
	UTC string `json:"utc"`
	// UnixTimestamp is the time of the advisory as a Unix timestamp.
	UnixTimestamp float64 `json:"unix_timestamp"`

	// Debug indicates the RA failed validation (only present when debug mode is enabled).
	Debug *bool `json:"debug"`
	// DFType is the downlink format type of the message.
	DFType *int `json:"df_type"`
	// FullBytes is the full hex-encoded message bytes (only present for valid RA messages).
	FullBytes *string `json:"full_bytes"`

	// Bytes is the 7 hex-encoded bytes of the ACAS RA message.
	Bytes string `json:"bytes"`

	// ARA is the Active Resolution Advisory bits (bits 9-15), as a binary string.
	ARA string `json:"ARA"`
	// RAT is the Resolution Advisory Terminated bit, as a binary string ("0" or "1").
	RAT string `json:"RAT"`
	// MTE is the Multiple Threat Encounter bit, as a binary string ("0" or "1").
	MTE string `json:"MTE"`
	// RAC is the Resolution Advisory Complement bits (bits 23-26), as a binary string.
	RAC string `json:"RAC"`

	// AdvisoryComplement is the human-readable interpretation of the RAC bits,
	// e.g. "Do not pass below; Do not turn left".
	AdvisoryComplement string `json:"advisory_complement"`
	// Advisory is the human-readable resolution advisory,
	// e.g. "Climb", "Descend", "Clear of Conflict", "Level Off".
	Advisory string `json:"advisory"`

	// TTI is the Threat Type Indicator bits (bits 29-30), as a binary string.
	TTI string `json:"TTI"`
	// ThreatIDHex is the ICAO address of the threat aircraft (only present when TTI == "01").
	ThreatIDHex *string `json:"threat_id_hex"`
}
