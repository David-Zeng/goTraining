////////////////////////////////////////////////////////////////////////////////
// name:
//
//
////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "chocolate", "martini"}
	fmt.Println(jb)

	mp := []string{"miss", "moneyPenny", "strawBerry", "hazelNut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
