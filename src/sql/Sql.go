package sql

import ()

const T_USER_LOGIN_SQL string = "select id as Id,user_name as UserName,user_real_name as RealName,user_pwd as Pwd,role_id as RoleId from t_user where user_name=?"
const T_SUBJECT_SELECT_SQL string = "select t.id as Id,t.subCode as SubCode,t.subName as SubName,t.subRemark as SubRemark from subject t"
const T_SUBJECT_SELECTONE_SQL string = "select t.id as Id,t.subCode as SubCode,t.subName as SubName,t.subRemark as SubRemark from subject t where t.id=?"
const T_SUBJECT_INSERT_SQL string = "insert into subject(subCode,subName,subRemark) values (?,?,?)"
const T_SUBJECT_UPDATE_SQL string = "update subject set subCode=?,subName=?,subRemark=? where id=?"
const T_SUBJECT_DELETE_SQL string = "delete from subject where id=?"
const T_DETAIL_SELECT_SQL = "SELECT t.id as Id,t.`year` as Year,t.`month` as Month,t.`day` as Day,t.cerNumber as CerNumber,t.remark as Remark,t.borrow as Borrow,t.pay as Pay,t.bor_pay as Bor_Pay,t.balance as Balance,t.subId AS SubId FROM detail t"
const T_DETAIL_SELECTONE_SQL = "SELECT t.id as Id,t.`year` as Year,t.`month` as Month,t.`day` as Day,t.cerNumber as CerNumber,t.remark as Remark,t.borrow as Borrow,t.pay as Pay,t.bor_pay as Bor_Pay,t.balance as Balance,t.subId AS SubId FROM detail t where t.id =?"
const T_DETAIL_INSERT_SQL = "insert into detail (`year`,`month`,`day`,cerNumber,remark,borrow,pay,bor_pay,balance,subId) values (?,?,?,?,?,?,?,?,?,?)"
const T_DETAIL_UPDATE_SQL = "update detail set `year`=?,`month`=?,`day`=?,cerNumber=?,remark=?,borrow=?,pay=?,bor_pay=?,balance=?,subId=? where id=?"
const T_DETAIL_DELETE_SQL = "delete from detail where id=?"
