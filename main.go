package main

func main()  {
	RunGin()
}




func main2()  {
	goodUsers := []*User{}

	ru := RecommendListGoodUser()
	goodUsers = append(goodUsers, ru...)

	du := DynamicListGoodUser()
	goodUsers = append(goodUsers, du...)

	//放进本地goodUsers
	for _,u := range goodUsers {
		if RedisCheckBlockUser(u) {
			continue
		}
		RedisAddGoodUser(u)
	}
}