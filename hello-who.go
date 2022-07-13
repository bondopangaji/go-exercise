/*
 * Copyright (c) 2022 Bondo Pangaji
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("======================")
	fmt.Println("===== Hello who? =====")
	fmt.Println("======================")

	for {
		fmt.Print("-> ")
		scanner.Scan()
		fmt.Println("Hello", scanner.Text())
	}
}
