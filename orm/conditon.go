/**
* @Description: 构建条件语句
* @Author: jinyidong
* @Date: 2021/6/19
* @Version V1.0
 */
package orm

import "strings"

type Condition interface {
	SetWhere(k string, v []interface{})
	SetOr(k string, v []interface{})
	SetOrder(k string)
	SetJoinOn(t, on string) Condition
}

type gormCondition struct {
	gormPublic
	Join []*gormJoin
}

type gormPublic struct {
	Where map[string][]interface{}
	Order []string
	Or    map[string][]interface{}
}

type gormJoin struct {
	Type   string
	JoinOn string
	gormPublic
}

func (e *gormJoin) SetJoinOn(t, on string) Condition {
	return nil
}

func (e *gormPublic) SetWhere(k string, v []interface{}) {
	if e.Where == nil {
		e.Where = make(map[string][]interface{})
	}
	e.Where[k] = v
}

func (e *gormPublic) SetOr(k string, v []interface{}) {
	if e.Or == nil {
		e.Or = make(map[string][]interface{})
	}
	e.Or[k] = v
}

func (e *gormPublic) SetOrder(k string) {
	if e.Order == nil {
		e.Order = make([]string, 0)
	}
	e.Order = append(e.Order, k)
}

func (e *gormCondition) SetJoinOn(t, on string) Condition {
	if e.Join == nil {
		e.Join = make([]*gormJoin, 0)
	}
	join := &gormJoin{
		Type:       t,
		JoinOn:     on,
		gormPublic: gormPublic{},
	}
	e.Join = append(e.Join, join)
	return join
}

type resolveSearchTag struct {
	Type   string
	Column string
	Table  string
	On     []string
	Join   string
}

// makeTag 解析search的tag标签
func makeTag(tag string) *resolveSearchTag {
	r := &resolveSearchTag{}
	tags := strings.Split(tag, ";")
	var ts []string
	for _, t := range tags {
		ts = strings.Split(t, ":")
		if len(ts) == 0 {
			continue
		}
		switch ts[0] {
		case "type":
			if len(ts) > 1 {
				r.Type = ts[1]
			}
		case "column":
			if len(ts) > 1 {
				r.Column = ts[1]
			}
		case "table":
			if len(ts) > 1 {
				r.Table = ts[1]
			}
		case "on":
			if len(ts) > 1 {
				r.On = ts[1:]
			}
		case "join":
			if len(ts) > 1 {
				r.Join = ts[1]
			}
		}
	}
	return r
}
