package qq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeywordSearch(t *testing.T) {
	results, err := _q.KeywordSearch(_ctx, "ζειΆθ‘")
	require.Nil(t, err)
	t.Log(results)
}
