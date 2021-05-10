package utilities

import (
	"testing"
)

func TestProgressBar(t *testing.T) {
	t.Parallel()

	tests := []struct {
		len, fill, all, filled, empty int
	}{
		{10, 3, 10, 3, 7},
		{20, 16, 100, 3, 17},
		{30, 100, 100, 30, 0},
	}

	for _, test := range tests {
		progress := ProgressBarString(test.len, test.fill, test.all)
		progressSymb, spaceSymb := 0, 0

		for _, ch := range progress {
			switch ch {
			case '#':
				progressSymb++
				break
			case '-':
				spaceSymb++
				break
			}
		}

		if test.filled != progressSymb || spaceSymb != test.empty {
			t.Errorf(
				"String:\n %sExpected %d filled, %d empty cells in progress bar\nFounded %d, %d",
				progress,
				test.filled,
				test.empty,
				progressSymb,
				spaceSymb,
			)
		}
	}
}
