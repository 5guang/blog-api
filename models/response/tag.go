package response

type ResTag struct {
	Creator     string `json:"creator"`
	Updater     string `json:"updater"`
	ID          int    `json:"id"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
	Name        string `json:"name"`
}

type ResTagData struct {
	TagList []ResTag `json:"tagList"`
	Count   int      `json:"count"`
}
