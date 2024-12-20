package main

import "fmt"

// 73	74	75	76	77	78	79	80	81
// 72	43	44	45	46	47	48	49	50
// 71	42	21	22	23	24	25	26	51
// 70	41	20	7	8	9	10	27	52
// 69	40	19	6	1	2	11	28	53
// 68	39	18	5	4	3	12	29	54
// 67	38	17	16	15	14	13	30	55
// 66	37	36	35	34	33	32	31	56
// 65	64	63	62	61	60	59	58	57

func main(){
	fmt.Println(findSumOfSpiralDiags(1001))
}

func findSumOfSpiralDiags(size int) int{
	sum := 1 
	for n := 3; n <= size; n += 2 {
		topRight := n * n
		topLeft := topRight - (n - 1)
		bottomLeft := topLeft - (n - 1)
		bottomRight := bottomLeft - (n - 1)
		
		sum += topRight + topLeft + bottomLeft + bottomRight
	}
	return sum
}