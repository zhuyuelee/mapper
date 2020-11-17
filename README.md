# Mapper 包

支持struct map to  struct  slice map to  slice
tag标签 mapper
 ```go
 Name string `mapper:"name"`
```
## 安装

```
go get github.com/zhuyuelee/mapper
```

## example for map to Struct
``` 
type BaseDto struct{
 CreatedAt time.Time
 UpdatedAt time.Time

}
type StudentDto struct{
    Name string
    BaseDto `mapper:"base"`
}

type StudentModel struct{
    Name string
    gorm.Model `mapper:"base"`
}

source:= Student{
    Name:"MyName",
    Model:gorm.Model{
        CreatedAt:time.Now(),
        UpdatedAt:time.Now(),
    },
}

to:=&StudentDto{}
mapper.Map(source,to)

```

## example struct
```
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

```
## example for map to struct
``` 
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

```
## example for map to Slice
``` 
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

```