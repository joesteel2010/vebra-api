package model

import (
	"fmt"
	"regexp"
	"strconv"
)

type ChangedPropertySummary struct {
	PropertyID  uint   `xml:"propid"`
	LastChanged string `xml:"lastchanged"`
	Url         string `xml:"url"`
	LastAction  string `xml:"action"`
}

type ChangedPropertySummaries struct {
	PropertySummaries []ChangedPropertySummary `xml:"property"`
}

func (cps *ChangedPropertySummary) SetLastAction(lastAction string) {
	cps.LastAction = lastAction
}

func (cps *ChangedPropertySummary) GetLastAction() string {
	return cps.LastAction
}

/**
 * get ClientId
 *
 * @return int $clientId
 */

func (cps *ChangedPropertySummary) GetClientID() (int, error) {
	re := regexp.MustCompile(`branch\/(\d+)/`)
	matches := re.FindStringSubmatch(cps.Url)

	if matches == nil || len(matches) < 2 {
		return 0, fmt.Errorf("couldnt match client ID in URL: [%s]", cps.Url)
	}

	return strconv.Atoi(matches[1])
}
