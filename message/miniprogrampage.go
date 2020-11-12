package message


// miniprogrampage  小程序消息
type Miniprogrampage struct {
	CommonToken

	Miniprogrampage_content struct {
		Title          string `json:"title"`
		Appid          string `json:"appid"`
		Pagepath       string `json:"pagepath"`
		Thumb_media_id string `json:"thumb_media_id"`
	}  `json:"miniprogrampage"`
}


func NewMiniprogrampage(title,appid,pagepath,thumb_media_id string ) *Miniprogrampage{
	miniprogrampage := new(Miniprogrampage)
	miniprogrampage.Miniprogrampage_content.Title = title
	miniprogrampage.Miniprogrampage_content.Appid = appid
	miniprogrampage.Miniprogrampage_content.Pagepath = pagepath
	miniprogrampage.Miniprogrampage_content.Thumb_media_id = thumb_media_id
	return miniprogrampage
}