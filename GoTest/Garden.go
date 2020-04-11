package main

import "fmt"

const Allflower = 97
type FlowerNames [Allflower]string
type NameLength [Allflower]int

func collectFlowers(t *FlowerNames, n *int) {
	var temp string
	var idx int = 0
	fmt.Println("Enter flower names :")
	fmt.Scanln(&temp)

	for temp != "STOP" {
		t[idx] = temp
		idx++
		fmt.Scanln(&temp)
	}
	*n = idx
}

func computeLength(t FlowerNames, n int, f *NameLength) {
	var idx int
	for idx = 0; idx <= n-1; idx++ {
		f[idx] = len(t[idx]) 
	}
}

func printFlowers(t FlowerNames, n int, names bool) {
	var idx int
	if names {
		fmt.Println("List of flowers :")
		for idx = 0; idx <= n-1; idx++ {
			fmt.Println(t[idx])
		}
	} else {
		fmt.Println("Length of names :")
		for idx = 0; idx <= n-1; idx++ {
			fmt.Println(len(t[idx]))
		}
	}
}

func sortByName(t *FlowerNames, n int) {
	var idx, idx2 int
    for idx = 0; idx < n; idx++ {
        var minIdx = idx
        for idx2 = idx; idx2 < n; idx2++ {
            if t[idx2] < t[minIdx] {
                minIdx = idx2
            }
        }
        t[idx], t[minIdx] = t[minIdx], t[idx]
	}
}

func sortByLength(t *FlowerNames, n int) {
	var idx, idx2 int
    for idx = 0; idx < n; idx++ {
        var minIdx = idx
        for idx2 = idx; idx2 < n; idx2++ {
            if len(t[idx2]) > len(t[minIdx]) {
                minIdx = idx2
            }
        }
        t[idx], t[minIdx] = t[minIdx], t[idx]
	}
}
func main() {
	var flower FlowerNames
	var length NameLength
	var M int

	collectFlowers(&flower, &M)
	computeLength(flower, M, &length)
	sortByName(&flower, M)
	printFlowers(flower, M, true)
	sortByLength(&flower, M)
	printFlowers(flower, M, false)

}