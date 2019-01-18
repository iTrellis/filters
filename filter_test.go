// MIT License

// Copyright (c) 2016 go-trellis

package filters_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-trellis/filters"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	manager = filters.New()
)

func init() {
	manager.AddCompareFunc(FilterTestFuncEquals, FilterTestEquals)
	manager.AddCompareFunc(FilterTestFuncTestInt, FilterTestInt)
}

func TestTarget(t *testing.T) {

	Convey("get target name's demensions", t, func() {
		Convey("when xxxx not in map", func() {
			Convey("will be not filtered", func() {
				filtered, e := manager.Compare(
					&filters.FilterParams{Names: []string{"xxxx"}},
					nil, nil)
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)
			})
		})
		Convey("when FilterTestFuncEquals in map", func() {
			Convey("will get FilterTestFuncEquals' dememsions", func() {
				filtered, e := manager.Compare(
					&filters.FilterParams{
						Names: []string{FilterTestFuncEquals}},
					nil, nil)
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)
			})
		})
	})

	Convey("test sequence filter", t, func() {
		Convey("when filter test_key does not equal", func() {
			Convey("will pass the filter", func() {
				filtered, e := manager.Compare(
					&filters.FilterParams{
						Names: []string{FilterTestFuncEquals}},
					nil,
					filters.FilterValues{})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)

				filtered, e = manager.Compare(
					&filters.FilterParams{Names: []string{FilterTestFuncEquals}},
					filters.FilterValues{testKeyEquals: ""},
					filters.FilterValues{})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)

				filtered, e = manager.Compare(
					&filters.FilterParams{Names: []string{FilterTestFuncEquals}},
					filters.FilterValues{testKeyEquals: "1", testKeyInt: 0},
					filters.FilterValues{testKeyEquals: "2"})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)

				filtered, e = manager.Compare(
					&filters.FilterParams{Names: []string{FilterTestFuncEquals}},
					filters.FilterValues{testKeyEquals: 1, testKeyInt: 2},
					filters.FilterValues{testKeyEquals: 2})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)

				filtered, e = manager.Compare(
					&filters.FilterParams{Names: []string{FilterTestFuncEquals}},
					filters.FilterValues{testKeyEquals: 1, testKeyInt: 2},
					filters.FilterValues{testKeyEquals: 2})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)
			})
		})
		Convey("when filter values' key exist", func() {
			Convey("will be filtered", func() {
				filtered, e := manager.Compare(
					&filters.FilterParams{Names: []string{FilterTestFuncEquals}},
					filters.FilterValues{},
					filters.FilterValues{testKeyEquals: "1"})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeTrue)

				filtered, e = manager.Compare(
					&filters.FilterParams{Names: []string{FilterTestFuncEquals}},
					filters.FilterValues{testKeyEquals: "1"},
					filters.FilterValues{testKeyEquals: "1"})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeTrue)
			})
		})
		Convey("when filter has two more names", func() {
			Convey("will be filtered", func() {
				filtered, e := manager.Compare(
					&filters.FilterParams{Names: []string{
						FilterTestFuncEquals,
						FilterTestFuncTestInt}},
					filters.FilterValues{testKeyEquals: "2", testKeyInt: "1"},
					filters.FilterValues{testKeyEquals: "1"})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeTrue)

				filtered, e = manager.Compare(
					&filters.FilterParams{Names: []string{
						FilterTestFuncTestInt, FilterTestFuncEquals}},
					filters.FilterValues{testKeyEquals: "1", testKeyInt: 2},
					filters.FilterValues{testKeyEquals: "1"})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeTrue)
			})
		})
	})

	Convey("test sequence filter", t, func() {
		Convey("will be filtered", func() {
			filtered, e := manager.Compare(
				&filters.FilterParams{
					Type: filters.CompareTypeConsistent,
					Names: []string{
						FilterTestFuncEquals,
						FilterTestFuncTestInt}},
				filters.FilterValues{},
				filters.FilterValues{testKeyEquals: "1"})
			So(e, ShouldBeNil)
			So(filtered, ShouldBeTrue)

			filtered, e = manager.Compare(
				&filters.FilterParams{
					Type: filters.CompareTypeConsistent,
					Names: []string{
						FilterTestFuncEquals,
						FilterTestFuncTestInt}},
				filters.FilterValues{testKeyEquals: "1"},
				filters.FilterValues{testKeyEquals: "1"})
			So(e, ShouldBeNil)
			So(filtered, ShouldBeTrue)

			filtered, e = manager.Compare(
				&filters.FilterParams{
					Type: filters.CompareTypeConsistent,
					Names: []string{
						FilterTestFuncEquals,
						FilterTestFuncTestInt}},
				filters.FilterValues{testKeyEquals: 2, testKeyInt: 2},
				filters.FilterValues{testKeyEquals: 2})
			So(e, ShouldBeNil)
			So(filtered, ShouldBeTrue)

			filtered, e = manager.Compare(
				&filters.FilterParams{
					Type: filters.CompareTypeConsistent,
					Names: []string{
						FilterTestFuncEquals,
						FilterTestFuncTestInt}},
				filters.FilterValues{testKeyEquals: 1, testKeyInt: "2"},
				filters.FilterValues{testKeyEquals: 2})
			So(e, ShouldBeNil)
			So(filtered, ShouldBeTrue)
		})
	})

	Convey("test filter timeout", t, func() {
		Convey("set invalid timeout", func() {
			Convey("will return error", func() {

				e := manager.SetCompareTimeout(-1 * time.Second)
				So(e.Error(), ShouldEqual, filters.ErrTimeoutMustAboveZero.New().Error())
			})
		})
		manager.AddCompareFunc(FilterTestFuncTimeout, FilterTestTimeout)
		manager.SetCompareTimeout(time.Second)
		Convey("input timeout function and exec sequence function", func() {
			Convey("will timeout", func() {

				filtered, e := manager.Compare(
					&filters.FilterParams{Names: []string{FilterTestFuncTimeout}},
					filters.FilterValues{testKeyEquals: 1, testKeyInt: "2"},
					filters.FilterValues{testKeyEquals: 2})
				So(e.Error(), ShouldEqual, filters.ErrFailedExecTimeout.New().Error())
				So(filtered, ShouldBeTrue)
			})
		})
		Convey("input timeout function and exec consistent function", func() {
			Convey("will timeout", func() {

				filtered, e := manager.Compare(
					&filters.FilterParams{
						Type: filters.CompareTypeConsistent,
						Names: []string{
							FilterTestFuncTimeout,
							FilterTestFuncEquals,
							FilterTestFuncTestInt}},
					filters.FilterValues{testKeyEquals: 1, testKeyInt: 2},
					filters.FilterValues{testKeyEquals: 2})
				So(e.Error(), ShouldEqual, filters.ErrFailedExecTimeout.New().Error())
				So(filtered, ShouldBeTrue)
			})
		})
	})
	return
}

// FilterTestFuncs
const (
	FilterTestFuncEquals  = "defaultTestEquals"
	FilterTestFuncTestInt = "defaultTestInt"
	FilterTestFuncTimeout = "defaultTestTimeout"

	testKeyEquals = "test_key_equals"
	testKeyInt    = "test_key_int"
)

func FilterTestEquals(req filters.FilterValues, cvs filters.FilterValues) (bool, error) {
	if len(cvs) == 0 {
		return false, nil
	}

	if len(req) == 0 {
		return true, nil
	}

	return req[testKeyEquals] == cvs[testKeyEquals], nil
}

func FilterTestTimeout(_ filters.FilterValues, _ filters.FilterValues) (bool, error) {
	time.Sleep(time.Second * 3)
	return false, nil
}

func FilterTestInt(req filters.FilterValues, _ filters.FilterValues) (bool, error) {
	vs := req[testKeyInt]
	if vs == nil {
		return true, nil
	}

	switch reflect.TypeOf(vs).Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		return false, nil
	}

	return true, nil
}
