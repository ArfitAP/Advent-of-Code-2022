from functools import cmp_to_key
import json
import collections.abc

def compareValues(left, right):
    if not isinstance(left, collections.abc.Sequence) and not isinstance(right, collections.abc.Sequence):
        if left < right:
            return -1
        elif left > right:
            return 1
        else:
            return 0

    elif isinstance(left, collections.abc.Sequence) and isinstance(right, collections.abc.Sequence):

        leftLen = len(left)
        rightLen = len(right)
        tmplen = max(leftLen, rightLen)
        for i in range(tmplen):
            if i >= leftLen:
                return -1
            elif i >= rightLen:
                return 1

            tmpcompare = compareValues(left[i], right[i])
            if tmpcompare == -1:
                return -1
            elif tmpcompare == 1:
                return 1
        return 0
    elif not isinstance(left, collections.abc.Sequence) and isinstance(right, collections.abc.Sequence):
        newleft = [left]
        return compareValues(newleft, right)
    elif isinstance(left, collections.abc.Sequence) and not isinstance(right, collections.abc.Sequence):
        newright = [right]
        return compareValues(left, newright)


file = open('../input.txt', 'r')

allPacks = []

while True:

    # Get next line from file
    line = file.readline()

    # if line is empty
    # end of file is reached
    if not line:
        break

    if len(line) <= 1:
        continue

    allPacks.append(json.loads(line))

firstElem = json.loads("[[2]]")
secondElem = json.loads("[[6]]")
allPacks.append(firstElem)
allPacks.append(secondElem)

allPacks.sort(key=cmp_to_key(compareValues))

print((allPacks.index(firstElem) + 1) * (allPacks.index(secondElem) + 1))

file.close()
