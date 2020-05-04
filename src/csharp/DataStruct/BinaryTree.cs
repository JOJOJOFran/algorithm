using System;

namespace DataStruct
{
    public class BinaryTree
    {
        public TreeNode<int> Root { get { return _root; } }

        private TreeNode<int> _root;

        public BinaryTree (TreeNode<int> root)
        {
            _root = root;
        }

        public bool InsertNode (TreeNode<int> node)
        {
            if (node == null)
                return false;

            var p = _root;

            while (p != null)
            {
                if (node.Value <= _root.Value)
                {
                    if(p.Left==null)
                    {
                        p.Left=node;
                        return true;
                    }
                    else
                    {
                        p=p.Left;
                    }
                }
                else
                {
                    if(p.Right==null)
                    {
                        p.Right=node;
                        return true;
                    }
                    else
                    {
                        p=p.Right;
                    }            
                }

            }
            p=node;
            return true;
        }

        public bool DeleteNode(int value)
        {
            var p= _root;
            var tmp =p;
            var tmpValue=0;
            var flag=false;
            while(p!=null)
            {
                if(p.Value<value&&!flag)
                {
                    p=p.Right;
                }

                if(p.Value>value&&!flag)
                {
                    p=p.Left;
                }

                if(p.Value==value&&!flag)
                {
                    tmp=p;
                    flag=true;
                }

                if(flag)
                {
                    tmpValue=p.Value;
                    p=p.Left;
                }
                
            }

            return true;


        }
    }

    public class TreeNode<T>
    {
        public T Value { get; set; }

        public TreeNode<T> Left { get; set; }

        public TreeNode<T> Right { get; set; }

        public TreeNode (T data)
        {
            Value = data;
        }
    }

}