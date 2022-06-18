package test_gorm

import (
	"testing"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ClassName string
}

type Student struct {
	gorm.Model
	ClassID uint
	IDCard IDCard
	Teachers []Teacher `gorm:"many2many:Student_Teacher;"`
	StudentName string
}

type IDCard struct {
	gorm.Model
	Num int
	StudentID uint
}

type Teacher struct {
	gorm.Model
	TeacherName string
	Students []Student `gorm:"many2many:Student_Teacher;"`
}

func TestMigrateRelationTable(t *testing.T) {
	db.AutoMigrate(&Class{}, &Teacher{}, &Student{}, &IDCard{})
}


