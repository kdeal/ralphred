package ralphred

import (
	"encoding/json"
	"errors"
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

type DevDocsDocEntry struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}

type DevDocsDocType struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
	Slug  string `json:"slug"`
}

type DevDocsDocIndex struct {
	Entries []DevDocsDocEntry `json:"entries"`
	Types   []DevDocsDocType  `json:"types"`
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

func fetchDevdocsDocIndex(docSlug string) (DevDocsDocIndex, error) {
	resp, err := http.Get(fmt.Sprintf("%s/docs/%s/index.json", DevdocsBaseUrl, docSlug))
	if err != nil {
		return DevDocsDocIndex{}, err
	}
	defer resp.Body.Close()
	var docIndex = DevDocsDocIndex{}
	err = json.NewDecoder(resp.Body).Decode(&docIndex)
	if err != nil {
		return DevDocsDocIndex{}, err
	}
	return docIndex, nil
}

func filterEntries(docEntries []DevDocsDocEntry, searchQuery []string) []DevDocsDocEntry {
	if len(searchQuery) == 0 {
		return docEntries
	}

	matchedEntries := []DevDocsDocEntry{}
	for _, entry := range docEntries {
		if queryMatches(entry.Name, searchQuery) {
			matchedEntries = append(matchedEntries, entry)
		}
	}
	return matchedEntries
}

func devdocsSearchDoc(docSlug string, searchQuery []string) ([]AlfredItem, error) {
	docsIndex, err := fetchDevdocsDocIndex(docSlug)
	if err != nil {
		return []AlfredItem{}, err
	}
	entries := filterEntries(docsIndex.Entries, searchQuery)
	if len(entries) > 25 {
		entries = entries[:25]
	}
	docItems := make([]AlfredItem, len(entries))
	for i, entry := range entries {
		docItems[i] = AlfredItem{
			UID:          entry.Name,
			Title:        entry.Name,
			Subtitle:     entry.Path,
			Arg:          []string{fmt.Sprintf("%s/%s/%s", DevdocsBaseUrl, docSlug, entry.Path)},
			Autocomplete: entry.Name,
		}
	}
	return docItems, nil
}

func devdocsCommand(args []string) ([]AlfredItem, error) {
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

	docItems = filterAlfredItems(docItems, args)

	return docItems, nil
}

func devdocsDocSetCommand(args []string) ([]AlfredItem, error) {
	if len(args) == 0 {
		return []AlfredItem{}, errors.New("No doc set specified")
	}

	searchQuery := []string{}
	if len(args) > 1 {
		searchQuery = args[1:]
	}
	return devdocsSearchDoc(args[0], searchQuery)
}
