package com.tanveershafeeprottoy.ds;

import java.util.Map;
import java.util.Stack;

public class Bracket {
    public static final char[][] BRACKETS = { { '{', '}' }, { '[', ']' }, { '(', ')' } };
    public static final Map<Character, Character> BRACKETS_MAP = Map.of(
        '{', '}',
        '[', ']',
        '(', ')'
    );

    public static boolean isOpen(char bracket) {
        for(char[] chars : BRACKETS) {
            if(chars[0] == bracket) {
                return true;
            }
        }
        return false;
    }

    public static boolean isMatched(char openBracket, char closedBracket) {
        for(char[] chars : BRACKETS) {
            if(chars[0] == openBracket) {
                return chars[1] == closedBracket;
            }
        }
        return false;
    }

    public static boolean isBalanced(String expression) {
        Stack<Character> brackets = new Stack<>();
        for(char c : expression.toCharArray()) {
            if(isOpen(c)) {
                brackets.push(c);
            }
            else if(brackets.isEmpty() || !isMatched(brackets.pop(), c)) {
                return false;
            }
        }
        return brackets.isEmpty();
    }
}
