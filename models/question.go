package models

import (
	"time"

	"gorm.io/datatypes"
)

type Question struct {
	Id           int            `gorm:"id" json:"id"`
	Content      string         `gorm:"content" json:"content"`           // 题目内容
	ImageId      int            `gorm:"image_id" json:"image_id"`         // 题目关联的图片
	RightAnswer  string         `gorm:"right_answer" json:"right_answer"` // 正确答案
	Answers      datatypes.JSON `gorm:"answers" json:"answers"`           // 选项
	CreatedAt    time.Time      `gorm:"created_at" json:"created_at"`     // 创建时间
	Keyword      string         `gorm:"keyword" json:"keyword"`           // 题目的关键词
	QuestionType string         `gorm:"question_type" json:"type"`        // 题目类型， exam =测试, query = 答疑
}
