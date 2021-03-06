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

func stringCommands() {
	resp := AlfredResponse{
		Items: []AlfredItem{
			alfredItemFromStringForwarded("length", true),
			alfredItemFromStringForwarded("words", true),
			alfredItemFromStringForwarded("lower", true),
			alfredItemFromStringForwarded("title", true),
			alfredItemFromStringForwarded("upper", true),
			alfredItemFromStringForwarded("pymod", true),
			alfredItemFromStringForwarded("unpymod", true),
			alfredItemFromStringForwarded("md5", true),
			alfredItemFromStringForwarded("sha1", true),
			alfredItemFromStringForwarded("sha256", true),
			alfredItemFromStringForwarded("sha512", true),
		},
	}
	resp.Print()
}

func hashString(hasher hash.Hash, toHash string) string {
	hasher.Write([]byte(toHash))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

func stringCommand(args []string) {
	if len(args) == 0 {
		stringCommands()
		return
	}

	if len(args) == 1 {
		return
	}

	subcmd := args[0]
	input_string := strings.Join(args[1:], " ")
	result := ""
	switch subcmd {
	case "length":
		result = fmt.Sprintf("%d", len(input_string))
	case "words":
		result = fmt.Sprintf("%d", len(strings.Fields(input_string)))
	case "lower":
		result = strings.ToLower(input_string)
	case "title":
		result = strings.ToTitle(input_string)
	case "upper":
		result = strings.ToUpper(input_string)
	case "pymod":
		no_slashes := strings.Replace(input_string, "/", ".", -1)
		result = strings.TrimSuffix(no_slashes, ".py")
	case "unpymod":
		result = strings.Replace(input_string, ".", "/", -1) + ".py"
	case "md5":
		result = hashString(md5.New(), input_string)
	case "sha1":
		result = hashString(sha1.New(), input_string)
	case "sha256":
		result = hashString(sha256.New(), input_string)
	case "sha512":
		result = hashString(sha512.New(), input_string)
	default:
		result = "Unknown string subcommand"
	}

	resp := AlfredResponse{
		Items: []AlfredItem{
			alfredItemFromString(result, false),
		},
	}
	resp.Print()
}
