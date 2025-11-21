package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:100;not null"`
	Email      string `gorm:"size:100;uniqueIndex;not null"`
	Age        uint
	PostsCount int `gorm:"default:0"`
	Posts      []Post
	Comments   []Comment
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type Post struct {
	ID             uint   `gorm:"primaryKey"`
	Title          string `gorm:"size:200;not null"`
	Content        string `gorm:"type:text;not null"`
	UserID         uint   `gorm:"not null"`
	CommentsStatus string
	Comments       []Comment
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	PostID    uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	// fmt.Println("Hello, World!")
	dsn := "root:root@tcp(localhost:3306)/go_demo?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("数据库连接失败", err)
	}
	fmt.Println("数据库连接成功")
	// 题目1
	// db.AutoMigrate(&User{}, &Post{}, &Comment{})
	// fmt.Println("自动迁移完成")
	// users := []User{
	// 	{Name: "张三", Email: "zhangsan@qq.com", Age: 50},
	// 	{Name: "李四", Email: "lisi@qq.com", Age: 40},
	// 	{Name: "王五", Email: "wangwu@qq.com", Age: 30},
	// }
	// err = batchCreateUsers(db, users)
	// if err != nil {
	// 	log.Fatalln("批量插入失败", err)
	// }
	// posts := []Post{
	// 	{Title: "第一篇文章", Content: "这是张三的第一篇文章内容", UserID: 1},
	// 	{Title: "第二篇文章", Content: "这是李四的第一篇文章内容", UserID: 2},
	// 	{Title: "第三篇文章", Content: "这是王五的第一篇文章内容", UserID: 3},
	// 	{Title: "第四篇文章", Content: "这是张三的第二篇文章内容", UserID: 1},
	// }
	// err = batchCreatePosts(db, posts)
	// if err != nil {
	// 	log.Fatalln("批量插入失败", err)
	// }
	// comments := []Comment{
	// 	{Content: "这是对第一篇文章的评论", PostID: 1, UserID: 2},
	// 	{Content: "这是对第二篇文章的评论", PostID: 2, UserID: 3},
	// 	{Content: "这是对第三篇文章的评论", PostID: 3, UserID: 1},
	// 	{Content: "这是对第四篇文章的评论", PostID: 4, UserID: 2},
	// }
	// err = batchCreateComments(db, comments)
	// if err != nil {
	// 	log.Fatalln("批量插入失败", err)
	// }

	// createUser(db)
	// createWithSelectedFields(db)
	//题目2
	// createWithRelations(db)
	// selectOne(db)
	// selectMostCommentbyPost(db)

	//题目三
	// addPost(db)
	deleteComment(db)

}

func batchCreateUsers(db *gorm.DB, users []User) error {
	result := db.Create(&users)
	return result.Error
}
func batchCreatePosts(db *gorm.DB, posts []Post) error {
	result := db.Create(&posts)
	return result.Error
}
func batchCreateComments(db *gorm.DB, comments []Comment) error {
	result := db.Create(&comments)
	return result.Error
}
func createUser(db *gorm.DB) {
	user := User{
		Name:  "赵六",
		Email: "zhaoliu@example.com",
	}

	result := db.Create(&user)
	if result.Error != nil {
		log.Printf("创建用户失败: %v", result.Error)
		return
	}

	fmt.Printf("创建用户成功，ID: %d\n", user.ID)
}

// 使用指定字段创建
func createWithSelectedFields(db *gorm.DB) {
	user := User{
		Name:  "钱七",
		Email: "qianqi@example.com",
		Age:   100,
	}

	// 只插入Name和Email字段
	result := db.Select("Name", "Email").Create(&user)
	if result.Error != nil {
		log.Printf("创建失败: %v", result.Error)
		return
	}

	fmt.Printf("创建成功，ID: %d\n", user.ID)
}

// 创建关联数据
func createWithRelations(db *gorm.DB) {
	user := User{
		Name:  "二狗子",
		Email: "zhangsan@example.com",
		Posts: []Post{
			{Title: "第五篇文章", Content: "这是内容..."},
			{Title: "第六篇文章", Content: "这是另一篇内容..."},
		},
	}

	result := db.Create(&user)
	if result.Error != nil {
		log.Printf("创建失败: %v", result.Error)
		return
	}
	fmt.Printf("✅ 创建用户及关联数据成功\n")
}

func selectOne(db *gorm.DB) {
	var user User
	// result := db.First(&user, 2) // 根据主键查询
	result := db.Preload("Posts").Preload("Comments").First(&user, 2)
	// result := db.Preload("posts").First(&user, 2)
	if result.Error != nil {
		log.Printf("查询失败: %v", result.Error)
		return
	}

	fmt.Printf("查询成功: %+v\n", user)
}

func selectMostCommentbyPost(db *gorm.DB) {
	// var comment Comment
	var post Post
	// result := db.First(&user, 2) // 根据主键查询
	// result := db.Model("group by postId  order by postId desc").Find(&comment)
	// var Count int64
	var commentResult struct {
		PostID uint
		Count  int
	}

	// result := db.Debug().Model(&comment).Select("post_id").Group("post_id").Order("post_id desc limit 1").Count(&Count).Limit(1)
	// result := db.Debug().Select("post_id","").Count(Group("post_id").Order("post_id desc").Find(&commentResult)

	// result := db.Preload("comment").First(&comment, 2)
	err := db.Raw(`
		SELECT count(c.post_id )count,c.post_id   FROM go_demo.comments AS c 
		GROUP by c.post_id  ORDER  by count  desc limit 1 
		`).Scan(&commentResult).Error
	if err != nil {
		fmt.Printf("查询失败: %v", err)
	}

	err1 := db.Preload("Comments").First(&post, commentResult.PostID).Error

	if err1 != nil {
		fmt.Printf("查询失败: %v", err1)
	}

	// fmt.Printf("查询成功: %+v\n", Count)
	fmt.Printf("查询成功: %+v\n", commentResult)
	fmt.Printf("查询成功: %+v\n", post)
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 在这里编写你的逻辑代码
	fmt.Println("AfterDelete 钩子被触发了！")
	var commentCount int64
	err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error
	if err != nil {
		fmt.Errorf("统计评论数量失败: %v", err)
		return err

	}
	fmt.Println("postId=", c.PostID)
	fmt.Printf("post=%+v", c)
	if commentCount == 0 {
		err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comments_status", "无评论").Error
		if err != nil {
			return fmt.Errorf("更新文章评论状态失败: %v", err)
		}
		fmt.Println("文章评论状态已更新为无评论")
	}
	return nil
}
func deleteComment(db *gorm.DB) error {
	var comment Comment
	err := db.Where("id = ?", 9).First(&comment).Error
	if err != nil {

		return fmt.Errorf("无此条数据: %v", err)
	}

	result := db.Delete(&comment)
	if result.Error != nil {

		return fmt.Errorf("删除评论失败: %v", err)
	}
	return nil
}
func addPost(db *gorm.DB) error {
	post := Post{
		Title:   "新增文章标题",
		Content: "新增文章内容",
		UserID:  2,
	}
	err := db.Create(&post)
	if err != nil {
		return fmt.Errorf("新增文章失败: %v", err)
	}
	fmt.Println("新增文章成功")

	return nil
}
func (p *Post) AfterCreate(tx *gorm.DB) error {
	fmt.Println("p=", p)
	// 在这里编写你的逻辑代码
	fmt.Println("AfterCreate 钩子被触发了！")
	var postCount int64
	err := tx.Model(&Post{}).Where("user_id = ?", p.UserID).Count(&postCount).Error
	if err != nil {
		return fmt.Errorf("统计文章数量失败: %v", err)
	}
	fmt.Printf("postCount=%d", postCount)

	err1 := tx.Model(&User{}).Where("id = ?", p.UserID).Update("posts_count", postCount).Error
	if err1 != nil {
		return fmt.Errorf("更新用户文章数量失败: %v", err1)
	}
	fmt.Println("用户文章数量已更新")
	return nil
}
