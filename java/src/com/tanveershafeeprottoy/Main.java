package com.tanveershafeeprottoy;

import com.tanveershafeeprottoy.algorithms.Sort;
import com.tanveershafeeprottoy.ds.Bracket;
import com.tanveershafeeprottoy.problems.General;

import java.util.Arrays;

public class Main {

    public static void main(String[] args) {
        System.out.println(
            Arrays.toString(
                Sort.insertionSort(
                    new int[] {
                        5, 2, 4, 6, 9
                    }
                )
            )
        );
        //System.out.println(LeetCode.lengthOfLongestSubstring("anviaj"));
        // System.out.println(LeetCode.lengthOfLongestSubstring("hkcpmprxxxqw"));
        // System.out.println(LeetCode.lengthOfLongestSubstringStr("hkcpmprxxxqw"));
        // System.out.println(LeetCode.lengthOfLongestSubstringOptimized("pwwkew"));
        // System.out.println(LeetCode.lengthOfLongestSubstringAscii("pwwkew"));
        /*System.out.println(
            LeetCode.findMedianSortedArrays(
                new int[] {
                    1, 2
                },
                new int[] {
                    3, 4
                }
            )
        );
        System.out.println(
            LeetCode.findMedianSortedArrays(
                new int[] {
                    1, 2, 3, 6, 7
                },
                new int[] {
                    4, 6, 8, 10, 11, 12, 13
                }
            )
        );*/
        int[] arr = { 1, 4, 2, 10, 2, 3, 1, 0, 20 };
        int k = 4;
        int n = arr.length;
        //System.out.println(General.maxSum(arr, n, k));
        //System.out.println(General.isPalindrome("abba"));
        System.out.println(Arrays.toString(General.reverseArray(new int[] { 1, 2, 3 })));
        System.out.println(Arrays.toString(General.reverseArray(new char[] { 'C', 'B', 'A' })));

        // ds
        System.out.println("" + Bracket.isBalanced(
            "{[()]}"
        ));
    }
}
