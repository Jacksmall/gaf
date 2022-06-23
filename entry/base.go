package entry

import "time"

type PageReq struct {
	// 通用参数--第几页 參數位置隨請求方式
	Page int `json:"page" form:"page" binding:"omitempty"`

	// 通用参数--每页数量 參數位置隨請求方式
	Limit int `json:"limit" form:"limit" binding:"omitempty,max=1000"`
}

// GetOffset 获取偏移量
func (s PageReq) GetOffset() int {
	return (s.GetPage() - 1) * s.GetLimit()
}

// GetPage 获取页码，不能小于0
func (s PageReq) GetPage() int {
	if s.Page <= 0 {
		return 1
	}
	return s.Page
}

// GetLimit 获取显示条数 不能小于0 不能大于1000
func (s PageReq) GetLimit() int {
	if s.Limit <= 0 {
		return 10
	} else if s.Limit > 1000 {
		return 1000
	} else {
		return s.Limit
	}
}

// PageRes 分页响应数据
type PageRes struct {
	// 总数 全量返回或者不需要数量得情况返回0
	Total int64 `json:"total"`
	// 列表数据
	List interface{} `json:"list"`
}

type BaseTimeField struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
