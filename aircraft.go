package ultrafeeder

// AircraftJSON represents the top-level aircraft.json response.
type AircraftJSON struct {
	Now float64 `json:"now"`
	// MessageCount is the total number of Mode S messages processed since readsb started.
	MessageCount int64      `json:"messages"`
	Aircraft     []Aircraft `json:"aircraft"`
}

// Aircraft represents a single aircraft entry in the aircraft.json response.
// Section references like (2.2.xyz) refer to DO-260B.
type Aircraft struct {
	// Identity.
	// ICAOAddr is the 24-bit ICAO identifier of the aircraft, as 6 hex digits. The identifier may start with '~', this means that the address is a non-ICAO address (e.g. from TIS-B).
	ICAOAddr string `json:"hex"`
	// Type is the type of underlying messages / best source of current data for this position / aircraft.
	Type string `json:"type"`
	// Callsign is the callsign, flight name, or aircraft registration as 8 chars (2.2.8.2.6).
	Callsign string `json:"flight"`
	// Squawk is the Mode A code (squawk), encoded as 4 octal digits.
	Squawk string `json:"squawk"`
	// Category is the emitter category to identify particular aircraft or vehicle classes (values A0 - D7) (2.2.3.2.5.2).
	Category string `json:"category"`

	// Registration.
	// Registration is the aircraft registration pulled from the database.
	Registration string `json:"r"`
	// AircraftType is the aircraft type pulled from database.
	AircraftType string `json:"t"`
	// LongAircraftType is the long aircraft type pulled from the database (this requires --db-file-lt).
	LongAircraftType string `json:"desc"`
	OwnerOperator    string `json:"ownOp"`
	Year             string `json:"year"`

	// DB fields (requires --db-file).
	// DBFlags is a bitfield encoding certain database flags.
	//   military = dbFlags & 1
	//   interesting = dbFlags & 2
	//   PIA = dbFlags & 4
	//   LADD = dbFlags & 8
	DBFlags int `json:"dbFlags"`

	// Position.
	// Latitude is the aircraft's latitude in decimal degrees.
	Latitude *float64 `json:"lat"`
	// Latitude is the aircraft's longitude in decimal degrees.
	Longitude *float64 `json:"lon"`
	// PositionLastSeen is how long ago (in seconds before "now") the position was last updated.
	PositionLastSeen *float64 `json:"seen_pos"`

	// Altitude.
	// BarometricAltitude is the aircraft barometric altitude in feet, or "ground".
	BarometricAltitude *BarometricAltitude `json:"alt_baro"`
	// GeometricAltitude is the geometric (GNSS / INS) altitude in feet referenced to the WGS84 ellipsoid.
	GeometricAltitude *float64 `json:"alt_geom"`

	// Speed.
	// GroundSpeed is the ground speed in knots.
	GroundSpeed *float64 `json:"gs"`
	// IndicatedAirspeed is the indicated air speed in knots.
	IndicatedAirspeed *float64 `json:"ias"`
	// TrueAirspeed is the true air speed in knots.
	TrueAirspeed *float64 `json:"tas"`
	// Mach number is the Mach number.
	MachNumber *float64 `json:"mach"`

	// Heading and track.
	// Track is the true track over ground in degrees (0-359).
	Track *float64 `json:"track"`
	// TrackRate is the rate of change of track, degrees/second.
	TrackRate *float64 `json:"track_rate"`
	// Roll is the roll in degrees. Negative is left roll.
	Roll *float64 `json:"roll"`
	// MagHeading is the heading in degrees clockwise from magnetic north.
	MagHeading *float64 `json:"mag_heading"`
	// TrueHeading is the heading in degrees clockwise from true north (usually only transmitted on ground, in the air usually derived from the magnetic heading using magnetic model WMM2020).
	TrueHeading *float64 `json:"true_heading"`

	// Vertical Rate.
	// BaroRate is the rate of change of barometric altitude in feet/minute.
	BaroRate *float64 `json:"baro_rate"`
	// GeomRate is the rate of change of geometric (GNSS / INS) altitude in feet/minute.
	GeomRate *float64 `json:"geom_rate"`

	// Navigation.
	// NavQNH is the altimeter setting (QFE or QNH/QNE) in hPa.
	NavQNH *float64 `json:"nav_qnh"`
	// NavAltitudeMCP is the selected altitude from the Mode Control Panel / Flight Control Unit (MCP/FCU) or equivalent equipment.
	NavAltitudeMCP *int `json:"nav_altitude_mcp"`
	// NavAltitudeFMS is the selected altitude from the Flight Management System (FMS) (2.2.3.2.7.1.3.3).
	NavAltitudeFMS *int `json:"nav_altitude_fms"`
	// NavHeading is the selected heading (True or Magnetic is not defined in DO-260B, mostly Magnetic as that is the de facto standard) (2.2.3.2.7.1.3.7).
	NavHeading *float64 `json:"nav_heading"`
	// NavModes is the set of engaged automation modes: 'autopilot', 'vnav', 'althold', 'approach', 'lnav', 'tcas'.
	NavModes []string `json:"nav_modes"`

	// Integrity and accuracy.
	// NIC is the Navigation Integrity Category (2.2.3.2.7.2.6).
	NIC *int `json:"nic"`
	// RadiusOfContainment is a measure of position integrity, in meters, derived from NIC and supplementary bits (2.2.3.2.7.2.6, Table 2-69).
	RadiusOfContainment *int `json:"rc"`
	// Version is the ADS-B Version Number 0, 1, 2 (3-7 are reserved) (2.2.3.2.7.5).
	Version *int `json:"version"`
	// NICBaro is the Navigation Integrity Category for Barometric Altitude (2.2.5.1.35).
	NICBaro *int `json:"nic_baro"`
	// NACP is the Navigation Accuracy for Position (2.2.5.1.35).
	NACP *int `json:"nac_p"`
	// NACV is the Navigation Accuracy for Velocity (2.2.5.1.19).
	NACV *int `json:"nac_v"`
	// SIL is the Source Integity Level (2.2.5.1.40).
	SIL *int `json:"sil"`
	// SILType is the interpretation of SIL: "unknown", "perhour", "persample".
	SILType string `json:"sil_type"`
	// GVA is the Geometric Vertical Accuracy (2.2.3.2.7.2.8).
	GVA *int `json:"gva"`
	// SDA is the System Design Assurance (2.2.3.2.7.2.4.6).
	SDA *int `json:"sda"`

	// Receiver location accuracy (included when the readsb flag --json-location-accuracy is set, which in the ultrafeeder container is set via the env var READSB_RX_LOCATION_ACCURACY).
	// ReceiverDistance is the distance from the receiver to the aircraft in nautical miles.
	ReceiverDistance *float64 `json:"r_dst"`
	// ReceiverDirection is the direction from the receiver to the aircraft in degrees.
	ReceiverDirection *float64 `json:"r_dir"`

	// Status.
	// Emergency is the ADS-B emergency/priority status, a superset of the 7x00 squawks (2.2.3.2.7.8.1.1) (none, general, lifeguard, minfuel, nordo, unlawful, downed, reserved).
	Emergency string `json:"emergency"`
	// Alert is the flight status alert bit (2.2.3.2.3.2).
	Alert *int `json:"alert"`
	// SPI is the flight status special position identification bit (2.2.3.2.3.2).
	SPI *int `json:"spi"`

	// Signal.
	// RSSI is the recent average RSSI in dbFS; this will always be negative.
	RSSI *float64 `json:"rssi"`

	// Weather (derived).
	// WindDirection is the wind direction calculated from ground track, true heading, true airspeed, and ground speed.
	WindDirection *float64 `json:"wd"`
	// WindSpeed is the wind speed calculated from ground track, true heading, true airspeed, and ground speed.
	WindSpeed *float64 `json:"ws"`
	// OutsideAirTemp is the outer/static air temperature in degrees C calculated from mach number and true airspeed (typically somewhat inaccurate at lower altitudes / mach numbers below 0.5, calculation is inhibited for mach < 0.395).
	OutsideAirTemp *float64 `json:"oat"`
	// TotalAirTemp is the total air temperature in degrees C calculated from mach number and true airspeed (typically somewhat inaccurate at lower altitudes / mach numbers below 0.5, calculation is inhibited for mach < 0.395).
	TotalAirTemp *float64 `json:"tat"`

	// Statistics.
	// MessageCount is the total number of Mode S messages received from this aircraft.
	MessageCount int64 `json:"messages"`
	// LastSeen is how long ago (in seconds before "now") a message was last received from this aircraft.
	LastSeen *float64 `json:"seen"`

	// Source tracking.
	// MLATFields is the list of fields derived from MLAT data.
	MLATFields []string `json:"mlat"`
	// TISBFields is the list of fields derived from TIS-B data.
	TISBFields []string `json:"tisb"`

	// LastPosition is the last known position when the regular lat and lon are more than 60 seconds old. Aircraft will only be in the JSON if a position has been received in the last 60 seconds or if any message has been received in the last 30 seconds.
	LastPosition *LastPosition `json:"lastPosition"`

	// Estimated position (multi-receiver aggregation).
	RoughLatitude  *float64 `json:"rr_lat"`
	RoughLongitude *float64 `json:"rr_lon"`

	// ACAS Resolution Advisory (experimental).
	// ACASResolutionAdvisory is the ACAS resolution advisory. This is experimental and subject to change. See format here: https://github.com/wiedehopf/readsb/blob/caf2c5a6dccf1e079d0ae63d2b7dd45d568de265/json_out.c#L262.
	ACASResolutionAdvisory *ACASResolutionAdvisory `json:"acas_ra"`

	// GPS status (experimental).
	// GPSOkBefore is the timestamp at which the GPS was last working well, if the aircraft has lost GPS or its GPS is heavily degraded. This is only included for 15 min after GPS is lost / degraded. This is experimental and subject to change.
	GPSOkBefore *int64 `json:"gpsOkBefore"`
}
