package com.tanveershafeeprottoy.algorithms;

/*
Loop invariant:
    Initialization: It is true prior to the first iteration of the loop.
    Maintenance: If it is true before an iteration of the loop, it remains true before the
    next iteration.
    Termination: When the loop terminates, the invariant gives us a useful property that helps show that the algorithm is correct.
    When the first two properties hold, the loop invariant is true prior to every iteration of the loop. (Of course, we are free to use established facts other than the loop invariant itself to prove that the loop invariant remains true before each iteration.) Note the similarity to mathematical induction, where to prove that a property holds, you prove a base case and an inductive step. Here, showing that the invariant holds before the first iteration corresponds to the base case, and showing that the invariant holds from iteration to iteration corresponds to the inductive step.
    The third property is perhaps the most important one, since we are using the loop invariant to show correctness. Typically, we use the loop invariant along with the condition that caused the loop to terminate. The termination property differs from how we usually use mathematical induction, in which we apply the inductive step infinitely; here, we stop the “induction” when the loop terminates.
    Let us see how these properties hold for insertion sort.
    Initialization: We start by showing that the loop invariant holds before the first
    1
    Maintenance: Next, we tackle the second property: showing that each iteration maintains the loop invariant. Informally, the body of the for loop works by moving AŒj 􏰐 1􏰀, AŒj 􏰐 2􏰀, AŒj 􏰐 3􏰀, and so on by one position to the right until it finds the proper position for AŒj 􏰀 (lines 4–7), at which point it inserts the value of AŒj 􏰀 (line 8). The subarray AŒ1 : : j 􏰀 then consists of the elements originally in AŒ1 : : j 􏰀, but in sorted order. Incrementing j for the next iteration of the for loop then preserves the loop invariant.
    A more formal treatment of the second property would require us to state and show a loop invariant for the while loop of lines 5–7. At this point, however,
    1When the loop is a for loop, the moment at which we check the loop invariant just prior to the first iteration is immediately after the initial assignment to the loop-counter variable and just before the first test in the loop header. In the case of INSERTION-SORT, this time is after assigning 2 to the variable j but before the first test of whether j 􏰎 A:length.
    loop iteration, when j D 2.
    of just the single element AŒ1􏰀, which is in fact the original element in AŒ1􏰀. Moreover, this subarray is sorted (trivially, of course), which shows that the loop invariant holds prior to the first iteration of the loop.
    The subarray AŒ1 : : j 􏰐 1􏰀, therefore, consists

    20 Chapter 2 Getting Started
    we prefer not to get bogged down in such formalism, and so we rely on our informal analysis to show that the second property holds for the outer loop.
    Termination: Finally, we examine what happens when the loop terminates. The condition causing the for loop to terminate is that j > A:length D n. Because each loop iteration increases j by 1, we must have j D n C 1 at that time. Substituting n C 1 for j in the wording of loop invariant, we have that the subarray AŒ1 : : n􏰀 consists of the elements originally in AŒ1 : : n􏰀, but in sorted order. Observing that the subarray AŒ1 : : n􏰀 is the entire array, we conclude that the entire array is sorted. Hence, the algorithm is correct.
    We shall use this method of loop invariants to show correctness later in this chapter and in other chapters as well.
*/

public class Sort {
    /*  insertion sort
        for j = 2 to A:length
            key = A[j]􏰀
            // Insert A[j] 􏰀 into the sorted sequence A[1..j - 1]
            i = j - 􏰐1
            while i > 0 and A[i]􏰀>key
                A[i +1] = A[i]􏰀
                i = i-1
            A[i+1]􏰀=key*/
    public static int[] insertionSort(int[] numArray) {
        int element, i;
        for(int j = 1; j < numArray.length; j++) {
            // numArray[1..j - 1] = currently shorted range in array
            // numArray[j + 1..n] = left to be shorted range in array
            element = numArray[j];
            // go back one index
            i = j - 1;
            while(i >= 0 && numArray[i] > element) {
                // this is where the sorting happens
                // this traverses from j - 1 to 0 [goes back] and sorts
                // previous values if required
                // move numArray[i] to next index
                numArray[i + 1] = numArray[i];
                // decrement i to go back 1 index
                i = i - 1;
            }
            // after previous elements sorting done
            // put current value in the last traversed index
            numArray[i + 1] = element;
        }
        return numArray;
    }

    public static void mergeSort(
        Comparable[] a,
        int lo,
        int mid,
        int hi
    ) { // Merge a[lo..mid] with a[mid+1..hi].
        int i = lo, j = mid + 1;
        Comparable[] aux = new Comparable[a.length];
        for(int k = lo; k <= hi; k++) // Copy a[lo..hi] to aux[lo..hi].
            aux[k] = a[k];
        for(int k = lo; k <= hi; k++) // Merge back to a[lo..hi].
            if(i > mid) a[k] = aux[j++];
            else if(j > hi) a[k] = aux[i++];
            // else if(less(aux[j], aux[i])) a[k] = aux[j++];
            else a[k] = aux[i++];
    }
}
