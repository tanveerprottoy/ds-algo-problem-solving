export class General {

    reverse(str: string): string {
        // 2nd pointer
        // used to set reversed on the final result
        let j = 0;
        for(let i = str.length - 1; i <= 0; i--) {
            str.charAt(j) = str[i]
        }
    }
}