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

    /*
    Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

    An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.
    */
    isValidParentheses(s) {
        const bracketsMap = new Map([
            ["(", ")"],
            ["{", "}"],
            ["[", "]"],
        ]);
        const brackets = new Array();
        for(const c of s) {
            // check if it's open
            const closedBracket = bracketsMap.get(c);
            if(closedBracket) {
                // c == open bracket
                // push to stack
                brackets.push(c);
            }
            else if(brackets.length === 0) {
                // c == closed bracket
                // but no other open
                // bracket was pushed
                // so it's invalid
                return false;
            }
            else {
                // c == closed bracket
                // stack has open brackets
                // pop and compare the closed brackets
                // if they match then it's good
                // to continue
                // if dont match then return false
                const openBracket = brackets.pop();
                const closed = bracketsMap.get(openBracket);
                if(c !== closed) {
                    return false;
                }
            }
        }
        // return brackets is empty
        return brackets.length === 0;
    }

    /*
    Given an integer array nums and an integer val, remove all occurrences of val 
    in nums in-place. The relative order of the elements may be changed.
    Since it is impossible to change the length of the array in some languages, 
    you must instead have the result be placed in the first part of the array nums. 
    More formally, if there are k elements after removing the duplicates, 
    then the first k elements of nums should hold the final result. 
    It does not matter what you leave beyond the first k elements.
    Return k after placing the final result in the first k slots of nums.
    Do not allocate extra space for another array. You must do this by modifying 
    the input array in-place with O(1) extra memory.

    Example 1:

    Input: nums = [3,2,2,3], val = 3
    Output: 2, nums = [2,2,_,_]
    Explanation: Your function should return k = 2, with the first two elements of 
    nums being 2.
    It does not matter what you leave beyond the returned k 
    (hence they are underscores).
    */
    removeElement(nums, val) {
        // 2 pointers solution
        if(nums.length === 0) {
            return 0;
        }
        // 2nd pointer
        // will track the index that will be overwritten
        // during a loop
        let k = 0;
        for(const n of nums) {
            // this algo will move the
            // non val values (n !== val)
            // to k index (2nd pointer)
            if(n !== val) {
                // move nums[i] to k
                nums[k] = n;
                k++;
            }
            // do not need to move n === val
        }
        console.log(nums)
        return k;
    }

    /*
    Given an integer array nums sorted in non-decreasing order, 
    remove the duplicates in-place such that each unique element appears only once.
    The relative order of the elements should be kept the same.
    Since it is impossible to change the length of the array in some languages, 
    you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
    Return k after placing the final result in the first k slots of nums.
    Do not allocate extra space for another array. You must do this by modifying 
    the input array in-place with O(1) extra memory.
    
    Example 1:
    Input: nums = [1,1,2]
    Output: 2, nums = [1,2,_]
    Explanation: Your function should return k = 2, with the first two elements
    of nums being 1 and 2 respectively.
    It does not matter what you leave beyond the returned k 
    (hence they are underscores).
    */
    removeDuplicates(nums) {
        if(nums.length === 0) {
            return 0;
        }
        // this algo will not use any
        // auxilary ds, instead it will
        // start from index 1 and check with the
        // previous/adjacent item if they are same or not
        // 2nd pointer, also starts from 1
        let k = 1;
        for(let i = 1; i < nums.length; i++) {
            // need to only rearrange
            // the non duplicate ones
            // as iterating need to check with
            // the previous one
            if(nums[i] !== nums[i - 1]) {
                // if diff than previous/adjacent
                // then rearrange
                nums[k] = nums[i];
                k++;
            }
        }
        console.log(nums);
        return k;
    }

    removeDuplicatesSet(nums) {
        if(nums.length === 0) {
            return 0;
        }
        // 2nd pointer
        let k = 0;
        // need a ds to store duplicates
        // set will be a good one for this
        const numSet = new Set();
        for(const n of nums) {
            // need to only rearrange
            // the non duplicate ones
            if(!numSet.has(n)) {
                nums[k] = n;
                k++;
                numSet.add(n);
            }
        }
        console.log(nums);
        return k;
    }

    /*
    Given an integer array nums sorted in non-decreasing order, 
    remove some duplicates in-place such that each unique element 
    appears at most twice. The relative order of the elements should be kept the same.
    Since it is impossible to change the length of the array in some languages, 
    you must instead have the result be placed in the first part of the array nums. 
    More formally, if there are k elements after removing the duplicates, then the 
    first k elements of nums should hold the final result. It does not matter what 
    you leave beyond the first k elements.
    Return k after placing the final result in the first k slots of nums.
    Do not allocate extra space for another array. You must do this by modifying 
    the input array in-place with O(1) extra memory.

    Example 1:
    Input: nums = [1,1,1,2,2,3]
    Output: 5, nums = [1,1,2,2,3,_]
    Explanation: Your function should return k = 5, with the first five elements 
    of nums being 1, 1, 2, 2 and 3 respectively.
    It does not matter what you leave beyond the returned k 
    (hence they are underscores).
    */
    removeDuplicatesOverTwo(nums) {
        if(nums.length === 0) {
            return 0;
        }
    }
}