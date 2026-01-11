package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notes struct {
	Id primitive.ObjectID `bson:"_id" json:"id"`
	Title string 					`bson:"title" json:"title"`
	Content string 				`bson:"content" json:"content"`
	Pinned bool 					`bson:"pinned" json:"pinned"`
	CreateAt  time.Time		`bson:"createAt" json:"createAt"`
	UpdateAt time.Time 		`bson:"updateAt" json:"updateAt"`
}

type createNoteRequest struct {
	 Title 		string 			`json:"title" binding:"required"`
	 Content 	string 			`json:"content" binding:"required"`
	 Pinned 	bool 				`json:"pinned" binding:"required"`
	 

}

func main() {

}