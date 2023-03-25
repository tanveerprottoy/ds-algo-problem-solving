export class Stack<T> {
    data: T[];

    constructor() {
        this.data = [];
    }

    push(item: T) {
        this.data.push(item)
    }

    pop(): T {
        if(this.data.length === 0) {
            throw new Error("stack is empty");
        }
        return this.data.pop();
    }
}