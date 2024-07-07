package cleanstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCleanstring(t *testing.T) {
	type testCase struct {
		Description string
		Input       string
		Expected    string
	}

	testCases := []testCase{
		{
			Description: "Non-matching strings should be unmodified.",
			Input:       "foo",
			Expected:    "foo",
		},
		{
			Description: "Whitespace-only prefix lines should be stripped.",
			Input:       "\n\nfoo",
			Expected:    "foo",
		},
		{
			Description: "Different whitespaces should be stripped",
			Input:       "\n\n\t\n  \t  \nfoo",
			Expected:    "foo",
		},
		{
			Description: "Trailing whitespace should be stripped",
			Input:       "foo\n\n",
			Expected:    "foo",
		},
		{
			Description: "Prefixes should be stripped",
			Input:       "  |foo",
			Expected:    "foo",
		},
		{
			Description: "Whitespace-only suffix lines should be stripped",
			Input:       "  |foo\n\n \n",
			Expected:    "foo",
		},
		{
			Description: "Prefix and suffix lines are trimmed but intermediate lines are untouched",
			Input:       "\n\n          |foo\n\n          |bar\n\n",
			Expected:    "foo\n\nbar",
		},
		{
			Description: "Whitespace only prefix lines preserve whitespace",
			Input:       "\n\n          |foo\n          |    \n          |bar\n\n",
			Expected:    "foo\n    \nbar",
		},
	}
	for _, tc := range testCases {
		output := Get(tc.Input)
		require.Equal(t, tc.Expected, output, tc.Description)
	}
}