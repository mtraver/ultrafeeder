package ultrafeeder

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	ground = "ground"
)

// BarometricAltitude represents an aircraft's barometric altitude, which can be
// either an integer number of feet or the special value "ground".
type BarometricAltitude struct {
	Feet   int
	Ground bool
}

func (b *BarometricAltitude) UnmarshalJSON(data []byte) error {
	// Try string first ("ground").
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		if s == ground {
			b.Ground = true
			return nil
		}
		return fmt.Errorf("unexpected string value for barometric altitude: %q", s)
	}

	// Otherwise expect an integer
	var n int
	if err := json.Unmarshal(data, &n); err != nil {
		return fmt.Errorf("barometric altitude must be an integer or \"ground\": %w", err)
	}
	b.Feet = n
	return nil
}

func (b BarometricAltitude) MarshalJSON() ([]byte, error) {
	if b.Ground {
		return json.Marshal(ground)
	}
	return json.Marshal(b.Feet)
}

func (b BarometricAltitude) String() string {
	if b.Ground {
		return ground
	}
	return strconv.Itoa(b.Feet)
}
