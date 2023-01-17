package com.tanveershafeeprottoy.problems;

public class Miscs {

    static String reverseString(String str) {
        StringBuilder reversed = new StringBuilder();
        for(int i = str.length() - 1; i >= 0; i--) {
            reversed.append(str.charAt(i));
        }
        return reversed.toString();
    }

    public static void main(String[] args) {
        System.out.println(
            Miscs.reverseString("abcde")
        );

        System.out.println(
            Miscs.reverseString("123456")
        );
    }
}
