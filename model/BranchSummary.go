package model

import (
	"strconv"
	"strings"
)

type BranchSummaries struct {
	Branches []BranchSummary `xml:"branch"`
}

type BranchSummary struct {
	Name     string `xml:"name"`
	FirmID   int    `xml:"firmid"`
	BranchID int    `xml:"branchid"`
	Url      string `xml:"url"`
}

func (bs BranchSummary) GetClientID() (int, error) {
	index := strings.LastIndex(bs.Url, "/")
	return strconv.Atoi(bs.Url[(index + 1):])
}
