# target-manager


## Build

* [![Build Status](https://travis-ci.org/go-trellis/filters.png)](https://travis-ci.org/go-trellis/filters)

## Sample

### Test Config

[target.sample.json](target.sample.json)

```json
{
  "target_demensions": [
    {
      "target_name": "test1",
      "demensions": []
    },
    {
      "target_name": "test2",
      "demensions": [
        {
          "target_key": "test2-key1",
          "description": "for-test2-key1"
        },
        {
          "target_key": "test2-key2",
          "description": "for-test2-key2"
        }
      ]
    }
  ]
}
```

### Init Config

```go
    filter, err := filters.NewManager(
	    filters.InitFiltersTypeFromFile, map[string]interface{}{"filename": "target.json"})
```


### TEST

[manager_test.go](manager_test.go)

```go
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
```
