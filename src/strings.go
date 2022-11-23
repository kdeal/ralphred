package ralphred

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"
)

type StringConversion struct {
	Convert func(string) string
}

var string_conversions  = map[string]StringConversion {
	"length": {
		Convert: func(input_string string) string {
			return fmt.Sprintf("%d", len(input_string))
		},
	},
	"words": {
		Convert: func(input_string string) string {
			return fmt.Sprintf("%d", len(strings.Fields(input_string)))
		},
	},
	"lower": {
		Convert: func(input_string string) string {
			return strings.ToLower(input_string)
		},
	},
	"title": {
		Convert: func(input_string string) string {
			return strings.Title(strings.ToLower(input_string))
		},
	},
	"upper": {
		Convert: func(input_string string) string {
			return strings.ToUpper(input_string)
		},
	},
	"pymod": {
		Convert: func(input_string string) string {
			no_slashes := strings.Replace(input_string, "/", ".", -1)
			return strings.TrimSuffix(no_slashes, ".py")
		},
	},
	"unpymod": {
		Convert: func(input_string string) string {
			return strings.Replace(input_string, ".", "/", -1) + ".py"
		},
	},
	"md5": {
		Convert: func(input_string string) string {
			return hashString(md5.New(), input_string)
		},
	},
	"sha1": {
		Convert: func(input_string string) string {
			return hashString(sha1.New(), input_string)
		},
	},
	"sha256": {
		Convert: func(input_string string) string {
			return hashString(sha256.New(), input_string)
		},
	},
	"sha512": {
		Convert: func(input_string string) string {
			return hashString(sha512.New(), input_string)
		},
	},
}

func stringCommands() []AlfredItem {
	helpText := make([]AlfredItem, len(string_conversions))
	i := 0
	for command := range string_conversions {
		helpText[i] = alfredItemFromStringForwarded(command, true)
		i++
	}
	return helpText
}

func hashString(hasher hash.Hash, toHash string) string {
	hasher.Write([]byte(toHash))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

func stringCommand(args []string) ([]AlfredItem, error) {
	if len(args) == 0 {
		return stringCommands(), nil
	}

	if len(args) == 1 {
		return []AlfredItem{}, nil
	}

	subcmd := args[0]
	input_string := strings.Join(args[1:], " ")
	result := "Unknown string subcommand"

	converter, exists := string_conversions[subcmd]
	if exists {
		result = converter.Convert(input_string)
	}

	resp := []AlfredItem{
		alfredItemFromString(result, false),
	}
	return resp, nil
}
