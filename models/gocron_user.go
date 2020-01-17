package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type GocronUser struct {
	Id       int       `orm:"column(id);pk"`
	Name     string    `orm:"column(name)"`
	Password string    `orm:"column(password)"`
	Salt     string    `orm:"column(salt)"`
	Email    string    `orm:"column(email)"`
	Created  time.Time `orm:"column(created);type(timestamp without time zone)"`
	Updated  time.Time `orm:"column(updated);type(timestamp without time zone);null"`
	IsAdmin  int16     `orm:"column(is_admin)"`
	Status   int16     `orm:"column(status)"`
}

func (t *GocronUser) TableName() string {
	return "gocron_user"
}

func init() {
	orm.RegisterModel(new(GocronUser))
}

// AddGocronUser insert a new GocronUser into database and returns
// last inserted Id on success.
func AddGocronUser(m *GocronUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGocronUserById retrieves GocronUser by Id. Returns error if
// Id doesn't exist
func GetGocronUserById(id int) (v *GocronUser, err error) {
	o := orm.NewOrm()
	v = &GocronUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGocronUser retrieves all GocronUser matches certain condition. Returns empty list if
// no records exist
func GetAllGocronUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GocronUser))
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

	var l []GocronUser
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

// UpdateGocronUser updates GocronUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateGocronUserById(m *GocronUser) (err error) {
	o := orm.NewOrm()
	v := GocronUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGocronUser deletes GocronUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGocronUser(id int) (err error) {
	o := orm.NewOrm()
	v := GocronUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GocronUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
