
def is_left_bracket(bracket):
    left_brackets = ('(', '{', '[')
    if bracket == left_brackets[0] or bracket == left_brackets[1] or bracket == left_brackets[2]:
        return True
    return False


def get_left_bracket(right_bracket):
    right_left_bracket_dict = {
        ')': '(',
        '}': '{',
        ']': '['
    }
    return right_left_bracket_dict[right_bracket]


# Given a string made up with following brackets: (), {}, []
# determine whether the brackets properly match
def match_brackets(brackets: str):
    brackets_stack = list()
    for bracket in brackets:
        if is_left_bracket(bracket):
            brackets_stack.append(bracket)
        elif not brackets_stack or brackets_stack.pop() != get_left_bracket(bracket):
            return False
    if not brackets_stack:
        return True
    return False


def is_string_balanced():
    pass


print(match_brackets('[[{}]]'))
