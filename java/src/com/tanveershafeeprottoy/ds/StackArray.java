package com.tanveershafeeprottoy.ds;

import java.util.Arrays;

public class StackArray<T> {
    private Object[] objArray;
    private int top;

    public StackArray(int size) {
        objArray = new Object[size];
        top = -1;
    }

    public void push(T obj) {
        if(top >= objArray.length - 1) {
            increaseCapacity();
        }
        objArray[++top] = obj;
    }

    @SuppressWarnings("unchecked")
    public T pop() throws Exception {
        if(isEmpty()) {
            throw new Exception("Stack underflow error");
        }
        T obj = (T) objArray[top];
        objArray[top--] = null;
        return obj;
    }

    public boolean isEmpty() {
        return top < 0;
    }

    public void increaseCapacity() {
        int newLength = objArray.length * 2;
        objArray = Arrays.copyOf(objArray, newLength);
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append('[');
        int length = objArray.length;
        for(int i = 0; i < length; i++) {
            sb.append(objArray[i].toString());
            if(i < length - 1) {
                sb.append(",");
            }
        }
        sb.append(']');
        return sb.toString();
    }
}
