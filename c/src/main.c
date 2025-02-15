#include <stdio.h>
#include <stdlib.h>

#include "./problems/leetcode.c"

int main() {
    // twoSum();

    // int* nums = (int*)malloc(5 * sizeof(int))

    int nums[10] = {0, 0, 1, 1, 1, 2, 2, 3, 3, 4};

    int* ptr = nums;

    int size = sizeof(nums) / sizeof(nums[0]);

    printf("%d", removeDuplicates(ptr, size));

    return 0;
}