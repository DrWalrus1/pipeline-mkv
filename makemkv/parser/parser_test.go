package parser_test

import (
	"servermakemkv/makemkv/commands/outputs"
	"servermakemkv/makemkv/parser"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestEmptyStringFailsParsing(t *testing.T) {
	actual, err := parser.Parse("")

	assert.Equal(t, parser.EmptyInput, err)
	assert.Equal(t, nil, actual)
}

func TestStringSanitisation(t *testing.T) {
	actual, err := parser.Parse(" SINFO:0,6,31,6121,<b>Track information</b><br>")
	assert.Equal(t, nil, err)
	assert.Equal(t, actual.(*outputs.StreamInformation).Value, "Track information")
}

func TestCurrentProgressTitleOutputParser(t *testing.T) {
	t.Run("Successful Parse", func(t *testing.T) {
		expected := outputs.CurrentProgressTitleOutput{
			Code: "1",
			ID:   "1",
			Name: "Test",
		}

		input := "PRGC:1,1,Test"

		actual, err := parser.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails when missing prefix", func(t *testing.T) {
		input := "1,1,Test"

		actual, err := parser.Parse(input)

		assert.Equal(t, parser.PrefixNotFound, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Fails when missing values", func(t *testing.T) {
		input := "PRGC:1,1"

		actual, err := parser.Parse(input)

		assert.Equal(t, parser.NotEnoughValues, err)
		assert.Equal(t, nil, actual)
	})
}

func TestParseDiscInformationOutput(t *testing.T) {
	t.Run("Successful Parse", func(t *testing.T) {
		expected := outputs.DiscInformationOutputMessage{
			TitleCount: 1,
		}

		input := "TCOUNT:1"

		actual, err := parser.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails when missing prefix", func(t *testing.T) {
		input := "1"

		actual, err := parser.Parse(input)

		assert.Equal(t, parser.PrefixNotFound, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Fails when missing values", func(t *testing.T) {
		input := "TCOUT:"
		actual, err := parser.Parse(input)

		assert.NotEqual(t, nil, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Fails when invalid value is provided", func(t *testing.T) {
		input := "TCOUT:THISISNOTANUMBER"

		actual, err := parser.Parse(input)

		assert.NotEqual(t, nil, err)
		assert.Equal(t, nil, actual)
	})
}

func TestParseDiscInfo(t *testing.T) {
	t.Run("Successful Parse", func(t *testing.T) {
		expected := outputs.DiscInformation{
			ID:            1,
			MessageCodeId: 1,
			Value:         "Value",
		}

		input := "CINFO:1,1,Value"

		actual, err := parser.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails when missing prefix", func(t *testing.T) {
		input := "1,CODE,Value"

		actual, err := parser.Parse(input)

		assert.Equal(t, parser.PrefixNotFound, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Fails when missing values", func(t *testing.T) {
		input := "CINFO:1,CODE"

		actual, err := parser.Parse(input)

		assert.Equal(t, parser.NotEnoughValues, err)
		assert.Equal(t, nil, actual)
	})
}

func TestDriveScanMessageParser(t *testing.T) {
	t.Run("Parse Successfully", func(t *testing.T) {
		expected := outputs.DriveScanMessage{
			DriveIndex: "1",
			Visible:    true,
			Enabled:    true,
			Flags:      "Flags",
			DriveName:  "Drive1",
			DiscName:   "Disc1",
			DeviceName: "/dev/sr0",
		}

		input := "DRV:1,1,1,Flags,Drive1,Disc1,/dev/sr0"

		actual, err := parser.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails when missing prefix", func(t *testing.T) {
		input := "1,true,true,Flags,Drive1,Disc1,/dev/sr0"

		actual, err := parser.Parse(input)

		assert.Equal(t, parser.PrefixNotFound, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Fails when there are not enough values", func(t *testing.T) {
		input := "DRV:1,true,true,Flags,Drive1,/dev/sr0"

		actual, err := parser.Parse(input)

		assert.Equal(t, parser.NotEnoughValues, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Parses when drives are not visible or enabled", func(t *testing.T) {
		input := "DRV:1,255,999,Flags,Drive1,Disc1,/dev/sr0"

		expected := outputs.DriveScanMessage{
			DriveIndex: "1",
			Visible:    false,
			Enabled:    false,
			Flags:      "Flags",
			DriveName:  "Drive1",
			DiscName:   "Disc1",
			DeviceName: "/dev/sr0",
		}
		actual, err := parser.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)

	})
}

func TestParseMessageOutput(t *testing.T) {
	t.Run("Successfully parse", func(t *testing.T) {
		t.Run("No params", func(t *testing.T) {
			expected := outputs.MessageOutput{
				Code:           "1",
				Flags:          "test",
				ParameterCount: 0,
				RawMessage:     "Test Message",
				FormatMessage:  "Test Message",
			}

			input := "MSG:1,test,0,Test Message,Test Message"

			actual, err := parser.Parse(input)

			assert.Equal(t, nil, err)
			assert.Equal(t, expected, actual)

			expected2 := outputs.MessageOutput{
				Code:           "1011",
				Flags:          "0",
				ParameterCount: 1,
				RawMessage:     "Using LibreDrive mode (v06.3 id=0FA242DD4D0B)",
				FormatMessage:  "%1",
				MessageParams:  []string{"Using LibreDrive mode (v06.3 id=0FA242DD4D0B)"},
			}
			input2 := `MSG:1011,0,1,"Using LibreDrive mode (v06.3 id=0FA242DD4D0B)","%1","Using LibreDrive mode (v06.3 id=0FA242DD4D0B)"`

			actual2, err2 := parser.Parse(input2)

			assert.Equal(t, nil, err2)
			assert.Equal(t, expected2, actual2)
		})

		t.Run("One param", func(t *testing.T) {
			expected := outputs.MessageOutput{
				Code:           "1",
				Flags:          "test",
				ParameterCount: 1,
				RawMessage:     "Test Message",
				FormatMessage:  "Test Message",
				MessageParams:  []string{"Test"},
			}

			input := "MSG:1,test,1,Test Message,Test Message,Test"

			actual, err := parser.Parse(input)

			assert.Equal(t, nil, err)
			assert.Equal(t, expected, actual)
		})

		t.Run("Three params", func(t *testing.T) {
			expected := outputs.MessageOutput{
				Code:           "1",
				Flags:          "test",
				ParameterCount: 3,
				RawMessage:     "Test Message",
				FormatMessage:  "Test Message",
				MessageParams:  []string{"Test1", "Test2", "Test3"},
			}

			input := "MSG:1,test,3,Test Message,Test Message,Test1,Test2,Test3"

			actual, err := parser.Parse(input)

			assert.Equal(t, nil, err)
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("Fails when missing prefix", func(t *testing.T) {
		input := "1,test,3,Test Message,Test Message,Test1,Test2,Test3"

		actual, err := parser.Parse(input)

		assert.Equal(t, parser.PrefixNotFound, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Fails when parameter count mismatch", func(t *testing.T) {
		t.Run("Param count is less than actual", func(t *testing.T) {
			expected := outputs.MessageOutput{
				Code:           "1",
				Flags:          "test",
				ParameterCount: 1,
				RawMessage:     "Test Message",
				FormatMessage:  "Test Message",
				MessageParams:  []string{"Test1"},
			}

			input := "MSG:1,test,1,Test Message,Test Message,Test1,Test2,Test3"

			actual, err := parser.Parse(input)

			assert.Equal(t, nil, err)
			assert.Equal(t, expected, actual)
		})

		t.Run("Param count is greater than actual", func(t *testing.T) {
			expected := outputs.MessageOutput{
				Code:           "1",
				Flags:          "test",
				ParameterCount: 3,
				RawMessage:     "Test Message",
				FormatMessage:  "Test Message",
			}

			input := "MSG:1,test,3,Test Message,Test Message,Test1"

			actual, err := parser.Parse(input)

			assert.Equal(t, nil, err)
			assert.Equal(t, expected, actual)
		})
	})
}

func TestParseProgressBarOutput(t *testing.T) {
	t.Run("Successful parse", func(t *testing.T) {
		expected := outputs.ProgressBarOutput{
			CurrentProgress: "1",
			TotalProgress:   "100",
			MaxProgress:     "200",
		}
		input := "PRGV:1,100,200"

		actual, err := parser.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails to parse", func(t *testing.T) {
		t.Run("Missing prefix", func(t *testing.T) {
			input := "1,100,200"

			actual, err := parser.Parse(input)

			assert.Equal(t, parser.PrefixNotFound, err)
			assert.Equal(t, nil, actual)
		})

		t.Run("Progress bar missing values", func(t *testing.T) {
			input := "PRGV:1,100"

			actual, err := parser.Parse(input)

			assert.Equal(t, parser.NotEnoughValues, err)
			assert.Equal(t, nil, actual)
		})
	})
}

func TestParseStreamInfo(t *testing.T) {
	t.Run("Successful parse", func(t *testing.T) {
		expected := outputs.StreamInformation{
			TitleIndex:    1,
			StreamIndex:   1,
			AttributeId:   1,
			MessageCodeId: 1,
			Value:         "Value",
		}

		input := "SINFO:1,1,1,1,Value"

		actual, err := parser.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails to parse", func(t *testing.T) {
		t.Run("Missing prefix", func(t *testing.T) {
			input := "1,CODE,Value"

			actual, err := parser.Parse(input)

			assert.Equal(t, parser.PrefixNotFound, err)
			assert.Equal(t, nil, actual)
		})

		t.Run("Missing values", func(t *testing.T) {
			input := "SINFO:1,CODE"

			actual, err := parser.Parse(input)

			assert.Equal(t, parser.NotEnoughValues, err)
			assert.Equal(t, nil, actual)
		})
	})
}

func TestParseTitleInfo(t *testing.T) {
	t.Run("Successful parse", func(t *testing.T) {
		expected := outputs.TitleInformation{
			TitleIndex:    1,
			AttributeId:   1,
			MessageCodeId: 1,
			Value:         "Value",
		}

		input := "TINFO:1,1,1,Value"

		actual, err := parser.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails to parse", func(t *testing.T) {
		input := "1,CODE,Value"

		actual, err := parser.Parse(input)

		assert.NotEqual(t, nil, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Missing values", func(t *testing.T) {
		input := "TINFO:1,CODE"

		actual, err := parser.Parse(input)

		assert.Equal(t, parser.NotEnoughValues, err)
		assert.Equal(t, nil, actual)
	})
}

func TestParseTotalTitleOutput(t *testing.T) {
	t.Run("Successful parse", func(t *testing.T) {
		expected := outputs.TotalProgressTitleOutput{
			Code: "1",
			ID:   "1",
			Name: "Test",
		}

		input := "PRGT:1,1,Test"

		actual, err := parser.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails to parse", func(t *testing.T) {
		t.Run("Missing prefix", func(t *testing.T) {
			input := "1,1,Test"

			actual, err := parser.Parse(input)

			assert.Equal(t, parser.PrefixNotFound, err)
			assert.Equal(t, nil, actual)
		})

		t.Run("Missing values", func(t *testing.T) {
			input := "PRGT:1,1"

			actual, err := parser.Parse(input)

			assert.Equal(t, parser.NotEnoughValues, err)
			assert.Equal(t, nil, actual)
		})
	})
}
