#include <stdio.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
    *returnSize = 2;
    int* map = (int*)malloc(numsSize * sizeof(int));
    int* result = (int*)malloc(*returnSize * sizeof(int));

    for (int i = 0; i < numsSize; i++) {
        map[i] = -1;
    }

    for (int i = 0; i < numsSize; i++) {
        int complement = target - nums[i];

        for (int j = 0; j < i; j++) {
            if (nums[j] == complement) {
                result[0] = j;
                result[1] = i;
                free(map);
                return result;
            }
        }
        map[i] = nums[i];
    }

    *returnSize = 0;
    free(map);

    free(result);

    return NULL;
}
