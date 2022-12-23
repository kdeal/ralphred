package ralphred

import (
	"encoding/json"
	"fmt"
	"os"
)

type AlfredItem struct {
	UID          string   `json:"uid,omitempty"`
	Title        string   `json:"title"`
	Subtitle     string   `json:"subtitle,omitempty"`
	Arg          []string `json:"arg"`
	Autocomplete string   `json:"autocomplete"`
}

func alfredItemFromString(str string, set_uid bool) AlfredItem {
	uid := ""
	if set_uid {
		uid = str
	}

	return AlfredItem{
		UID:          uid,
		Title:        str,
		Subtitle:     "",
		Arg:          []string{str},
		Autocomplete: str,
	}
}

func alfredItemFromStringForwarded(str string, set_uid bool) AlfredItem {
	uid := ""
	if set_uid {
		uid = str
	}

	return AlfredItem{
		UID:          uid,
		Title:        str,
		Subtitle:     "",
		Arg:          []string{fmt.Sprintf("%s ", str)},
		Autocomplete: str,
	}
}

type AlfredResponse struct {
	Items []AlfredItem `json:"items"`
}

func (resp AlfredResponse) Print() {
	json_data, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error formatting string commands json")
		return
	}
	fmt.Println(string(json_data))
}

func errorAlfredItems(errMsg string) []AlfredItem {
	return []AlfredItem{
		alfredItemFromString(errMsg, false),
	}
}

// Filter alfred items by the search query. This shouldn't be used when there
// are a large amount of items
func filterAlfredItems(items []AlfredItem, searchQuery []string) []AlfredItem {
	matchedItems := []AlfredItem{}
	for _, item := range items {
		if queryMatches(item.Title, searchQuery) {
			matchedItems = append(matchedItems, item)
		}
	}

	return matchedItems
}
