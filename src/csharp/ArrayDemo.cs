using System;
using System.Collections.Generic;

namespace csharp
{
    public static class ArrayAlgo
    {
        //https://leetcode-cn.com/problems/two-sum/
        public static int[] TwoSum (int[] nums, int target)
        {
            var ret = new int[2];
            var dic = new Dictionary<int, int> ();

            for (int i = 0; i < nums.Length; i++)
            {
                if (dic.ContainsKey (nums[i]))
                {
                    ret[0] = dic[nums[i]];
                    ret[1] = i;
                }
                else
                {
                    dic.Add (target - nums[i], i);
                }
            }

            return ret;
        }

        #region 二分查找
        public static bool BinarySearch (this int[] nums, int target)
        {
            if(nums==null||nums.Length<=0)
                return false;

            var start = 0;
            var end = nums.Length - 1;
            var mid = start + ((end - start) >> 1);

            while (start < end)
            {
                if (nums[mid] < target)
                    start = mid + 1;

                if (nums[mid] > target)
                    end = mid-1;

                if (nums[mid] == target)
                    return true;

                mid = start + (end - start) >> 1;
            }

            return nums[mid] == target;
        }

        public static bool BinarySearchRecursive (this int[] nums, int target)
        {
            if(nums==null||nums.Length<=0)
                return false;
            return BinarySearchPartion(nums,0,nums.Length,target);
        }

        private static bool BinarySearchPartion (int[] nums, int start, int end, int target)
        {
            var mid = start + ((end - start) >>1);
            while (start <end)
            {
                if (nums[mid] < target)
                    start = mid + 1;

                if (nums[mid] > target)
                    end = mid-1;

                if (nums[mid] == target)
                    return true;
                return BinarySearchPartion (nums, start, end, target);
            }

             return nums[mid] == target;

        }
        #endregion

    }
}