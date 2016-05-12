package service

import (
	"common"
	"dao"
	"entity"
	"sql"
)

type SubjectService struct {
}

func (subjectService *SubjectService) GetSubjectListPage(pageInfo *common.PageInfo, subject *entity.Subject) (*common.PageInfo, []*entity.Subject, error) {
	var baseDao dao.BaseDao
	info, data, err := baseDao.SelectSplitPageStruct(pageInfo, subject, sql.T_SUBJECT_SELECT_SQL)
	if err != nil {
		return nil, nil, err
	}
	dataList := convertToSubject(data)
	return info, dataList, nil
}
func (subjectService *SubjectService) GetSubjectList(subject *entity.Subject) ([]*entity.Subject, error) {
	var baseDao dao.BaseDao
	data, err := baseDao.SelectMultStruct(subject, sql.T_SUBJECT_SELECT_SQL)
	if err != nil {
		return nil, err
	}
	dataList := convertToSubject(data)
	return dataList, nil
}
func (subjectService *SubjectService) SaveSubject(subject *entity.Subject) int {
	var baseDao dao.BaseDao
	id, _ := baseDao.Add(sql.T_SUBJECT_INSERT_SQL, subject.SubCode, subject.SubName, subject.SubRemark)
	return id
}
func (subjectService *SubjectService) UpdateSubject(subject *entity.Subject) int {
	var baseDao dao.BaseDao
	count, _ := baseDao.Update(sql.T_SUBJECT_UPDATE_SQL, subject.SubCode, subject.SubName, subject.SubRemark, subject.Id)
	return count
}
func (subjectService *SubjectService) DeleteSubject(subject *entity.Subject) int {
	var baseDao dao.BaseDao
	count, _ := baseDao.Delete(sql.T_SUBJECT_DELETE_SQL, subject.Id)
	return count
}
func (subjectService *SubjectService) GetSubject(subject *entity.Subject) *entity.Subject {
	var baseDao dao.BaseDao
	obj, _ := baseDao.SelectOneStruct(subject, sql.T_SUBJECT_SELECTONE_SQL, subject.Id)
	if v, ok := obj.(*entity.Subject); ok {
		return v
	}
	return nil
}
func convertToSubject(data []interface{}) []*entity.Subject {
	subList := make([]*entity.Subject, len(data))
	for i, obj := range data {
		if v, ok := obj.(*entity.Subject); ok {
			subList[i] = v
		}
	}
	return subList
}
