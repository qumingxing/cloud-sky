package service

import (
	"dao"
	"entity"
	"sql"
	"common"
)

type DetailService struct {
}

func (detailService *DetailService) GetDetailListPage(pageInfo *common.PageInfo, detail *entity.Detail) (*common.PageInfo, []*entity.Detail, error) {
	var baseDao dao.BaseDao
	info, data, err := baseDao.SelectSplitPageStruct(pageInfo, detail, sql.T_DETAIL_SELECT_SQL)
	if err != nil {
		return nil, nil, err
	}
	dataList := convertToDetail(data)
	return info, dataList, nil
}
func (detailService *DetailService) SaveDetail(detail *entity.Detail) int {
	var baseDao dao.BaseDao
	id, _ := baseDao.Add(sql.T_DETAIL_INSERT_SQL, detail.Year, detail.Month, detail.Day, detail.CerNumber, detail.Remark, detail.Borrow, detail.Pay, detail.Bor_Pay, detail.Balance, detail.SubId)
	return id
}
func (detailService *DetailService) UpdateDetail(detail *entity.Detail) int {
	var baseDao dao.BaseDao
	count, _ := baseDao.Update(sql.T_DETAIL_UPDATE_SQL, detail.Year, detail.Month, detail.Day, detail.CerNumber, detail.Remark, detail.Borrow, detail.Pay, detail.Bor_Pay, detail.Balance, detail.SubId, detail.Id)
	return count
}
func (detailService *DetailService) DeleteDetail(detail *entity.Detail) int {
	var baseDao dao.BaseDao
	count, _ := baseDao.Delete(sql.T_DETAIL_DELETE_SQL, detail.Id)
	return count
}
func (detailService *DetailService) GetDetail(detail *entity.Detail) *entity.Detail {
	var baseDao dao.BaseDao
	obj, _ := baseDao.SelectOneStruct(detail, sql.T_DETAIL_SELECTONE_SQL, detail.Id)
	if v, ok := obj.(*entity.Detail); ok {
		return v
	}
	return nil
}
func convertToDetail(data []interface{}) []*entity.Detail {
	subList := make([]*entity.Detail, len(data))
	for i, obj := range data {
		if v, ok := obj.(*entity.Detail); ok {
			subList[i] = v
		}
	}
	return subList
}
