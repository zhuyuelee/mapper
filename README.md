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

example
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