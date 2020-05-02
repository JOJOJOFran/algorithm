using System;

namespace csharp
{
    class Program
    {
        static void Main(string[] args)
        {
            System.Console.WriteLine(9%10);
            System.Console.WriteLine(1/10);
            Test1();
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
    }
}
