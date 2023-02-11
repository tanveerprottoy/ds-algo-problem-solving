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

    static int binarySearch(int array[], int x, int low, int high) {

        // Repeat until the pointers low and high meet each other
        while (low <= high) {
            int mid = low + (high - low) / 2;

            if (array[mid] == x)
                return mid;

            if (array[mid] < x)
                low = mid + 1;

            else
                high = mid - 1;
        }

        return -1;
    }

    static int gcd(int a, int b) {
        if (a == 0)
            return b;
        return gcd(b % a, a);
    }

    // function to calculate
    // lcm of two numbers.
    static int lcm(int a, int b) {
        return (a * b) / gcd(a, b);
    }

    public static void main(String[] args) {
        /*
         * System.out.println(
         * reverseString("abcde"));
         * 
         * System.out.println(
         * reverseString("123456"));
         */
        int arr[] = { 20, 5, 4, 6, 9 };
        int arrSize = arr.length;
        int[] result = largestSmallestNumberInArray(arr, arrSize);
        System.out.print(result);

        int array[] = { 3, 4, 5, 6, 7, 8, 9 };
        int n = array.length;
        int x = 4;
        int result1 = binarySearch(array, x, 0, n - 1);
        if (result1 == -1)
            System.out.println("Not found");
        else
            System.out.println("Element found at index " + result);
    }
}
