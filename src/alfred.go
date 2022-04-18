package ralphred

import (
	"encoding/json"
	"fmt"
	"os"
)

type AlfredItem struct {
	UID string `json:"uid,omitempty"`
	Title string `json:"title"`
	Subtitle string `json:"subtitle,omitempty"`
	Arg []string `json:"arg"`
	Autocomplete string `json:"autocomplete"`
}

func alfredItemFromString(str string, set_uid bool) AlfredItem {
	uid := ""
	if set_uid {
		uid = str
	}

	return AlfredItem{
		UID: uid,
		Title: str,
		Subtitle: "",
		Arg: []string{str},
		Autocomplete: str,
	}
}


func alfredItemFromStringForwarded(str string, set_uid bool) AlfredItem {
	uid := ""
	if set_uid {
		uid = str
	}

	return AlfredItem{
		UID: uid,
		Title: str,
		Subtitle: "",
		Arg: []string{fmt.Sprintf("%s ", str)},
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

func errorAlfredResponse(errMsg string) AlfredResponse {
	return AlfredResponse {
		Items: []AlfredItem{
			alfredItemFromString(errMsg, false),
		},
	}
}
