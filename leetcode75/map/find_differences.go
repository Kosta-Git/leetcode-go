package _map

func findDifference(nums1 []int, nums2 []int) [][]int {
    map1, map2 := make(map[int]bool, len(nums1)), make(map[int]bool, len(nums2))
    for _, val := range nums1 {
        map1[val] = true
    }
    for _, val := range nums2 {
        map2[val] = true
    }
    output := make([][]int, 2)
    output[0], output[1] = []int{}, []int{}
    for key, _ := range map1 {
        if _, ok := map2[key]; !ok {
            output[0] = append(output[0], key)
        }
    }
    for key, _ := range map2 {
        if _, ok := map1[key]; !ok {
            output[1] = append(output[1], key)
        }
    }
    return output
}
