package target_manager_test

import (
	"testing"

	tm "github.com/go-rut/target_manager"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	CONFIG_FILE = "target.conf.sample"
)

var (
	manager *tm.Manager
)

func TestInitTargetDemensions(t *testing.T) {
	manager := tm.NewManager()
	manager.InitTargetDemensions(CONFIG_FILE)
	return
}

func TestTarget(t *testing.T) {

	Convey("get target name's demensions", t, func() {
		Convey("when xxxx not in map", func() {
			Convey("will xxxx not found", func() {
				filtered, e := manager.Compare("xxxx", nil, nil)
				So(e, ShouldNotBeNil)
				So(e, ShouldEqual, tm.ERR_TARGET_NAME_NOT_EXIST)
				So(filtered, ShouldBeTrue)
			})
		})
		Convey("when test1 in map", func() {
			Convey("will get test1's dememsions", func() {
				_, e := manager.Compare("test1", nil, nil)
				So(e, ShouldBeNil)
			})
		})
	})

	Convey("target's dememsions not exist", t, func() {
		Convey("when target values is not nil", func() {
			Convey("will get filtered false", func() {
				filtered, e := manager.Compare("test1", tm.TargetValues{"": ""}, nil)
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)
			})
		})
		Convey("when compare values is not nil", func() {
			Convey("will get test1's dememsions", func() {
				filtered, e := manager.Compare("test1", nil, tm.CompareValues{"": ""})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)
			})
		})
	})

	Convey("target's dememsions exist", t, func() {
		Convey("when compare values' key not exist", func() {
			Convey("will get error: invalid dememsion", func() {
				filtered, e := manager.Compare("test2", nil, tm.CompareValues{"": ""})
				So(e, ShouldNotBeNil)
				So(e, ShouldEqual, tm.ERR_INVALID_DEMEMSION)
				So(filtered, ShouldBeTrue)

				filtered, e = manager.Compare("test2", tm.TargetValues{"": ""}, tm.CompareValues{"": ""})
				So(e, ShouldNotBeNil)
				So(e, ShouldEqual, tm.ERR_INVALID_DEMEMSION)
				So(filtered, ShouldBeTrue)
			})
		})
		Convey("when traget values not equal compare values", func() {
			Convey("will get filtered true", func() {
				filtered, e := manager.Compare("test2",
					tm.TargetValues{"": ""},
					tm.CompareValues{"test2-key1": ""})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeTrue)

				filtered, e = manager.Compare("test2",
					tm.TargetValues{"test2-key1": 1},
					tm.CompareValues{"test2-key1": ""})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeTrue)

				filtered, e = manager.Compare("test2",
					tm.TargetValues{"test2-key1": "1"},
					tm.CompareValues{"test2-key1": "1",
						"test2-key2": 0})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeTrue)
			})
		})
		Convey("when traget values equals compare values", func() {
			Convey("will get filtered true", func() {
				filtered, e := manager.Compare("test2",
					tm.TargetValues{"test2-key1": ""},
					tm.CompareValues{"test2-key1": ""})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)

				filtered, e = manager.Compare("test2",
					tm.TargetValues{"test2-key1": 1},
					tm.CompareValues{"test2-key1": 1})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)

				filtered, e = manager.Compare("test2",
					tm.TargetValues{"test2-key1": "1", "test2-key2": 0},
					tm.CompareValues{"test2-key1": "1", "test2-key2": 0})
				So(e, ShouldBeNil)
				So(filtered, ShouldBeFalse)
			})
		})
	})
	return
}
