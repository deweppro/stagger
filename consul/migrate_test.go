package consul

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateConsulMigrate(t *testing.T) {
	cli, err := CreateConsulMigrate("consul:8500")
	require.NoError(t, err)

	actual, err := cli.Dump()
	require.NoError(t, err)
	require.Len(t, actual, 0)

	required := KVList{
		{Key: "aaa", Value: "aaa", Type: ""},
		{Key: "bbb", Value: "bbb", Type: ""},
	}
	require.NoError(t, cli.Migrate(required))

	actual, err = cli.Dump()
	require.NoError(t, err)
	require.Len(t, actual, 2)

	sort.Slice(actual, func(i, j int) bool {
		return actual[i].Key < actual[j].Key
	})

	require.Equal(t, required, actual, "required %v actual %v", required, actual)

	required = KVList{
		{Key: "aaa", Value: "", Type: ""},
		{Key: "bbb", Value: "", Type: ""},
	}
	require.NoError(t, cli.Migrate(required))

	actual, err = cli.Dump()
	require.NoError(t, err)
	require.Len(t, actual, 0)
}
