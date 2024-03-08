package main

import (
	"context"
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/fireyun/gorm-gen-demo/gendemo/dal"
	"github.com/fireyun/gorm-gen-demo/gendemo/dal/query"
)

func main() {
	ops1()
}

func ops1() {
	var que = query.Use(dal.DB.Debug())
	ctx := context.Background()
	m := que.TemplateSets
	q := que.TemplateSets.WithContext(ctx)

	r, err := q.Where(m.ID.Eq(1)).First()
	fmt.Println(r, err)

	// q.Where(m.ID.Eq(1)).Update(m.TemplateIDs, gorm.Expr("JSON_ARRAY_APPEND(template_ids, '$', ?)", 9))
	rs, err := q.Where(m.ID.Eq(1)).Not(gen.Cond(datatypes.JSONArrayQuery("template_ids").Contains(1))...).
		Update(m.TemplateIDs, gorm.Expr("JSON_ARRAY_APPEND(template_ids, '$', ?)", 1))
	fmt.Printf("rs:%#v, err:%v\n", rs, err)

	rs2, err := q.Where(gen.Cond(datatypes.JSONArrayQuery("bound_apps").Contains(5))...).Find()
	fmt.Println(len(rs2), err)
	for _, r := range rs2 {
		fmt.Printf("id: %d, name: %s, template_ids: %v\n", r.ID, r.Spec.Name, r.Spec.BoundApps)
	}

	tmplID := 7
	// subQuery get the array of template ids after delete the target template id, set it to '[]' if no records found
	subQuery := "COALESCE ((SELECT JSON_ARRAYAGG(oid) new_oids FROM " +
		"JSON_TABLE (template_ids, '$[*]' COLUMNS (oid BIGINT (1) UNSIGNED PATH '$')) AS t1 WHERE oid<> ?), '[]')"
	rs3, err := q.Where(m.BizID.Eq(2)).
		Where(gen.Cond(datatypes.JSONArrayQuery("template_ids").Contains(tmplID))...).
		Update(m.TemplateIDs, gorm.Expr(subQuery, tmplID))
	fmt.Printf("rs3:%#v, err:%v\n", rs3, err)

}
