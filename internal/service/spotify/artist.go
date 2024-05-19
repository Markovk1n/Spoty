package spotify

import "fmt"

type ArtistService service

type Artist struct {
	Name string `json:"name,omitempty"`
	ID   ID     `json:"id,omitempty"`
	URI  URI    `json:"uri,omitempty"`
	Href string `json:"href,omitempty"`

	Type       string   `json:"type,omitempty"`
	Genres     []string `json:"genres,omitempty"`
	Images     []Image  `json:"images,omitempty"`
	Popularity int      `json:"popularity,omitempty"`
}

type Artists struct {
	Href     string   `json:"href,omitempty"`
	Items    []Artist `json:"items,omitempty"`
	Limit    int      `json:"limit,omitempty"`
	Next     string   `json:"next,omitempty"`
	Offset   int      `json:"offset,omitempty"`
	Previous string   `json:"previous,omitempty"`
	Total    int      `json:"total,omitempty"`
}

type ArtistParams struct {
	IDs string `url:"ids,omitempty"`
}

type CountryParams struct {
	Country string `url:"country,omitempty"`
}

type ArtistAlbumsParams struct {
	IncludeGroups string `url:"include_groups,omitempty"`
	Market        string `url:"market,omitempty"`
	Limit         int    `url:"limit,omitempty"`
	Offset        int    `url:"offset,omitempty"`
}

type ArtistsResult struct {
	Artists []Artist `json:"artists,omitempty"`
}

func (s *ArtistService) List(IDs string) (*ArtistsResult, error) {
	params := &ArtistParams{IDs: IDs}
	var err error
	res := &ArtistsResult{}
	s.client.base.Get("artists").QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *ArtistService) Get(ID string) (*Artist, error) {
	var err error
	res := new(Artist)

	s.client.base.Path("artists/").Get(ID).Receive(res, err)
	return res, err
}

func (s *ArtistService) GetTopTracks(ID, country string) (*TracksResult, error) {
	var err error
	res := new(TracksResult)
	// params := &CountryParams{country}
	artists := fmt.Sprintf("artists/%s/", ID)

	s.client.base.Path(artists).Get("top-tracks").Receive(res, err)

	return res, err
}

func (s *ArtistService) GetAlbums(ID, includeGroups, market string, limit, offset int) (*Albums, error) {
	var err error
	res := new(Albums)
	params := &ArtistAlbumsParams{includeGroups, market, limit, offset}
	artists := fmt.Sprintf("artists/%s/", ID)
	s.client.base.Path(artists).Get("albums").QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *ArtistService) RelatedArtists(ID string) (*ArtistsResult, error) {
	var err error
	res := new(ArtistsResult)
	artists := fmt.Sprintf("artists/%s/", ID)
	s.client.base.Path(artists).Get("related-artists").Receive(res, err)
	return res, err
}
