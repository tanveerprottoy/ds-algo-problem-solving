export class NumberLinkedList {
    value
    next

    constructor(value, next = null) {
        this.value = value;
        this.next = next;
    }
}