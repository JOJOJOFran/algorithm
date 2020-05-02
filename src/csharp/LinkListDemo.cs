using System;

namespace csharp
{
    public class ListNode
    {
        public int val;
        public ListNode next;
        public ListNode (int x) { val = x; }
    }

    public static class Sloution
    {
        //https://leetcode-cn.com/problems/add-two-numbers/
        public static ListNode AddTwoNumbers (ListNode l1, ListNode l2)
        {
            int add = 0;
            int newVal = 0;
            int sum = 0;
            var ret = new ListNode (0);
            var p = ret;

            while (l1 != null || l2 != null)
            {

                sum = ((l1?.val??0) + (l2?.val??0) + add);
                newVal = sum % 10;
                add = sum / 10;

                p.val = newVal;
                if (l1?.next != null || l2?.next != null || add > 0)
                {
                    p.next = new ListNode (add);
                    p = p.next;
                }

                l1 = l1?.next;
                l2 = l2?.next;
            }

            return ret;
        }
    }

}