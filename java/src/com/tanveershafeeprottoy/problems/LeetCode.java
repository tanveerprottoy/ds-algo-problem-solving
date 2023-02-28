package com.tanveershafeeprottoy.problems;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.PriorityQueue;
import java.util.Set;

public class LeetCode {

    /*
     * sliding window technique used The technique can be best understood with
     * the window pane in bus, consider a window of length n and the pane which is
     * fixed in it of length k. Consider, initially the pane is at extreme left i.e.
     * at 0 units from the left. Now, co-relate the window with array arr[] of size
     * n and pane with current_sum of size k elements. Now, if we apply force on the
     * window such that it moves a unit distance ahead. The pane will cover next k
     * consecutive elements. Consider an array arr[] = {5, 2, -1, 0, 3} and value
     * of k = 3 and n = 5 Applying sliding window technique : We compute the sum of
     * first k elements out of n terms using a linear loop and store the sum in
     * variable
     * window_sum. Then we will graze linearly over the array till it reaches the
     * end
     * and simultaneously keep track of maximum sum. To get the current sum of block
     * of
     * k elements just subtract the first element from the previous block and add
     * the last
     * element of the current block.
     */
    public static int lengthOfLongestSubstring(String s) {
        int n = s.length();
        Set<Character> set = new HashSet<>();
        // j is fast (goes forward) i is slow (points to the start of a new substring)
        int ans = 0, i = 0, j = 0;
        while (i < n && j < n) {
            // try to extend the range [i, j]
            if (!set.contains(s.charAt(j))) {
                set.add(s.charAt(j++));
                ans = Math.max(ans, j - i);
            } else {
                set.remove(s.charAt(i++));
            }
        }
        System.out.println(set);
        return ans;
    }

    public static int lengthOfLongestSubstringOptimized(String s) {
        // j is fast (goes forward) i is slow (points to the start of a new substring)
        int n = s.length(), ans = 0;
        Map<Character, Integer> map = new HashMap<>(); // current index of character
        for (int j = 0, i = 0; j < n; j++) {
            if (map.containsKey(s.charAt(j))) {
                i = Math.max(map.get(s.charAt(j)), i);
            }
            ans = Math.max(ans, j - i);
            map.put(s.charAt(j), j);
        }
        return ans;
    }

    public static int lengthOfLongestSubstringAscii(String s) {
        // j is fast (goes forward) i is slow (points to the start of a new substring)
        int n = s.length(), ans = 0;
        int[] index = new int[128]; // current index of character
        // try to extend the range [i, j]
        for (int j = 0, i = 0; j < n; j++) {
            i = Math.max(index[s.charAt(j)], i);
            ans = Math.max(ans, j - i);
            index[s.charAt(j)] = j;
        }
        return ans;
    }

    public static int lengthOfLongestSubstringStr(String s) {
        int n = s.length();
        Set<Character> set = new HashSet<>();
        // j is fast (goes forward) i is slow (points to the start of a new substring)
        int ans = 0, i = 0, j = 0;
        StringBuilder stringBuilder = new StringBuilder();
        char c;
        String result = "";
        while (i < n && j < n) {
            // try to extend the range [i, j]
            c = s.charAt(j);
            if (!set.contains(c)) {
                set.add(c);
                j++;
                ans = Math.max(ans, j - i);
                stringBuilder.append(c);
            } else {
                c = s.charAt(i);
                set.remove(c);
                i++;
                stringBuilder.deleteCharAt(stringBuilder.indexOf(String.valueOf(c)));
            }
            if (stringBuilder.length() > result.length()) {
                result = stringBuilder.toString();
            }
        }
        System.out.println(set);
        System.out.println(result);
        return ans;
    }

    /*
     * need to find the left part of the merged array (won't be merged though)
     * cut at random position of a and b
     */
    public static double findMedianSortedArrays(int[] a, int[] b) {
        int m = a.length;
        int n = b.length;
        // to ensure m<=n
        if (m > n) {
            int[] temp = a;
            a = b;
            b = temp;
            int tmp = m;
            m = n;
            n = tmp;
        }
        int iMin = 0, iMax = m, halfLen = (m + n + 1) / 2;
        while (iMin <= iMax) {
            // cut [a] array in position i
            int i = (iMin + iMax) / 2;
            // cut [b] array in position j
            int j = halfLen - i;
            if (i < iMax && b[j - 1] > a[i]) {
                // i is too small
                // increase iMin to to cut position + 1
                iMin = i + 1;
            } else if (i > iMin && a[i - 1] > b[j]) {
                // i is too big
                // decrease iMax to to cut position - 1
                iMax = i - 1;
            } else {
                // i is perfect
                int maxLeft;
                if (i == 0) {
                    maxLeft = b[j - 1];
                } else if (j == 0) {
                    maxLeft = a[i - 1];
                } else {
                    maxLeft = Math.max(a[i - 1], b[j - 1]);
                }
                if ((m + n) % 2 == 1) {
                    return maxLeft;
                }
                int minRight;
                if (i == m) {
                    minRight = b[j];
                } else if (j == n) {
                    minRight = a[i];
                } else {
                    minRight = Math.min(b[j], a[i]);
                }

                return (maxLeft + minRight) / 2.0;
            }
        }
        return 0.0;
    }

    /*
     * Given an array of integers arr and an integer k.
     * 
     * Find the least number of unique integers after removing exactly k elements.
     * Example 1:
     * Input: arr = [5,5,4], k = 1
     * Output: 1
     * Explanation: Remove the single 4, only 5 is left.
     * Example 2:
     * Input: arr = [4,3,1,1,3,3,2], k = 3
     * Output: 2
     * Explanation: Remove 4, 2 and either one of the two 1s or three 3s. 1 and 3
     * will be left.
     * 
     * Constraints:
     * 
     * 1 <= arr.length <= 10^5
     * 1 <= arr[i] <= 10^9
     * 0 <= k <= arr.length
     */
    static int findLeastNumOfUniqueInts(int[] arr, int k) {
        Map<Integer, Integer> count = new HashMap<>();
        for (int a : arr)
            count.put(a, 1 + count.getOrDefault(a, 0));
        PriorityQueue<Integer> pq = new PriorityQueue<>(count.values());
        while (k > 0)
            k -= pq.poll();
        return k < 0 ? pq.size() + 1 : pq.size();
    }

    static int reverseInt(int x) {
        int rev = 0;
        while (x != 0) {
            int pop = x % 10;
            x /= 10;
            if (rev > Integer.MAX_VALUE / 10 || (rev == Integer.MAX_VALUE / 10 && pop > 7))
                return 0;
            if (rev < Integer.MIN_VALUE / 10 || (rev == Integer.MIN_VALUE / 10 && pop < -8))
                return 0;
            rev = rev * 10 + pop;
        }
        return rev;
    }

    static String longestPalindrome(String s) {
        if (s == null || s.length() < 1)
            return "";
        int start = 0, end = 0;
        for (int i = 0; i < s.length(); i++) {
            int len1 = expandAroundCenter(s, i, i);
            int len2 = expandAroundCenter(s, i, i + 1);
            int len = Math.max(len1, len2);
            if (len > end - start) {
                start = i - (len - 1) / 2;
                end = i + len / 2;
            }
        }
        return s.substring(start, end + 1);
    }

    static int expandAroundCenter(String s, int left, int right) {
        int L = left, R = right;
        while (L >= 0 && R < s.length() && s.charAt(L) == s.charAt(R)) {
            L--;
            R++;
        }
        return R - L - 1;
    }

    public static void main(String[] args) {
        /* Set set = new HashSet<String>();
        set. */
        int arr[] = { 2, 1, 1, 3, 3, 3 };
        int result = findLeastNumOfUniqueInts(arr, 3);
        System.out.print(result);
        System.out.print(reverseInt(-2147483412));
    }
}
