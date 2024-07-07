package system

import (
	"lixuefei.com/go-admin/common/request"
	"lixuefei.com/go-admin/common/response"
	"lixuefei.com/go-admin/internal/dao"
)

type SysUserService struct {
}

func (s *SysUserService) QueryUserList(pageInfo request.PageInfo) *response.PageResult {
	pageNo := pageInfo.Page
	pageSize := pageInfo.PageSize
	keyword := pageInfo.Keyword
	userList, total, err := dao.SysUserDao{}.QueryUserList(pageNo, pageSize, keyword)
	if err != nil {
		return nil
	}

	pageResult := response.PageResult{
		Page:     pageNo,
		PageSize: pageSize,
		Total:    total,
		List:     userList,
	}
	return &pageResult
}
