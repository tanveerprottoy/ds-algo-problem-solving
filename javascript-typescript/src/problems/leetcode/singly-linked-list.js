export class SinglyLinkedList {
    value
    next

    constructor(value, next = null) {
        this.value = value;
        this.next = next;
    }
}