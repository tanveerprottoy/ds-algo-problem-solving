#include <iostream>

using namespace std;

int* largestSmallestNumberInArray(int arr[], unsigned int arrSize) {
    int* result = new int[2]{-1, -1};
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

int main() {
    int arr[] = {20, 5, 4, 6, 9};
    unsigned int arrSize = sizeof(arr) / sizeof(arr[0]);
    int* result = largestSmallestNumberInArray(arr, arrSize);
    cout << result[0] << ", " << result[1];
    return 0;
}
