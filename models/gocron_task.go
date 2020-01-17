package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type GocronTask struct {
	Id               int       `orm:"column(id);pk"`
	Name             string    `orm:"column(name)"`
	Level            int16     `orm:"column(level)"`
	DependencyTaskId string    `orm:"column(dependency_task_id)"`
	DependencyStatus int16     `orm:"column(dependency_status)"`
	Spec             string    `orm:"column(spec)"`
	Protocol         int16     `orm:"column(protocol)"`
	Command          string    `orm:"column(command)"`
	HttpMethod       int16     `orm:"column(http_method)"`
	Timeout          int       `orm:"column(timeout)"`
	Multi            int16     `orm:"column(multi)"`
	RetryTimes       int16     `orm:"column(retry_times)"`
	RetryInterval    int16     `orm:"column(retry_interval)"`
	NotifyStatus     int16     `orm:"column(notify_status)"`
	NotifyType       int16     `orm:"column(notify_type)"`
	NotifyReceiverId string    `orm:"column(notify_receiver_id)"`
	NotifyKeyword    string    `orm:"column(notify_keyword)"`
	Tag              string    `orm:"column(tag)"`
	Remark           string    `orm:"column(remark)"`
	Status           int16     `orm:"column(status)"`
	Created          time.Time `orm:"column(created);type(timestamp without time zone)"`
	Deleted          time.Time `orm:"column(deleted);type(timestamp without time zone);null"`
}

func (t *GocronTask) TableName() string {
	return "gocron_task"
}

func init() {
	orm.RegisterModel(new(GocronTask))
}

// AddGocronTask insert a new GocronTask into database and returns
// last inserted Id on success.
func AddGocronTask(m *GocronTask) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGocronTaskById retrieves GocronTask by Id. Returns error if
// Id doesn't exist
func GetGocronTaskById(id int) (v *GocronTask, err error) {
	o := orm.NewOrm()
	v = &GocronTask{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGocronTask retrieves all GocronTask matches certain condition. Returns empty list if
// no records exist
func GetAllGocronTask(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GocronTask))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []GocronTask
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateGocronTask updates GocronTask by Id and returns error if
// the record to be updated doesn't exist
func UpdateGocronTaskById(m *GocronTask) (err error) {
	o := orm.NewOrm()
	v := GocronTask{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGocronTask deletes GocronTask by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGocronTask(id int) (err error) {
	o := orm.NewOrm()
	v := GocronTask{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GocronTask{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
