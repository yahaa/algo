import collections


def commonChars(A):
    c = collections.Counter(A[0])
    print(c)
    for i in range(1, len(A)):
        c &= collections.Counter(A[i])
        print(c)
    return c.elements()


if __name__ == '__main__':
    print(commonChars(["abbc", "bbb", "b"]))
