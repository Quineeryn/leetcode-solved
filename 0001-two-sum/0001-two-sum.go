package main

func twoSum(nums []int, target int) []int {
    idx := make(map[int]int)

    for i, v := range nums{
        if j, ok := idx[target-v]; ok{
            return []int{j, i}
        }
        idx[v] = i
    }
    return nil
}