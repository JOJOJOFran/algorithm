using System;
using System.Collections.Generic;

namespace csharp
{
    public static class ArrayAlgo
    {
        //https://leetcode-cn.com/problems/two-sum/
        public static int[] TwoSum (int[] nums, int target)
        {
            var ret=new int[2];
            var dic=new Dictionary<int,int>();

            for(int i=0;i<nums.Length;i++)
            {
                if(dic.ContainsKey(nums[i]))
                {
                    ret[0]=dic[nums[i]];
                    ret[1]=i;
                }
                else
                {
                    dic.Add(target-nums[i],i);
                }
            }

            return ret;
        }
    }
}