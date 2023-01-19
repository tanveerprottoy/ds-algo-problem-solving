package com.tanveershafeeprottoy.problems;

public class Miscs {

    static String reverseString(String str) {
        StringBuilder reversed = new StringBuilder();
        for (int i = str.length() - 1; i >= 0; i--) {
            reversed.append(str.charAt(i));
        }
        return reversed.toString();
    }

    static int[] largestSmallestNumberInArray(int arr[], int arrSize) {
        int[] result = new int[] { -1, -1 };
        // Corner case
        if (arrSize <= 0) {
            return result;
        }
        int val = 0;
        for (int i = 0; i < arrSize; i++) {
            val = arr[i];
            // check large
            if (result[0] < val) {
                result[0] = val;
            }
            if (result[1] == -1 || result[1] > val) {
                result[1] = val;
            }
        }
        return result;
    }

    public static void main(String[] args) {
        /*
         * System.out.println(
         * Miscs.reverseString("abcde"));
         * 
         * System.out.println(
         * Miscs.reverseString("123456"));
         */
        int arr[] = { 20, 5, 4, 6, 9 };
        int arrSize = arr.length;
        int[] result = largestSmallestNumberInArray(arr, arrSize);
        System.out.print(result);
    }
}
