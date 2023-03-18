package com.tanveershafeeprottoy.ds;

import java.util.Iterator;

public class LinkedList<T> implements Iterable<T> {
    private int size = 0;
    private Node<T> head;

    private class Node<T> {
        private T data;
        private Node<T> next;

        Node(T data, Node<T> next) {
            this.data = data;
            this.next = next;
        }
    }

    @Override
    public Iterator<T> iterator() {
        return null;
    }
}
