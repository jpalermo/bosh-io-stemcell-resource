package boshio

type Stemcell struct {
	Name         string
	Version      string
	Light        *Metadata `json:"light"`
	Regular      *Metadata `json:"regular"`
	ForceRegular bool
	ForceLight   bool
}

type Metadata struct {
	URL    string
	Size   int64
	MD5    string
	SHA1   string
	SHA256 string
}

func (s Stemcell) Details() Metadata {
	if s.Light != nil && s.ForceRegular == false {
		return *s.Light
	}

	return *s.Regular
}

type Stemcells []Stemcell

func (s Stemcells) FindStemcellByVersion(version string) (Stemcell, bool) {
	for _, stemcell := range s {
		if stemcell.Version == version {
			return stemcell, true
		}
	}
	return Stemcell{}, false
}

func (s Stemcells) FilterByType() Stemcells {
	if s.forceRegular() {
		return s.regularStemcellsOnly()
	} else if s.forceLight()  {
		return s.lightStemcellsOnly()
	} else if s.supportsLight() == false {
		// all stemcells are Regular, no need to filter
		return s
	} else {
		return s.lightStemcellsOnly()
	}
}

func (s Stemcells) lightStemcellsOnly() Stemcells {
	filterFunc := func(stemcell Stemcell) bool {
		return stemcell.Light != nil
	}
	return s.filterStemcells(filterFunc)
}

func (s Stemcells) regularStemcellsOnly() Stemcells {
	filterFunc := func(stemcell Stemcell) bool {
		return stemcell.Regular != nil
	}
	return s.filterStemcells(filterFunc)
}

func (s Stemcells) filterStemcells(filterFunc func(Stemcell) bool) Stemcells {
	filteredStemcells := Stemcells{}
	for _, stemcell := range s {
		if filterFunc(stemcell) {
			filteredStemcells = append(filteredStemcells, stemcell)
		}
	}
	return filteredStemcells
}

func (s Stemcells) forceRegular() bool {
	for _, stemcell := range s {
		if stemcell.ForceRegular {
			return true
		}
	}
	return false
}

func (s Stemcells) forceLight() bool {
	for _, stemcell := range s {
		if stemcell.ForceLight {
			return true
		}
	}
	return false
}

func (s Stemcells) supportsLight() bool {
	for _, stemcell := range s {
		if stemcell.Light != nil {
			return true
		}
	}
	return false
}
