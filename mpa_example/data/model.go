package data

type Country struct {
	Name struct {
		Common     string `json:"common"`
		Official   string `json:"official"`
		NativeName struct {
			Spa struct {
				Official string `json:"official"`
				Common   string `json:"common"`
			} `json:"spa"`
		} `json:"nativeName"`
	} `json:"name"`
	Demonyms struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
		Fra struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"fra"`
	} `json:"demonyms"`
	Flags struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
		Alt string `json:"alt"`
	} `json:"flags"`
	Car struct {
		Side  string   `json:"side"`
		Signs []string `json:"signs"`
	} `json:"car"`
	Currencies struct {
		USD struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies"`
	Maps struct {
		GoogleMaps     string `json:"googleMaps"`
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps"`
	PostalCode struct {
		Format string `json:"format"`
		Regex  string `json:"regex"`
	} `json:"postalCode"`
	CoatOfArms struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"coatOfArms"`
	Status      string `json:"status"`
	StartOfWeek string `json:"startOfWeek"`
	Fifa        string `json:"fifa"`
	Cca2        string `json:"cca2"`
	Region      string `json:"region"`
	Subregion   string `json:"subregion"`
	Languages   struct {
		Spa string `json:"spa"`
	} `json:"languages"`
	Ccn3 string `json:"ccn3"`
	Cca3 string `json:"cca3"`
	Flag string `json:"flag"`
	Cioc string `json:"cioc"`
	Idd  struct {
		Root     string   `json:"root"`
		Suffixes []string `json:"suffixes"`
	} `json:"idd"`
	Continents  []string  `json:"continents"`
	Borders     []string  `json:"borders"`
	Latlng      []float64 `json:"latlng"`
	CapitalInfo struct {
		Latlng []float64 `json:"latlng"`
	} `json:"capitalInfo"`
	AltSpellings []string `json:"altSpellings"`
	Capital      []string `json:"capital"`
	Tld          []string `json:"tld"`
	Timezones    []string `json:"timezones"`
	Population   int      `json:"population"`
	Gini         struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini"`
	Landlocked  bool `json:"landlocked"`
	UnMember    bool `json:"unMember"`
	Independent bool `json:"independent"`
}
