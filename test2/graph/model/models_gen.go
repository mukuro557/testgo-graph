// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewAddress struct {
	City       string `json:"city"`
	District   string `json:"district"`
	SubDistric string `json:"subDistric"`
}

type NewOwner struct {
	Name     string `json:"name"`
	IsSingle bool   `json:"isSingle"`
}

type OldOwner struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsSingle bool   `json:"isSingle"`
}

type Owner struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsSingle bool   `json:"isSingle"`
	Address  string `json:"address"`
}

type Delete struct {
	ID string `json:"id"`
}
