//input: k=3, str=deeedbbcccbdaa
//output: str=aa
//1. ddbbbdaa
//2. dddaa
//3. aa

package main

import "fmt"

func main() {
	k := 3
	str := "deeedbbcccbdaa"
	//str := "aabbccc"
	newstr := RemoveConscutive(str, k)
	fmt.Println("output:")
	fmt.Println(newstr)
}

func RemoveConscutive(str string, k int) string {
	removedstr := false
	for j := 0; j < len(str)-(k-1); j++ {
		for i := 1; i <= k; i++ {
			if i == k {
				str = string(str[:j]) + string(str[j+i:])
				removedstr = true
			} else if str[j] != str[j+i] {
				break
			}
		}
	}
	if !removedstr {
		return str
	}
	str = RemoveConscutive(str, k)
	return str

}
