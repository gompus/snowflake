package snowflake

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type test struct {
	Snowflake Snowflake `json:"snowflake,omitempty"`
}

func TestParse(t *testing.T) {
	require := require.New(t)

	raw := "829388262825132092"
	s, err := Parse(raw)

	require.NoError(err)

	fmt.Println(s)
}

func TestSnowflake_UnmarshalJSONWithInt(t *testing.T) {
	require := require.New(t)

	raw, test := "{\"snowflake\": 829388262825132092}", new(test)

	err := json.Unmarshal([]byte(raw), &test)
	require.NoError(err)

	fmt.Println(test.Snowflake)
}

func TestSnowflake_UnmarshalJSONWithString(t *testing.T) {
	require := require.New(t)

	raw, test := "{\"snowflake\": \"829388262825132092\"}", new(test)

	err := json.Unmarshal([]byte(raw), &test)
	require.NoError(err)

	fmt.Println(test.Snowflake)
}

func TestSnowflake_UnmarshalJSONWithNonsense(t *testing.T) {
	require := require.New(t)

	raw, test := "{\"snowflake\": true}", new(test)

	err := json.Unmarshal([]byte(raw), &test)
	require.Error(err)
}
