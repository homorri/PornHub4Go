package pornapi

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type PornApi struct {
	document  *goquery.Document
	videoInfo *VideInfo
	page      int
	url       string
}
type VideInfo struct {
	Name   string
	Rating string
	Url    string
	Count  string
}

const (
	BaseUrl  = "https://pornhub.com"
	VideoUrl = BaseUrl + "/video/"
	ViewUrl  = BaseUrl + "/view_video.php"
)

///html/body/div[7]/div/div[3]/div/div/ul/li[2]
//
func NewPornApi() *PornApi {
	return &PornApi{videoInfo: &VideInfo{}}
}
func (p *PornApi) Search(word string) (*goquery.Document, error) {
	Url := VideoUrl + word
	p.url = Url
	doc, err := goquery.NewDocument(Url)
	p.document = doc
	return doc, err
}

func (p *PornApi) GetVideos() []string {
	var vkeys []string
	p.document.Find(".pcVideoListItem").Each(func(_ int, s *goquery.Selection) {
		attr, exists := s.Attr("_vkey")
		if exists {
			vkeys = append(vkeys, attr)
		}
	})
	return vkeys
}
func (p *PornApi) GetVideoInfo(vkey string) (*VideInfo, error) {
	url := ViewUrl + "?viewkey=" + vkey
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	inlineFree := doc.Find(".inlineFree")
	percent := doc.Find(".percent")
	count := doc.Find(".count")
	p.videoInfo.Count = count.Text()
	p.videoInfo.Name = inlineFree.Text()
	p.videoInfo.Rating = percent.Text()
	p.videoInfo.Url = url
	return p.videoInfo, nil
}
func (p *PornApi) GetVideoInfo2(vkeys []string) ([]*VideInfo, error) {
	var videos []*VideInfo
	for _, vkey := range vkeys {
		video, err := p.GetVideoInfo(vkey)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}
func (p *PornApi) NextPage() error {
	p.page++
	var err error
	p.document, err = goquery.NewDocument(p.url + "&page=" + strconv.Itoa(p.page))
	if err != nil {
		return err
	}
	return nil
}
