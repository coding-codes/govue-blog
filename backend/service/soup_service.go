package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coding-codes/tools"
	"log"
	"strconv"
)

type Soup struct {
	ID      int    `json:"id" db:"id"`
	Content string `json:"content" db:"content" binding:"required"`
}

type Soups struct {
	Items []Soup `json:"items"`
	Total int    `json:"total"`
}

func (s *Soup) Create() (Soup, error) {
	r, e := db.Exec("insert into blog_soup (content) values (?)", s.Content)
	if e != nil {
		return Soup{}, e
	}
	id, _ := r.LastInsertId()
	return Soup{int(id), s.Content}, nil
}

func (s *Soup) Delete() error {
	_, e := db.Exec("delete from blog_soup where id=?", s.ID)
	return e
}

func (s *Soup) Edit() error {
	_, e := db.Exec("update blog_soup set content=? where id=?", s.Content, s.ID)
	return e
}

func (s *Soup) GetOne() (Soup, error) {
	var soup Soup
	e := db.Get(&soup, "select * from blog_soup where id=?", s.ID)
	return soup, e
}

func (s Soup) GetRandOne() (Soup, error) {
	var Soup Soup
	e := db.Get(&Soup, `SELECT t1.* FROM blog_soup AS t1 JOIN (SELECT ROUND(RAND() * ((SELECT MAX(id) FROM blog_soup) - (SELECT MIN(id) FROM blog_soup)) + (SELECT MIN(id) FROM blog_soup)) AS id
) AS t2 WHERE t1.id >= t2.id LIMIT 1`)
	return Soup, e
}

func (s Soup) GetAll(limit, page string) (data Soups, err error) {
	baseSql := "select %s from blog_soup s"
	var p, l int
	// 把字符串转换为 int ,并设置 redis 的 key name
	if limit != "" && page != "" {
		p, _ = strconv.Atoi(page)
		l, _ = strconv.Atoi(limit)
	}
	key := fmt.Sprintf("soup_%d_%d", l, p)

	// 如有有cache，从cache 读取，然后退出
	cacheData, e := getSoupCache(key)
	if e != nil {
		log.Fatal("read cache error")
	}
	if cacheData.Total != 0 {
		return cacheData, nil
	}

	// 如果没有 cache 就从数据库读取，然后设置 cache
	soups := make([]Soup, 0)
	offset := (p - 1) * l
	selectSql := fmt.Sprintf(baseSql, "s.*") + fmt.Sprintf(" ORDER BY s.id DESC limit %d offset %d", l, offset)
	if err = db.Select(&soups, selectSql); err != nil {
		return
	}
	var total int
	if err = db.Get(&total, fmt.Sprintf(baseSql, "count(1)")); err != nil {
		return
	}

	data.Total = total
	data.Items = soups

	if e := setSoupCache(key, data); e != nil {
		log.Fatal("write cache error")
	}

	return
}

//func soupCacheKey(limit, page int) string {
//	if limit == 0 || page == 0 {
//		return fmt.Sprintf("soup_%d_%d", 10, 1)
//	}
//	return fmt.Sprintf("soup_%d_%d", limit, page)
//}

func getSoupCache(key string) (s Soups, err error) {
	data, e := tools.GetKey(key)
	if e != nil || data == nil {
		return s, e
	}

	v, ok := data.([]uint8)
	fmt.Println(v)
	if ok {
		if e := json.Unmarshal([]byte(v[:]), &s); e != nil {
			return s, e
		}
		return s, nil
	} else {
		return s, errors.New("返回数据类型有误，json无法解析")
	}
}

func setSoupCache(key string, value Soups) error {
	marshal, _ := json.Marshal(value)
	e := tools.SetKey(key, marshal, tools.SetTimeout(true))
	return e
}
