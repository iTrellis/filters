// MIT License

// Copyright (c) 2015 rutcode-go

package filters_test

import (
	"testing"

	"github.com/go-trellis/filters"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	filter filters.FilterRepo
)

func TestInitTargetDemensions(t *testing.T) {
	Convey("initial filters with filename", t, func() {
		var err error
		Convey("when initial type not exists", func() {
			Convey("will get error", func() {
				filter, err = filters.NewManager("test", nil)
				So(err, ShouldEqual, filters.ErrNotFoundInitialFunction)
			})
		})
		Convey("when options nil", func() {
			Convey("will get error", func() {
				filter, err = filters.NewManager(filters.InitFiltersTypeFromFile, nil)
				So(err, ShouldEqual, filters.ErrNeedConfigFile)

				filter, err = filters.NewManager(
					filters.InitFiltersTypeFromFile, map[string]interface{}{"filename": 1})
				So(err, ShouldNotBeNil)
			})
		})
		Convey("when filename not exists", func() {
			Convey("will get error", func() {
				filter, err = filters.NewManager(
					filters.InitFiltersTypeFromFile, map[string]interface{}{"filename": "target.json"})
				So(err, ShouldNotBeNil)
			})
		})
		Convey("when filename exists", func() {
			Convey("will be initialed", func() {
				filter, err = filters.NewManager(
					filters.InitFiltersTypeFromFile, map[string]interface{}{"filename": "target.sample.json"})
				So(err, ShouldBeNil)
				So(filter, ShouldNotBeNil)
			})
		})
	})
	return
}

func TestTarget(t *testing.T) {

	Convey("get target name's demensions", t, func() {
		Convey("when xxxx not in map", func() {
			Convey("will xxxx not found", func() {
				filtered, e := filter.Compare("xxxx", nil, nil)
				So(e, ShouldNotBeNil)
				So(e, ShouldEqual, filters.ErrTargetNameNotExists)
				So(filtered, ShouldBeTrue)
			})
		})
		Convey("when test1 in map", func() {
			Convey("will get test1's dememsions", func() {
				_, e := filter.Compare("test1", nil, nil)
				So(e, ShouldBeNil)
			})
		})
	})

	Convey("target's dememsions not exist", t, func() {
		Convey("when target values is not nil", func() {
			Convey("will get filtered false", func() {
				filtered, e := filter.Compare("test1", filters.TargetValues{"": ""}, nil)
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)
				filtered, e = filter.Compare("test2", filters.TargetValues{"test2-key1": "1"}, nil)
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)
			})
		})
		Convey("when compare values is not nil", func() {
			Convey("will get test1's dememsions", func() {
				filtered, e := filter.Compare("test1", nil, filters.CompareValues{"": ""})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)
			})
		})
	})

	Convey("target's dememsions exist", t, func() {
		Convey("when compare values' key not exist", func() {
			Convey("will get error: invalid dememsion", func() {
				filtered, e := filter.Compare("test2", nil, filters.CompareValues{"": ""})
				So(e, ShouldNotBeNil)
				So(e, ShouldEqual, filters.ErrInvalidDemension)
				So(filtered, ShouldBeTrue)

				filtered, e = filter.Compare("test2", filters.TargetValues{"": ""}, filters.CompareValues{"": ""})
				So(e, ShouldNotBeNil)
				So(e, ShouldEqual, filters.ErrInvalidDemension)
				So(filtered, ShouldBeTrue)
			})
		})
		Convey("when traget values not equal compare values", func() {
			Convey("will get filtered true", func() {
				filtered, e := filter.Compare("test2",
					filters.TargetValues{"": ""},
					filters.CompareValues{"test2-key1": ""})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeTrue)

				filtered, e = filter.Compare("test2",
					filters.TargetValues{"test2-key1": 1},
					filters.CompareValues{"test2-key1": ""})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeTrue)

				filtered, e = filter.Compare("test2",
					filters.TargetValues{"test2-key1": "1"},
					filters.CompareValues{"test2-key1": "1",
						"test2-key2": 0})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeTrue)
			})
		})
		Convey("when traget values equals compare values", func() {
			Convey("will get filtered true", func() {
				filtered, e := filter.Compare("test2",
					filters.TargetValues{"test2-key1": ""},
					filters.CompareValues{"test2-key1": ""})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)

				filtered, e = filter.Compare("test2",
					filters.TargetValues{"test2-key1": 1},
					filters.CompareValues{"test2-key1": 1})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)

				filtered, e = filter.Compare("test2",
					filters.TargetValues{"test2-key1": "1", "test2-key2": 0},
					filters.CompareValues{"test2-key1": "1", "test2-key2": 0})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)
			})
		})
	})
	return
}
