package blog

import (
	"database/sql"
	"time"

	"github.com/ligenhw/goshare/orm"
)

type Comment struct {
	Id       int       `json:"id"`
	BlogId   int       `json:"blogId"`
	UserId   int       `json:"userId"`
	ParentId *int      `json:"parentId" orm:"null"`
	ReplyTo  *int      `json:"replyTo" orm:"null"`
	Content  string    `json:"content"`
	Time     time.Time `json:"time"`
}

type CommentWithChild struct {
	ParentId      int        `json:"parentId"`
	ParentUserId  int        `json:"parentUserId"`
	ParentContent string     `json:"parentContent"`
	ParentTime    time.Time  `json:"parentTime"`
	SubComments   []*Comment `json:"subComments"`
}

func init() {
	orm.RegisterModel(new(Comment))
}

// CreateComment create a parent comment
func CreateComment(blogId, userId int, content string) (id int64, err error) {
	b := &Comment{
		BlogId:  blogId,
		UserId:  userId,
		Content: content,
		Time:    time.Now(),
	}

	id, err = o.Insert(b)
	return
}

func CreateReply(blogId, userId, parentId, replyTo int, content string) (err error) {
	b := &Comment{
		BlogId:   blogId,
		UserId:   userId,
		ParentId: &parentId,
		ReplyTo:  &replyTo,
		Content:  content,
		Time:     time.Now(),
	}
	_, err = o.Insert(b)
	return
}

const queryCommentsList = `
select c1.id parent_id, c1.user_id parent_user_id, c1.content parent_content, c1.time parent_time,
 c2.id, c2.user_id, c2.reply_to, c2.content, c2.time from
 comment c1 left join comment c2 on c2.parent_id = c1.id
 where c1.parent_id is null and c1.blog_id = ?
 order by c1.time desc, c2.time ;`

func QueryCommentsByBlogId(blogId int) (comments []*CommentWithChild, err error) {
	var rows *sql.Rows
	if rows, err = db.Query(queryCommentsList, blogId); err != nil {
		return
	}

	comments = make([]*CommentWithChild, 0)
	var comment *CommentWithChild
	for rows.Next() {
		var parent_id int
		var parent_user_id int
		var parent_content string
		var parent_time time.Time
		var subComment *Comment

		var id sql.NullInt64
		var userId sql.NullInt64
		var replyTo sql.NullInt64
		var content sql.NullString
		var timeValue interface{}
		if err = rows.Scan(&parent_id, &parent_user_id, &parent_content, &parent_time, &id, &userId, &replyTo, &content, &timeValue); err != nil {
			return
		} else if id.Valid {
			replyToPtr := int(replyTo.Int64)
			subComment = &Comment{
				Id:      int(id.Int64),
				UserId:  int(userId.Int64),
				ReplyTo: &replyToPtr,
				Content: content.String,
				Time:    timeValue.(time.Time),
			}
		}
		if comment == nil || comment.ParentId != parent_id {
			comment = &CommentWithChild{
				ParentId:      parent_id,
				ParentUserId:  parent_user_id,
				ParentContent: parent_content,
				ParentTime:    parent_time,
				SubComments:   make([]*Comment, 0),
			}
			comments = append(comments, comment)
		}

		comment.SubComments = append(comment.SubComments, subComment)
	}
	return
}

func QueryCommentsByReply(replyTo int) (c []*Comment, err error) {
	c = make([]*Comment, 0)
	qb := o.QueryTable(new(Comment))
	_, err = qb.Filter("reply_to", replyTo).OrderBy("-time").All(&c)
	return
}

// func QueryByBlogId(blogId int) (comments []*Comment, err error) {
// 	comments = make([]*Comment, 0)
// 	qs := o.QueryTable(new(Comment))
// 	_, err = qs.Filter("blog_id", blogId).OrderBy("-time").All(&comments)
// 	return
// }
