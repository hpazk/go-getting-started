[![](https://img.shields.io/github/license/golangkit/formatime)](https://img.shields.io/github/license/golangkit/formatime)
[![GoDoc](https://godoc.org/github.com/golangkit/formatime?status.svg)](https://godoc.org/github.com/golangkit/formatime)

# formatime

自定义golang时间的格式，能与gorm结合使用，无需再使用sql.nullTime，并能在struct的json序列化中返回指定的时间格式

## 安装
```bash
go get github.com/golangkit/formatime
```

## 使用
```go

type foo struct {
	CreatedAt formatime.Date `gorm:"column:created_at" json:"created_at"`
	UpdatedAt formatime.Timestamp `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt formatime.CustomTime `gorm:"column:deleted_at" json:"deleted_at"`
}

func Test_datetime(t *testing.T) {
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
```

```javascript
// log结果如下
{
    "created_at":"2019-10-29",
    "updated_at":"1572312916",
    "deleted_at":"Oct 29 09:35:16"
}
```
