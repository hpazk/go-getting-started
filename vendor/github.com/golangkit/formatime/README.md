[![](https://img.shields.io/github/license/golangkit/formatime)](https://img.shields.io/github/license/golangkit/formatime)
[![GoDoc](https://godoc.org/github.com/golangkit/formatime?status.svg)](https://godoc.org/github.com/golangkit/formatime)

# formatime

[中文](./README.cn.md)

Customize the format of the time in golang, can be used in conjunction with the gorm, no need to use sql.nullTime, and can return the specified time format in the json serialization of the struct
## Installation
```bash
go get github.com/golangkit/formatime
```

## Usage 
```go

type foo struct {
    ID        int64              `gorm:"column:id"`
	CreatedAt formatime.Date `gorm:"column:created_at" json:"created_at"`
	UpdatedAt formatime.Timestamp `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt formatime.CustomTime `gorm:"column:deleted_at" json:"deleted_at"`
}

func Test_json(t *testing.T) {
	b := foo{
		CreatedAt: formatime.NewDateNow(),
		UpdatedAt: formatime.NewTimestampNow(),
		DeletedAt: formatime.NewCustomTimeNow("Jan _2 15:04:05"),
	}

	text, err := json.Marshal(&b)
	if err != nil {
		t.Fatal(err)
	} else {
		log.Print(string(text))
	}
}

func Test_gorm(t *testing.T) {
	b := &foo{}
    	gorm.DB.First(b)
	text, err := json.Marshal(&b)
	if err != nil {
		t.Fatal(err)
	} else {
		log.Print(string(text))
	}
}
```

```javascript
// log 
{
    "created_at":"2019-10-29",
    "updated_at":"1572312916",
    "deleted_at":"Oct 29 09:35:16"
}
```
