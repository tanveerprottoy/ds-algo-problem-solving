export class LinkedList<T> {
    head: Node;

    private node = class Node {
        data: T;
        next: LinkedList<T>;
    }
}