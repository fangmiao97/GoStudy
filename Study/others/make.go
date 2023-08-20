package others

import "fmt"

type UserSegmentParam struct {
	UserID      int64
	UserSegment string
}

var params = []UserSegmentParam{
	{UserID: 1200379093, UserSegment: "A"},
	{UserID: 1200379094, UserSegment: "B"},
	{UserID: 1200379095, UserSegment: "C"},
	{UserID: 1200379096, UserSegment: "D"},
}

func makefun() {
	queryMap := make(map[int64][]string, 8)
	// print queryMap length
	fmt.Printf("len(queryMap): %d\n", len(queryMap))
	// print queryMap
	fmt.Printf("queryMap: %v\n", queryMap)
	for _, param := range params {
		queryMap[param.UserID] = append(queryMap[param.UserID], param.UserSegment)
	}
	// print queryMap length
	fmt.Printf("len(queryMap): %d\n", len(queryMap))
	// print queryMap
	fmt.Printf("queryMap: %v\n", queryMap)

}
