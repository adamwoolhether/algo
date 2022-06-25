package algo

import "testing"

func TestLinter(t *testing.T) {
	type test struct {
		input string
		exp   bool
	}
	testCases := map[string]test{
		"validBraces": {
			input: "( var x = { y: [1, 2, 3] } )",
			exp:   true,
		},
		"invalidBraces": {
			input: "var x = { y : [1, 2, 3] } )",
			exp:   false,
		},
		"validGo": {
			input: "func test(input []string) string { return string[0] }",
			exp:   true,
		},
		"invalidGo": {
			input: "func fail(arr ]string) { return } ",
			exp:   false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			linter := newLinter()
			if got := linter.lint(tc.input); got != tc.exp {
				t.Errorf("exp %t, got %t", tc.exp, got)
			}
		})
	}
}
