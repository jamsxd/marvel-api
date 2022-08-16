package domain

type Character struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Modified    string    `json:"modified"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	ResourceURI string    `json:"resourceURI"`
	Comics      Comics    `json:"comics"`
	Series      Comics    `json:"series"`
	Stories     Stories   `json:"stories"`
	Events      Comics    `json:"events"`
	Urls        []URL     `json:"urls"`
}

type Comics struct {
	Available     int64        `json:"available"`
	CollectionURI string       `json:"collectionURI"`
	Items         []ComicsItem `json:"items"`
	Returned      int64        `json:"returned"`
}

type ComicsItem struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}

type Stories struct {
	Available     int64         `json:"available"`
	CollectionURI string        `json:"collectionURI"`
	Items         []StoriesItem `json:"items"`
	Returned      int64         `json:"returned"`
}

type StoriesItem struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Type        Type   `json:"type"`
}

type Thumbnail struct {
	Path      string `json:"path"`
	Extension string `json:"extension"`
}

type URL struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Type string

const (
	Cover         Type = "cover"
	InteriorStory Type = "interiorStory"
)
