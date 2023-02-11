package com.tanveershafeeprottoy.problems;

import java.util.Collections;
import java.util.List;

public class HackerRank {

    static String timeInWords(int hour, int minute) {
        String timeWord = "";
        final String[] hourStrings = { "0", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
                "ten", "eleven", "twelve" };

        final String[] minuteStrings = { " o' clock", "one", "two",
                "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen",
                "fourteen", "quarter", "sixteen",
                "seventeen", "eighteen",
                "nineteen", "twenty",
                "twenty one", "twenty two",
                "twnenty three", "twenty four",
                "twenty five", "twenty six",
                "twenty seven", "twenty eight",
                "twenty nine", "half" };
        if (minute == 0)
            timeWord = hourStrings[hour] + minuteStrings[minute];
        else if (minute == 1)
            timeWord = minuteStrings[minute] + " minute past " +
                    hourStrings[hour];
        else if (minute == 15)
            timeWord = "quarter past " + hourStrings[hour];
        else if (minute <= 29)
            timeWord = minuteStrings[minute] + " minutes past " +
                    hourStrings[hour];
        else if (minute == 30)
            timeWord = "half past " + hourStrings[hour];
        else if (minute == 45)
            timeWord = "quarter to " + hourStrings[hour + 1];
        else if (minute <= 59)
            timeWord = minuteStrings[60 - minute] + " minutes to " +
                    hourStrings[hour + 1];
        return timeWord;
    }

    static int getTotalX(List<Integer> a, List<Integer> b) {
        int max = Collections.max(a);
        int min = Collections.min(b);
        int count = 0;
        for (int i = max; i <= min; i++) {
            /*
             * Since the range is going to be the max element in the a list and the min
             * element in the b list
             *
             * variable to keep track the number of element in the b list that are divisible
             * by i
             */
            int k = 0;
            for (int j = 0; j < b.size(); j++) {
                if (b.get(j) % i == 0) {
                    k++;
                }
            }
            if (k == b.size()) {
                /* to keep track of number of elements that can divide i in the list a */
                int n = 0;
                for (int j = 0; j < a.size(); j++) {
                    if (i % a.get(j) == 0) {
                        n++;
                    }
                }
                if (n == a.size()) {
                    /*
                     * if all conditions are satisfies then take that number into account by
                     * increasing the count
                     */
                    count++;
                }
            }
        }
        return count;
    }

    public static void main(String[] args) {
        String result = timeInWords(5, 47);
        System.out.println(result);
    }
}
