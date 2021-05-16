package ll1

import (
	"strings"
)

type Empty struct{}

var empty Empty

type Set struct {
	m map[string]Empty
}

func SetFactory() *Set {
	return &Set{
		m: map[string]Empty{},
	}
}

func (s *Set) Union(val *Set) {
	if val != nil {
		for v := range val.m {
			s.m[v] = empty
		}
	}

}

//添加元素
func (s *Set) Add(val string) {
	s.m[val] = empty
}

//判断元素是否存在
func (s *Set) Get(val string) bool {
	for i := range s.m {
		if i == val {
			return true
		}
	}
	return false
}

//删除元素
func (s *Set) Remove(val string) {
	delete(s.m, val)
}

//获取长度
func (s *Set) Len() int {
	return len(s.m)
}

//清空set
func (s *Set) Clear() {
	s.m = make(map[string]Empty)
}

//遍历set
func (s *Set) Traverse() []string {
	var setSlice []string
	for v := range s.m {
		setSlice = append(setSlice, v)
	}
	return setSlice
}

func initInput(input string) ([]string, *Set, *Set) {
	arr := strings.Fields(input)

	noEndSet := SetFactory()
	yesEndSet := SetFactory()
	for _, tempString := range arr {

		tempByte := []byte(tempString)

		noEndSet.Add(string(tempByte[0]))
		for l := 3; l < len(tempString); l++ {
			if !(tempString[l] >= 'A' && tempString[l] <= 'Z') {
				if tempString[l] != '#' && tempString[l] != '|' {
					yesEndSet.Add(string(tempString[l]))
				}
			}

		}
	}
	yesEndSet.Add("$")
	return arr, noEndSet, yesEndSet

}

func ExuFirst(arr []string, noEndSet *Set) (map[string][]string, *Set) {

	first := make(map[string]*Set)
	firstTemp := make(map[string]*Set)

	//有空集的非终结符集合
	emptySET := SetFactory()

	for _, tempString := range arr {

		lenTemp := len(tempString)
		tempByte := []byte(tempString)

		if first[string(tempByte[0])] == nil {
			first[string(tempByte[0])] = SetFactory()
		}
		if firstTemp[string(tempByte[0])] == nil {
			firstTemp[string(tempByte[0])] = SetFactory()
		}

		if !(tempByte[3] >= 'A' && tempByte[3] <= 'Z') {
			first[string(tempByte[0])].Add(string(tempByte[3]))
		} else {
			firstTemp[string(tempByte[0])].Add(string(tempByte[3]))
		}

		if tempByte[3] == '#' {
			emptySET.Add(string(tempByte[0]))
		}

		for l := 3; l < lenTemp; l++ {
			if tempByte[l] == '#' {
				emptySET.Add(string(tempByte[0]))
			}
		}

	}

	for _, tempString := range arr {
		lenTemp := len(tempString)
		tempByte := []byte(tempString)
		flag := true
		for l := 3; l+1 < lenTemp; l++ {
			if emptySET.Get(string(tempByte[l])) && flag == true {
				if !(tempByte[l+1] >= 'A' && tempByte[l+1] <= 'Z') {
					first[string(tempByte[0])].Add(string(tempByte[l+1]))
				} else {
					firstTemp[string(tempByte[0])].Add(string(tempByte[l+1]))
				}
			} else {
				flag = false
			}
			if l < lenTemp && tempByte[l] == '|' {
				flag = true
				if !(tempByte[l+1] >= 'A' && tempByte[l+1] <= 'Z') {
					first[string(tempByte[0])].Add(string(tempByte[l+1]))
				} else {
					firstTemp[string(tempByte[0])].Add(string(tempByte[l+1]))
				}
			}
		}
	}

	noEndPoint := noEndSet.Traverse()

	for _, e := range noEndPoint {
		if firstTemp[e].Len() != 0 {
			for i, j := range firstTemp {
				set := j.Traverse()
				for _, k := range set {
					first[i].Union(first[k])
				}
			}
		}
	}

	FirstMap := make(map[string][]string)

	for _, j := range noEndPoint {
		FirstMap[j] = first[j].Traverse()
	}

	return FirstMap, emptySET
}

func ExuFollow(arr []string, noEndSet *Set, emptySet *Set, firstMap map[string][]string) map[string][]string {

	follow := make(map[string]*Set)
	followTemp := make(map[string]*Set)

	for _, tempString := range arr {

		tempByte := []byte(tempString)

		if follow[string(tempByte[0])] == nil {
			follow[string(tempByte[0])] = SetFactory()
		}
		if followTemp[string(tempByte[0])] == nil {
			followTemp[string(tempByte[0])] = SetFactory()
		}
	}
	follow[string(arr[0][0])].Add("$")

	for _, tempString := range arr {
		lenTemp := len(tempString)
		tempByte := []byte(tempString)

		for l := 3; l+1 < lenTemp; l++ {
			if noEndSet.Get(string(tempByte[l])) {
				if noEndSet.Get(string(tempByte[l+1])) {
					t := firstMap[string(tempByte[l+1])]
					for _, tt := range t {
						if tt != "#" {
							follow[string(tempByte[l])].Add(tt)
						}
					}
				} else if !noEndSet.Get(string(tempByte[l+1])) && string(tempByte[l+1]) != "|" && string(tempByte[l+1]) != "#" {
					follow[string(tempByte[l])].Add(string(tempByte[l+1]))
				}
				if emptySet.Get(string(tempByte[l+1])) {
					followTemp[string(tempByte[l])].Add(string(tempByte[0]))
				}
			}
		}

		for l := 3; l < lenTemp; l++ {
			if l+1 < lenTemp && string(tempByte[l+1]) == "|" && noEndSet.Get(string(tempByte[l])) {
				followTemp[string(tempByte[l])].Add(string(tempByte[0]))
			}
			if noEndSet.Get(string(tempByte[l])) && l == lenTemp-1 {
				followTemp[string(tempByte[l])].Add(string(tempByte[0]))
			}
		}

	}

	FollowMap := make(map[string][]string)

	noEndPoint := noEndSet.Traverse()

	for _, e := range noEndPoint {
		if followTemp[e].Len() != 0 {
			for i, j := range followTemp {
				set := j.Traverse()
				for _, k := range set {
					follow[i].Union(follow[k])
				}
			}
		}
	}

	for _, j := range noEndPoint {
		FollowMap[j] = follow[j].Traverse()
	}

	return FollowMap
}

func ExcTable(arr []string, noEndSet *Set, yesEndSet *Set, firstMap map[string][]string, followMap map[string][]string) (map[string]map[string]string, bool) {
	var newArr []string
	conflict := false
	for _, tempString := range arr {
		tempLen := len(tempString)
		var at []int
		for l := 3; l < tempLen; l++ {
			if tempString[l] == '|' {
				at = append(at, l)
			}
		}
		if len(at) == 0 {
			newArr = append(newArr, tempString)
		} else if len(at) == 1 {
			tempS := tempString[0:at[0]]
			newArr = append(newArr, tempS)
			tempS = tempString[0:3] + tempString[at[0]+1:]
			newArr = append(newArr, tempS)
		} else {
			tempS := tempString[0:at[0]]
			newArr = append(newArr, tempS)
			for l := 1; l < len(at); l++ {
				tempS := tempString[0:3] + tempString[at[l-1]+1:at[l]]
				newArr = append(newArr, tempS)
			}
			tempS = tempString[0:3] + tempString[at[len(at)-1]+1:]
			newArr = append(newArr, tempS)

		}
	}
	table := make(map[string]map[string]string)
	for _, j := range noEndSet.Traverse() {
		table[j] = make(map[string]string)
	}
	for _, tempString := range newArr {
		if yesEndSet.Get(string(tempString[3])) && tempString[3] != '#' {
			if table[string(tempString[0])][string(tempString[3])] == "" {
				table[string(tempString[0])][string(tempString[3])] = tempString
			} else {
				table[string(tempString[0])][string(tempString[3])] = "冲突"
				conflict = true
			}
		}
		if tempString[3] == '#' {
			followSlice := followMap[string(tempString[0])]
			for _, s := range followSlice {
				if table[string(tempString[0])][s] == "" {
					table[string(tempString[0])][s] = tempString
				} else {
					table[string(tempString[0])][s] = "冲突"
					conflict = true
				}
			}
		}
		if noEndSet.Get(string(tempString[3])) {
			slice := firstMap[string(tempString[3])]
			fEmpty := false
			for _, s := range slice {
				if s == "#" {
					fEmpty = true
				} else {
					if table[string(tempString[0])][s] == "" {
						table[string(tempString[0])][s] = tempString
					} else {
						table[string(tempString[0])][s] = "冲突"
						conflict = true
					}
				}
			}
			if fEmpty == true {
				followSlice := followMap[string(tempString[0])]
				for _, s := range followSlice {
					if table[string(tempString[0])][s] == "" {
						table[string(tempString[0])][s] = tempString
					} else {
						table[string(tempString[0])][s] = "冲突"
						conflict = true
					}
				}
			}
		}
	}
	return table, conflict
}

func Analyse(input string) (map[string][]string, map[string][]string, map[string]map[string]string, bool, []string, []string) {
	//input := "G->S S->Tfg T->(S)ST|a|#"
	//input := "S->iEtST|a T->bS|# E->c"

	arr, noEndSet, yesEndSet := initInput(input)
	firstMap, emptySET := ExuFirst(arr, noEndSet)
	followMap := ExuFollow(arr, noEndSet, emptySET, firstMap)
	table, conflict := ExcTable(arr, noEndSet, yesEndSet, firstMap, followMap)
	return firstMap, followMap, table, conflict, yesEndSet.Traverse(), noEndSet.Traverse()
}
