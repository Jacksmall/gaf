package models

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"time"
)

var _db *gorm.DB

func SetDB(db *gorm.DB) {
	_db = db
}

type Where struct {
	Field string
	Op    string
	Value interface{}
}

func (w Where) toWhere(query *gorm.DB) *gorm.DB {
	switch strings.ToUpper(w.Op) {
	case "IN":
		// 可能会出现uint8 的类型
		if uv, ok := w.Value.([]uint8); ok {
			nv := make([]int, 0)
			for _, v := range uv {
				nv = append(nv, int(v))
			}
			query = query.Where(fmt.Sprintf("%s IN (?)", w.Field), nv)
		} else {
			query = query.Where(fmt.Sprintf("%s IN (?)", w.Field), w.Value)
		}
	case "LIKE":
	case "NOT LIKE":
		if strings.HasPrefix(w.Value.(string), "%") || strings.HasSuffix(w.Value.(string), "%") {
			query = query.Where(fmt.Sprintf("%s %s ?", w.Field, w.Op), w.Value.(string))
		} else {
			query = query.Where(fmt.Sprintf("%s %s ?", w.Field, w.Op), "%"+w.Value.(string)+"%")
		}
	case "BETWEEN":
		v := w.Value.([]interface{})
		query = query.Where(fmt.Sprintf("%s BETWEEN ? AND ?", w.Field), v[0], v[1])
	case "FIND_IN_SET":
		query = query.Where(fmt.Sprintf("FIND_IN_SET(?, %s)", w.Field), w.Value)
	case "RAW":
		query = query.Where(w.Field, w.Value.([]interface{})...)
	default:
		query = query.Where(fmt.Sprintf("%s %s ?", w.Field, w.Op), w.Value)
	}
	return query
}

// 任何实现该接口的结构体都可以作为M()函数的第一个参数
type table interface {
	TableName() string
}

type Common struct {
	t  table
	db *gorm.DB
}

// M NewCommon 函数
func M(t table, tx ...*gorm.DB) Common {
	db := _db
	if len(tx) > 0 {
		db = tx[0]
	}
	return Common{t, db}
}

// ToWhere 构建查询query DB
func ToWhere(query *gorm.DB, wheres []Where) *gorm.DB {
	if wheres == nil {
		return query
	}
	for _, w := range wheres {
		query = w.toWhere(query)
	}
	return query
}

// First 获取指定排序的第一条记录到target
func (c Common) First(where []Where, order string, target interface{}) error {
	return ToWhere(c.db.Table(c.t.TableName()), where).Order(order).First(target).Error
}

// Find 查找多条记录
func (c Common) Find(where []Where, order string, target interface{}) error {
	return ToWhere(c.db.Table(c.t.TableName()), where).Order(order).Find(target).Error
}

// Insert t table 就是实现table函数的model 结构体
func (c Common) Insert(t table) error {
	return c.db.Create(t).Error
}

// Save t table 保存所有字段
func (c Common) Save(t table) error {
	return c.db.Save(t).Error
}

// Update t table 更新多列
func (c Common) Update(where []Where, data map[string]interface{}) (int64, error) {
	obj := reflect.TypeOf(c.t)
	if _, exists := obj.FieldByName("UpdatedAt"); exists {
		if data["updated_at"] == nil || reflect.ValueOf(data["updated_at"]).IsZero() {
			data["updated_at"] = time.Now().Unix()
		}
	}
	res := ToWhere(c.db.Table(c.t.TableName()), where).UpdateColumns(data)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

// Delete 删除
func (c Common) Delete(where []Where) (int64, error) {
	res := ToWhere(c.db.Table(c.t.TableName()), where).Delete(c.t)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

// SoftDelete 软删除
func (c Common) SoftDelete(where []Where) (int64, error) {
	timestamp := time.Now().Unix()
	return c.Update(where, map[string]interface{}{
		"delete_at": timestamp,
	})
}

// ListPageWithCount 获取分页列表和总数一起
func (c Common) ListPageWithCount(where []Where, offset, limit int, order string, target interface{}) (count int64, err error) {
	query := ToWhere(_db.Table(c.t.TableName()), where)
	err = query.Count(&count).Error
	if count > 0 {
		err = query.Offset(offset).Limit(limit).Order(order).Find(target).Error
	}
	return
}
