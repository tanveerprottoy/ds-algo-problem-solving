package com.tanveershafeeprottoy.ds;

public class LinkedList<T> {

    private class Node<T> {
        private T data;
        private Node<T> next;

        Node(T data, Node<T> next) {
            this.data = data;
            this.next = next;
        }
    }

    public boolean hasCycle(Node<T> head) {
        if(head == null) return false;
        // fast starts from 2nd
        Node<T> fast = head.next;
        Node<T> slow = head;
        while(fast != null && fast.next != null && slow != null) {
            if(fast == slow) {
                return true;
            }
            // move forward
            // fast goes two steps
            fast = fast.next.next;
            slow = slow.next;
        }
        return false;
    }
}
