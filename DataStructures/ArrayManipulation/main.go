package main

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	ar := make([]int64, n+1)
	for i := 0; i < len(queries); i++ {
		ar[queries[i][0]-1] += int64(queries[i][2])
		ar[queries[i][1]] += -int64(queries[i][2])
	}
	var max, runningCount int64 = 0, 0
	for i := 0; i < len(ar); i++ {
		runningCount += ar[i]
		if runningCount > max {
			max = runningCount
		}

	}
	return max
}

func main() {
	var n int32 = 40
	var queries = [][]int32{
		{29, 40, 787},
		{9, 26, 219},
		{21, 31, 214},
		{8, 22, 719},
		{15, 23, 102},
		{11, 24, 83},
		{14, 22, 321},
		{5, 22, 300},
		{11, 30, 832},
		{5, 25, 29},
		{16, 24, 577},
		{3, 10, 905},
		{15, 22, 335},
		{29, 35, 254},
		{9, 20, 20},
		{33, 34, 351},
		{30, 38, 564},
		{11, 31, 969},
		{3, 32, 11},
		{29, 35, 267},
		{4, 24, 531},
		{1, 38, 892},
		{12, 18, 825},
		{25, 32, 99},
		{3, 39, 107},
		{12, 37, 131},
		{3, 26, 640},
		{8, 39, 483},
		{8, 11, 194},
		{12, 37, 502}}

	result := arrayManipulation(n, queries)
	println(result)

}
