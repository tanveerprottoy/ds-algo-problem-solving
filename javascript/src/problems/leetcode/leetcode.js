import { NumberLinkedList } from "./number-linked-list.js";

export class LeetCode {

    twoSum(nums, target) {
        indiceValueMap = new Map();
        indices = [];
        for(let i = 0; i < nums.length; i++) {
            const num = nums[i];
            // subtract from target
            const required = target - num;
            // check if the required value to add up to target exists in the map
            const index = indiceValueMap.get(required);
            if(index) {
                indices.push(index);
                indices.push(i);
                continue;
            }
            // store num in map
            indiceValueMap.set(num, i)
        }
        return indices;
    }

    addTwoNumbers(numberLs0, numberLs1) {
        let temp0 = numberLs0;
        let temp1 = numberLs1;
        let result = null;
        let carry = -1;
        let sumVal = 0;
        while(temp0 != null || temp1 != null || carry > 0) {
            if(temp0 != null) {
                // sum up
                sumVal += temp0.value;
                // set to next node
                temp0 = temp0.next;
            }
            if(temp1 != null) {
                // sum up
                sumVal = temp1.value;
                // set to next node
                temp1 = temp1.next;
            }
            if(sumVal >= 10) {
                // need to separate carry value
                // get floor value 
                carry = Math.floor(sumVal / 10);
                sumVal = sumVal % 10;
            }
            const current = new NumberLinkedList(sumVal);
            if(result === null) {
                // store current to result if result is null
                result = current;
            }
            else {
                // traverse to the last node of result and add current there
                let temp2 = result;
                while(temp2.next != null) {
                    temp2 = temp2.next;
                }
                temp2.next = current;
            }
        }
        return result;
    }

    addTwoNumbersRecursive(numberLs0, numberLs1, carry) {
        if(!numberLs0) {
            return numberLs1;
        }
        if(!numberLs1) {
            return numberLs0;
        }
        const sumVal = numberLs0.value + numberLs1.value + carry;
        if(sumVal >= 10) {
            // need to separate carry value
            // get floor value 
            carry = Math.floor(sumVal / 10);
            sumVal = sumVal % 10;
        }
        const numberLs = new NumberLinkedList(sumVal);
        numberLs.next = addTwoNumbersRecursive(numberLs0.next, numberLs1.next, carry);
        return numberLs;
    }

    lengthOfLongestSubstring(str) {
        const arrLength = str.length;
        let fast = slow = 0;
        const set = new Set();
        let result = 0;
        while(fast < arrLength && slow < arrLength) {
            const value = str[fast];
            const existingValue = set.get(value);
            if(existingValue) {
                // value exist, break the substring count
                // extend the slow range, this reset the start value
                // remove the existing value from set
                set.remove(existingValue);
                slow += slow;
            }
            else {
                // value does not exist, continue substring count
                // extend the fast range, this extend the result to next index
                // add the value to set
                // calculate result
                set.add(value);
                fast += 1;
                result = Math.max(result, (fast - slow) + 1);
            }
        }
    }

    // ineffiecient solution
    medianOfTwoSortedArray(nums1, nums2) {
        if(!nums1 || !nums2) {
            return 0.0;
        }
        const mergedArr = nums1.concat(nums2);
        mergedArr.sort(function (a, b) { return a - b; });
        const mergedArrLength = mergedArr.length;
        const midPoint = parseInt((mergedArrLength + 1) / 2);
        let median = 0.0;
        if(mergedArrLength % 2 === 0) {
            // for even length, median is avg of midPoint - 1 + midPoint elements
            // median = (mergedArr[midPoint - 1] + mergedArr[midPoint]) / 2
            median = (mergedArr[midPoint - 1] + mergedArr[midPoint]) / 2;
        }
        else {
            median = mergedArr[midPoint - 1];
        }
        return median;
    }
}