func FindOccurences() {
	// find the  occurance of character 
	s := "helloeveryone"
	varmap := make(map[string]int)
	for _, v := range s {
		varmap[string(v)] = 0
	}
	var newvalue int
	for _, v := range s{
		value, _ := varmap[string(v)]
		if strings.Contains(s, string(v)) {
			newvalue = value + 1 
			varmap[string(v)] = newvalue
		}
	} 
	fmt.Println(varmap)
	
	//find the most occured character
	var maxchar string
	var maxint int
	for i:=0; i < len(varmap)-1 ; i++{
		for _, v := range s{
			value, _ := varmap[string(v)]
			if maxint < value {
				maxint = value 
				maxchar = string(v)
			}
		} 
	}
	fmt.Println("Most Occured character :", maxchar)
}