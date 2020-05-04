using System;

namespace csharp
{
    class Program
    {
        static void Main(string[] args)
        {
            System.Console.WriteLine(9%10);
            System.Console.WriteLine(1/10);
            //Test1();
            BinarySearchTest();
            Console.WriteLine("Hello World!");
        }

        static void Test1()
        {
            var l1 =new ListNode(1);
            var l2 =new ListNode(8);
            //l1.next=new ListNode(4);
            //l1.next.next=new ListNode(3);
            //l2.next=new ListNode(6);
            //l2.next.next=new ListNode(4);

            Sloution.AddTwoNumbers(l1,l2);
        }

        static void BinarySearchTest()
        {
            var nums =new int[]{1,4,5,7,9,10,29,78,100};
            if(nums.BinarySearchRecursive(29))
                System.Console.WriteLine("29存在");

            if(nums.BinarySearchRecursive(7))
                System.Console.WriteLine("7存在");

            if(nums.BinarySearchRecursive(4))
                System.Console.WriteLine("4存在");
            
            if(nums.BinarySearchRecursive(30))
                System.Console.WriteLine("30存在");

            if(nums.BinarySearchRecursive(100))
                System.Console.WriteLine("100存在");
        }
    }
}
