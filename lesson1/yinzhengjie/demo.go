package main

import(
	"fmt"
)

type province struct {
	id  int
	val string
}
type city struct {
	fid int
	id  int
	val string
}
type conty struct {
	fid int
	id  int
	val string
}

var (
	provinces = make(map[int]province)
	citys     = make(map[int]city)
	contys    = make(map[int]conty)
)

func main() {
	provinces[1] = province{
		id:  1,
		val: "山东",
	}
	provinces[2] = province{
		id:  2,
		val: "河北",
	}

	citys[10] = city{
		fid: provinces[1].id,
		id:  10,
		val: "青岛",
	}
	citys[11] = city{
		fid: provinces[1].id,
		id:  11,
		val: "烟台",
	}
	contys[100] = conty{
		fid: citys[1].id,
		id:  100,
		val: "平谷区",
	}
	contys[101] = conty{
		fid: citys[1].id,
		id:  101,
		val: "平谷区",
	}
	res := need_city(1)
	fmt.Println(res)
}
func need_city(provinc_id int) []int {
	var cit []int
	for k, v := range citys {
		
		if v.fid == provinc_id {
			cit = append(cit, v.id)
		}
	}
	return cit
}

// func need_conty(city_id) []int {
// 	return
// }


