package main

import (
	"log"

	"./pornapi"
)

func main() {
	api := pornapi.NewPornApi()

	//検索ワード指定
	api.Search("Japan")

	//ビデオ取得
	vkeys := api.GetVideos()

	//次のページ表示できる
	api.NextPage()

	//一番最初の動画だけ取得するなら
	/*
		vkey := vkeys[0]
		video,err := api.GetVideoInfo(vkey)
		log.Printf(`
		タイトル%v
		URL:%v
		評価:%v
		再生数:%v`,video.Name,video.Url,video.Rating,video.Count)
	*/
	videos, err := api.GetVideoInfo2(vkeys)
	if err != nil {
		log.Fatal(err)
	}
	for _, video := range videos {
		log.Printf(`
		タイトル:%v
		URL:%v
		評価:%v
		再生数:%v`, video.Name, video.Url, video.Rating, video.Count)
	}
}
