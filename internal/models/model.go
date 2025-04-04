package models

type CreateLinkResponse struct {
	ShortLink string `json:"new_link"`
}

type CreateLinkData struct {
	CustomLink  string `json:"custom_link"`
	InitialLink string `json:"source_link"`
	UserId      int64  `json:"user_id"`
}

type RedirectLinkData struct {
	Err      error  `json:"error"`
	FullLink string `json:"full_link"`
	Code     int
}
