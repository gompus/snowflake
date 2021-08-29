package snowflake

import (
	"strconv"
	"time"
)

// Snowflake represents a Discord snowflake.
// See https://discord.com/developers/docs/reference
// for more information.
type Snowflake int64

// Nil represents an empty Snowflake.
var Nil = (*Snowflake)(nil)

// Zero represents the zero value for a Snowflake.
var Zero = Snowflake(0)

// discordEpoch is the time of the first second of 2015, in milliseconds.
const discordEpoch = 1420070400000

// Parse constructs a Snowflake from s.
func Parse(s string) (Snowflake, error) {
	if s == "" {
		return Zero, nil
	}
	parsed, err := strconv.ParseInt(s, 10, 0)
	return Snowflake(parsed), err
}

// MustParse is like Parse, but panics if it encounters an error.
func MustParse(s string) Snowflake {
	flake, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return flake
}

// FromTimestamp generates a Snowflake for the given time.
func FromTimestamp(time time.Time) Snowflake {
	return Snowflake(time.UnixMilli()-discordEpoch) << 22
}

// Timestamp returns the number of milliseconds since discordEpoch.
func (s Snowflake) Timestamp() int64 {
	return (int64(s) >> 22) + discordEpoch
}

// WorkerID returns the Snowflake's worker id.
func (s Snowflake) WorkerID() int64 {
	return (int64(s) & 0x3E0000) >> 17
}

// ProcessID returns the Snowflake's process id.
func (s Snowflake) ProcessID() int64 {
	return (int64(s) & 0x1F000) >> 12
}

// Increment returns the number of id's that have
// been generated on the Snowflake's process.
func (s Snowflake) Increment() int64 {
	return int64(s) & 0xFFF
}

// UnmarshalJSON unmarshals data into s.
func (s *Snowflake) UnmarshalJSON(data []byte) error {
	start, stop := 0, len(data)
	if data[0] == '"' && data[len(data)-1] == '"' {
		start++
		stop--
	}

	raw := data[start:stop]
	if len(raw) == 0 {
		*s = 0
		return nil
	}

	parsed, err := strconv.Atoi(string(raw))
	if err != nil {
		return UnmarshalTypeError{
			value: string(data),
			typ:   "int or string",
		}
	}

	*s = Snowflake(parsed)
	return nil
}

// MarshalJSON returns the JSON representation of s.
func (s Snowflake) MarshalJSON() ([]byte, error) {
	return []byte(s.String()), nil
}

// String returns the string representation of s.
func (s Snowflake) String() string {
	return strconv.FormatInt(int64(s), 10)
}
