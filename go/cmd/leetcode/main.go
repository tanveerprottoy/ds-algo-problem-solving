package main

import (
	"log"

	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/leetcode"
)

func main() {
	// fmt.Printf("%f", leetcode.MedianOfTwoSortedArray([]int{3}, []int{-2, -1}))
	// fmt.Printf("%f", leetcode.MedianOfTwoSortedArray([]int{1,2}, []int{3, 4}))
	// fmt.Printf("%d", leetcode.LengthOfLongestSubstring("abcabcbb"))
	// fmt.Printf("%d", leetcode.LengthOfLongestSubstring("bbbbb"))
	// fmt.Println(leetcode.LengthOfLongestSubstring("tmmzuxt"))
	/* fmt.Println(leetcode.LeastNumOfUniqueInts([]int{5, 5, 4}, 1))
	fmt.Println(leetcode.LeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3))
	fmt.Println(leetcode.LeastNumOfUniqueInts([]int{2, 1, 1, 3, 3, 3}, 3)) */
	// fmt.Println(leetcode.LeastNumOfUniqueInts([]int{4, 4, 4, 4, 1}, 2))
	// fmt.Println(leetcode.LeastNumOfUniqueInts([]int{1, 2, 2, 2, 2}, 4))
	// fmt.Println(leetcode.LeastNumOfUniqueInts([]int{2, 1, 1, 3, 3, 3}, 3))
	// fmt.Println(leetcode.ReverseInteger(123))
	// fmt.Println(leetcode.ReverseInteger(-541))
	//fmt.Println(leetcode.ReverseInteger(-2147483412))
	// fmt.Println(leetcode.ReverseInteger(-1463847412))
	// fmt.Println(leetcode.ArrayRankTransform([]int{40, 10, 20, 30}))
	// fmt.Println(leetcode.ArrayRankTransform([]int{100, 100, 100}))
	// fmt.Println(leetcode.ArrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))
	// fmt.Println(leetcode.RomanToInt("MCMXCIV"))
	// fmt.Println(leetcode.LongestCommonPrefixVertical([]string{"flower","flow","flight"}))
	// fmt.Println(leetcode.LongestCommonPrefixHorizontal([]string{"flower","flow","flight"}))
	/* l2 := new(leetcode.MergerListNode)
	l2.Val = 4
	l1 := new(leetcode.MergerListNode)
	l1.Val = 2
	l1.Nxt = l2
	l0 := new(leetcode.MergerListNode)
	l0.Val = 1
	l0.Nxt = l1
	m2 := new(leetcode.MergerListNode)
	m2.Val = 4
	m1 := new(leetcode.MergerListNode)
	m1.Val = 3
	m1.Nxt = m2
	m0 := new(leetcode.MergerListNode)
	m0.Val = 1
	m0.Nxt = m1
	fmt.Println(leetcode.MergeTwoLinkedLists(l0, m0)) */
	// fmt.Println(leetcode.IsValidParentheses("()[]{}"))
	// fmt.Println(leetcode.RemoveElement([]int{3, 2, 2, 3}, 3))
	// fmt.Println(leetcode.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	// fmt.Println(leetcode.RemoveDuplicates([]int{1, 1, 2}))
	// fmt.Println(leetcode.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	/* 	fmt.Println(leetcode.RemoveDuplicatesOverTwo([]int{1, 1, 1, 2, 2, 3}))
	   	fmt.Println(leetcode.RemoveDuplicatesOverTwo([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})) */
	/* fmt.Println(leetcode.PlusOne([]int{1, 2, 3}))
	fmt.Println(leetcode.PlusOne([]int{9}))
	fmt.Println(leetcode.PlusOne([]int{9, 9})) */
	/* fmt.Println(leetcode.IsPalindromeAfterRemoval("A man, a plan, a canal: Panama"))
	fmt.Println(leetcode.IsPalindromeAfterRemoval("race a car"))
	fmt.Println(leetcode.IsPalindromeAfterRemoval(" ")) */
	// fmt.Println(leetcode.IsPalindromeAfterRemoval(" "))
	// fmt.Println(leetcode.RemoveDuplicatesOverTwo([]int{1, 1, 1, 2, 2, 3}))
	// fmt.Println(leetcode.RemoveDuplicatesOverTwo([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
	/* fmt.Println(leetcode.SearchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(leetcode.SearchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(leetcode.SearchInsert([]int{1, 3, 5, 6}, 0)) */
	// fmt.Println(leetcode.LengthOfLastWord("Hello World"))
	// fmt.Println(leetcode.LengthOfLastWord("   fly me   to   the moon  "))
	// fmt.Println(leetcode.LengthOfLastWord("moon"))
	// fmt.Println(leetcode.LengthOfLastWord("moon"))
	// fmt.Println(leetcode.ContainsDuplicate([]int{1,2,3,4}))
	// fmt.Println(leetcode.MoveZeroes([]int{0, 1, 0, 3, 12}))
	// fmt.Println(leetcode.MajorityElement([]int{3, 2, 3}))
	// fmt.Println(leetcode.MajorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	// fmt.Println(leetcode.MajorityElement([]int{1}))
	// fmt.Println(leetcode.SingleNumber([]int{2, 2, 1}))
	// fmt.Println(leetcode.SingleNumber([]int{4, 1, 2, 1, 2}))
	/* l3 := linkedlist.NewNode(4, nil)
	l2 := linkedlist.NewNode(3, l3)
	l1 := linkedlist.NewNode(2, nil)
	l0 := linkedlist.NewNode(1, l1)
	res := leetcode.MergeTwoLists(l0, l2)
	res.Traverse(res) */

	/* l2 := linkedlist.NewNode(1, nil)
	l0 := linkedlist.NewNode(2, nil)
	res := leetcode.MergeTwoLists(l0, l2)
	res.Traverse(res) */
	// leetcode.MergeSortedArraysSimple([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
	/* leetcode.MergeSortedArrays([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
	leetcode.MergeSortedArrays([]int{0}, 0, []int{1}, 1)
	leetcode.MergeSortedArraysForDLC([]int{0}, 0, []int{1}, 1) */
	// fmt.Println(leetcode.SortedSquares([]int{-4, -1, 0, 3, 10}))
	/* Input: head = [1,1,2]
	Output: [1,2]
		Input: head = [1,1,2,3,3]
	Output: [1,2,3] */
	/* l2 := linkedlist.NewNode(2, nil)
	l1 := linkedlist.NewNode(1, l2)
	l0 := linkedlist.NewNode(1, l1)
	res := leetcode.DeleteDuplicates(l0)
	res.Traverse(res)

	l4 := linkedlist.NewNode(3, nil)
	l3 := linkedlist.NewNode(3, l4)
	l2 = linkedlist.NewNode(2, l3)
	l1 = linkedlist.NewNode(1, l2)
	l0 = linkedlist.NewNode(1, l1)
	res = leetcode.DeleteDuplicates(l0)
	res.Traverse(res)

	l2 = linkedlist.NewNode(1, nil)
	l1 = linkedlist.NewNode(1, l2)
	l0 = linkedlist.NewNode(1, l1)
	res = leetcode.DeleteDuplicates(l0)
	res.Traverse(res) */

	/* l4 := linkedlist.NewNode(3, nil)
	l3 := linkedlist.NewNode(3, l4)
	l2 := linkedlist.NewNode(1, l3)
	l1 := linkedlist.NewNode(1, l2)
	l0 := linkedlist.NewNode(1, l1)
	res := leetcode.DeleteDuplicates2(l0)
	res.Traverse(res) */

	// fmt.Println(leetcode.AddBinary("11", "1"))
	// fmt.Println(leetcode.AddBinary("10", "1011"))
	/* l0 := generic.NewNode(1, nil)
	fmt.Println(leetcode.HasCycleSimple(l0)) */

	// two sum
	log.Println(leetcode.TwoSum([]int{2, 7, 11, 15}, 9))
}
