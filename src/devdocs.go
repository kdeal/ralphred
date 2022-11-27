package ralphred

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const DevdocsBaseUrl string = "https://devdocs.io/"

type DevdocsDocSet struct {
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Type    string `json:"type"`
	Version string `json:"version"`
	Release string `json:"release"`
	Mtime   int64  `json:"mtime"`
	DBSize  int64  `json:"db_size"`
}

func fetchDevdocsDocsList() ([]DevdocsDocSet, error) {
	resp, err := http.Get(fmt.Sprintf("%s/docs/docs.json", DevdocsBaseUrl))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var docsList = []DevdocsDocSet{}
	err = json.NewDecoder(resp.Body).Decode(&docsList)
	if err != nil {
		return nil, err
	}
	return docsList, nil
}

func devdocsListDocs() ([]AlfredItem, error) {
	docsList, err := fetchDevdocsDocsList()
	if err != nil {
		return nil, err
	}

	docItems := make([]AlfredItem, len(docsList))
	for i, doc := range docsList {
		docItems[i] = AlfredItem{
			UID:          doc.Slug,
			Title:        fmt.Sprintf("%s %s", doc.Name, doc.Release),
			Subtitle:     doc.Slug,
			Arg:          []string{fmt.Sprintf("%s ", doc.Slug)},
			Autocomplete: doc.Name,
		}
	}

	return docItems, nil
}

func devdocsCommand(args []string) ([]AlfredItem, error) {
	if len(args) == 0 {
		return devdocsListDocs()
	}

	resp := []AlfredItem{
		alfredItemFromString("A response", false),
	}
	return resp, nil
}
