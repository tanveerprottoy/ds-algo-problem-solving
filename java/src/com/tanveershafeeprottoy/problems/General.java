package com.tanveershafeeprottoy.problems;

public class General {

    // O(n) solution for finding
    // maximum sum of a subarray
    // of size k
    static int maxSum(int[] arr, int n, int k) {
        // n must be greater
        if(n < k) {
            System.out.println("Invalid");
            return -1;
        }

        // Compute sum of first window of size k
        int max_sum = 0;
        for(int i = 0; i < k; i++)
            max_sum += arr[i];

        // Compute sums of remaining windows by
        // removing first element of previous
        // window and adding last element of
        // current window.
        int window_sum = max_sum;
        for(int i = k; i < n; i++) {
            window_sum += arr[i] - arr[i - k];
            max_sum = Math.max(max_sum, window_sum);
        }

        return max_sum;
    }

    static int maxSum2(int[] arr, int n, int k) {
        if(n < k) {
            System.out.println("Invalid");
            return -1;
        }
        return 0;
    }

    public static boolean isPalindrome(String value) {
        int i = 0, j = value.length() - 1;
        while(i < value.length() && j >= 0) {
            if(value.charAt(i) != value.charAt(j)) {
                return false;
            }
            i++;
            j--;
        }
        return true;
    }

    public static int[] reverseArray(int[] value) {
        int length = value.length;
        int[] result = new int[length];
        int i = value.length - 1, j = 0;
        while(i >= 0 && j < length) {
            result[j] = value[i];
            i--;
            j++;
        }
        return result;
    }

    public static char[] reverseArray(char[] value) {
        int length = value.length;
        char[] result = new char[length];
        int i = value.length - 1, j = 0;
        while(i >= 0 && j < length) {
            result[j] = value[i];
            i--;
            j++;
        }
        return result;
    }
}
