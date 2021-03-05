package model

type MapArticleTag struct {
	MapId     uint   `gorm:"column:map_id" json:"map_id"`
	ArticleId uint   `gorm:"column:article_id" json:"article_id"`
	TagId     uint16 `gorm:"column:tag_id" json:"tag_id"`
}

// type MapATResult struct {
// }

var TagList []string

//Create map
func CreateATMap(id int) {

}

//Delete map

//Edit map

//Get one map

//Get map list
