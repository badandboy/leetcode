package canVisitAllRooms

func dfs(rooms [][]int, room int, visted map[int]bool) {

}

func CanVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visted := map[int]bool{
		0: true,
	}

	var dfs func(room int)
	dfs = func(room int) {
		keys := rooms[room]
		for _, key := range keys {
			if _, ok := visted[key]; !ok {
				visted[key] = true
				dfs(key)
			}
		}
	}

	dfs(0)
	if len(visted) < n {
		return false
	}

	return true
}
