package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/fireyun/gorm-gen-demo/gendemo/dal"
	"github.com/fireyun/gorm-gen-demo/gendemo/dal/query"
)

func TestEq(t *testing.T) {
	var que = query.Use(dal.DB)
	ctx := context.Background()
	m := que.App
	q := que.App.WithContext(ctx)

	r, err := q.Where(m.BizID.Eq(2), m.Name.Eq("cmdb-file001")).Find()
	fmt.Println(r, len(r), err)
}

func TestLike(t *testing.T) {
	var que = query.Use(dal.DB)
	ctx := context.Background()
	m := que.App
	q := que.App.WithContext(ctx)

	r, err := q.Where(m.BizID.Eq(2), m.Name.Like("%cmdb%")).Find()
	fmt.Println(r, len(r), err)
}

func TestRegex(t *testing.T) {
	var que = query.Use(dal.DB)
	ctx := context.Background()
	m := que.App
	q := que.App.WithContext(ctx)

	r, err := q.Where(m.BizID.Eq(2), m.Name.Regexp("(?i)cmdb")).Find()
	fmt.Println(r, len(r), err)
}

func BenchmarkEq(b *testing.B) {
	var que = query.Use(dal.DB)
	ctx := context.Background()
	m := que.App
	q := que.App.WithContext(ctx)
	for i := 0; i < b.N; i++ {
		q.Where(m.BizID.Eq(2), m.Name.Eq("cmdb-file001")).Find()
	}
}

func BenchmarkLike(b *testing.B) {
	var que = query.Use(dal.DB)
	ctx := context.Background()
	m := que.App
	q := que.App.WithContext(ctx)
	for i := 0; i < b.N; i++ {
		q.Where(m.BizID.Eq(2), m.Name.Like("%cmdb%")).Find()
	}
}

func BenchmarkRegex(b *testing.B) {
	var que = query.Use(dal.DB)
	ctx := context.Background()
	m := que.App
	q := que.App.WithContext(ctx)
	for i := 0; i < b.N; i++ {
		q.Where(m.BizID.Eq(2), m.Name.Regexp("(?i)cmdb")).Find()
	}
}
