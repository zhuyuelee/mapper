package mapper_test

import (
	"testing"
	"time"

	"github.com/zhuyuelee/mapper"
)

type BaseDto struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
type HobbyDto struct {
	Name string
}
type StudentDto struct {
	Name    string
	Hobbys  []HobbyDto
	BaseDto `mapper:"base"`
}

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
type HobbyModel struct {
	Name string
}

type StudentModel struct {
	Name      string
	Hobbys    []HobbyModel
	BaseModel `mapper:"base"`
}

func TestMapStruct(t *testing.T) {
	source := StudentModel{
		Name: "MyName",
		Hobbys: []HobbyModel{
			{
				Name: "Football",
			},
			{
				Name: "Basketball",
			},
		},
		BaseModel: BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	to := &StudentDto{}
	mapper.Map(source, to)
	t.Logf("to = %+v \n", to)
}

func TestMapSlice(t *testing.T) {
	source := []StudentModel{
		{
			Name: "MyName",
			Hobbys: []HobbyModel{
				{
					Name: "Football",
				},
				{
					Name: "Basketball",
				},
			},
			BaseModel: BaseModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	to := make([]StudentDto, 0)
	mapper.Map(source, &to)
	t.Logf("to = %+v \n", to)
}
