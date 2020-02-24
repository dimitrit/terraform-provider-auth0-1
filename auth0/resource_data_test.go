package auth0

import (
	"strconv"
	"testing"
	"time"
)

func TestMapData(t *testing.T) {
	d := MapData{
		"one":  1,
		"zero": 0,
	}

	for key, shouldBeOk := range map[string]bool{
		"one":  true,
		"zero": false,
	} {
		if _, ok := d.GetOkExists(key); ok != shouldBeOk {
			t.Errorf("d.GetOkExists(%s) should retport ok == %t", key, shouldBeOk)
		}
	}
}

func TestJSON(t *testing.T) {
	d := MapData{"json": `{"foo": 123}`}
	v, err := JSON(d, "json")
	if err != nil {
		t.Error(err)
	}
	j, ok := v["foo"]
	if !ok {
		t.Errorf("Expected result to be a int, instead it was %T\n", j)
	}
}

func TestInt(t *testing.T) {
	k := "some-key"

	t.Run("int", func(t *testing.T) {
		v := 42
		d := MapData(map[string]interface{}{
			k: v,
		})
		a := Int(d, k)
		if *a != v {
			t.Errorf("unexpected value %d", *a)
		}
	})

	t.Run("string", func(t *testing.T) {
		v := "123"
		d := MapData(map[string]interface{}{
			k: v,
		})
		a := Int(d, k)
		w, _ := strconv.Atoi(v)
		if *a != w {
			t.Errorf("unexpected value %d", *a)
		}
	})

	t.Run("not an int", func(t *testing.T) {
		d := MapData(map[string]interface{}{
			k: time.Now(),
		})
		a := Int(d, k)
		if a != nil {
			t.Errorf("unexpected value %v", *a)
		}
	})
}
