package leetcode

import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	foodMap := make(map[string]struct{})
	tableMap := make(map[int]map[string]int)
	for _, order := range orders {
		id, _ := strconv.Atoi(order[1])
		food := order[2]
		foodMap[food] = struct{}{}
		if tableMap[id] == nil {
			tableMap[id] = make(map[string]int)
		}
		tableMap[id][food]++
	}
	foods := make([]string, 0, len(foodMap))
	for food := range foodMap {
		foods = append(foods, food)
	}
	sort.Strings(foods)
	tables := make([]int, 0, len(tableMap))
	for table := range tableMap {
		tables = append(tables, table)
	}
	sort.Ints(tables)
	res := make([][]string, len(tableMap)+1)
	res[0] = make([]string, 1, len(tableMap)+1)
	res[0][0] = "Table"
	res[0] = append(res[0], foods...)
	for i, table := range tables {
		cnt := tableMap[table]
		res[i+1] = make([]string, len(foodMap)+1)
		res[i+1][0] = strconv.Itoa(table)
		for j, food := range foods {
			res[i+1][j+1] = strconv.Itoa(cnt[food])
		}
	}
	return res
}
