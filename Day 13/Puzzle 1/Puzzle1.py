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

rightOrderIndices = []
leftList = False
rightList = False
left = []
right = []
index = 1

while True:

    # Get next line from file
    line = file.readline()

    # if line is empty
    # end of file is reached
    if not line:
        break

    if len(line) <= 1:
        leftList = False
        rightList = False
        continue

    if not leftList:
        left = json.loads(line)
        leftList = True
    elif not rightList:
        right = json.loads(line)
        rightList = True

    if leftList and rightList:
        #print(left)
        #print(right)

        leftLen = len(left)
        rightLen = len(right)
        tmplen = max(leftLen, rightLen)
        for i in range(tmplen):
            if i >= leftLen:
                rightOrderIndices.append(index)
                break
            elif i >= rightLen:
                break

            tmpcompare = compareValues(left[i], right[i])
            if tmpcompare == -1:
                rightOrderIndices.append(index)
            if tmpcompare != 0:
                break
        index += 1

print(sum(rightOrderIndices))

file.close()
