package model

// ResData ...
type ResData struct {
	Showcnt, Clickcnt int
}

// ResMap ...
type ResMap map[string]map[string]*ResData
