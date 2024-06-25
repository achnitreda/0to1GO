package converter

import "testing"

func TestConvertBase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test hexadecimal conversion",
			input:    "1E (hex) files were added",
			expected: "30 files were added",
		},
		{
			name:     "Test binary conversion",
			input:    "It has been 10 (bin) years",
			expected: "It has been 2 years",
		},
		{
			name:     "Test lowercase conversion",
			input:    "I should stop SHOUTING (low)",
			expected: "I should stop shouting",
		},
		{
			name:     "Test uppercase conversion",
			input:    "Ready, set, go (up) !",
			expected: "Ready, set, GO!",
		},
		{
			name:     "Test capitalization",
			input:    "Welcome to the Brooklyn bridge (cap)",
			expected: "Welcome to the Brooklyn Bridge",
		},
		{
			name:     "Test modifying words with number (up, <number>)",
			input:    "This is so exciting (up, 2)",
			expected: "This is SO EXCITING",
		},
		{
			name:     "Test fixing indefinite articles",
			input:    "There it was. A amazing rock!",
			expected: "There it was. An amazing rock!\n",
		},
		{
			name:     "Test punctuation formatting",
			input:    "I was thinking ... You were right. I was sitting over there ,and then BAMM !!",
			expected: "I was thinking... You were right. I was sitting over there, and then BAMM!!\n",
		},
		{
			name:     "Test fixing apostrophes",
			input:    "I am exactly how they describe me: ' awesome '. As Elton John said: ' I am the most well-known homosexual in the world '",
			expected: "I am exactly how they describe me: 'awesome'. As Elton John said: 'I am the most well-known homosexual in the world'",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ConvertBase(test.input)
			if result != test.expected {
				t.Errorf("Expected: %q, but got: %q", test.expected, result)
			}
		})
	}
}
