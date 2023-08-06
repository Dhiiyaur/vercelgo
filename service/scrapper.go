package service

import (
	"context"
	"strings"
	"verceltest/entity"

	"github.com/gocolly/colly"
)

func MangatownIndex(ctx context.Context) ([]entity.MangaData, error) {

	var returnData []entity.MangaData

	c := colly.NewCollector()
	c.OnHTML("body > section > article > div > div.manga_pic_content > ul.manga_pic_list > li", func(e *colly.HTMLElement) {
		var mangaId, lastChapter string

		mangaIdCheck := strings.Split(e.ChildAttr("a", "href"), "/")
		mangaId = mangaIdCheck[len(mangaIdCheck)-2]

		lastChapterCheck := strings.Split(e.ChildText("p.new_chapter"), " ")
		lastChapter = lastChapterCheck[len(lastChapterCheck)-1]

		mangaCoverCheck := strings.Replace(e.ChildAttr("a > img", "src"), "https://fmcdn.mangahere.com/", "http://fmcdn.mangatown.com/", -1)

		returnData = append(returnData, entity.MangaData{
			Id:          mangaId,
			Title:       e.ChildAttr("a", "title"),
			Cover:       mangaCoverCheck,
			LastChapter: lastChapter,
		})

	})
	err := c.Visit("https://www.mangatown.com/hot/" + "1" + ".htm?wviews.za")
	if err != nil {
		return returnData, err
	}

	return returnData, nil

}
