package response

type MusicMediaResponse struct {
	Id             uint   `json:"Id"`
	Name           string `json:"Name"`
	CreatorAddress string `json:"CreatorAddress"`
	Mp3File        []byte `json:"Mp3File"`
	ImgFile        []byte `json:"ImgFile"`
}
