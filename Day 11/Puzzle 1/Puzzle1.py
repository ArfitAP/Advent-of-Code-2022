
class Monkey(object):
    # Initializing to None
    def __init__(self):
        self.number = 0
        self.items = []
        self.operation = ""
        self.testDivisable = 0
        self.trueAction = 0
        self.falseAction = 0


def round():

    for i in range(len(monkies)):

        monkey = monkies[i]
        while len(monkey.items) > 0:

            item: int = monkey.items.pop()
            itemInspections[i] += 1

            operator = monkey.operation[4]

            leftop, rightop = monkey.operation.split(operator)

            if leftop.strip() == "old":
                left = item
            else:
                left = int(leftop.strip())
            if rightop.strip() == "old":
                right = item
            else:
                right = int(rightop.strip())

            newvalue = 0
            if operator == "+":
                newvalue = left + right
            elif operator == "*":
                newvalue = left * right

            newvalue = int(newvalue / 3)

            if newvalue % monkey.testDivisable == 0:
                throwTo(newvalue, monkey.trueAction)
            else:
                throwTo(newvalue, monkey.falseAction)


def throwTo(item, monkey):
    monkies[monkey].items.insert(0, item)


file = open('../input.txt', 'r')

monkies = []
itemInspections = []

while True:

    # Get next line from file
    line = file.readline()

    # if line is empty
    # end of file is reached
    if not line:
        break

    if line.split(" ")[0] == "Monkey":
        num = int(line.split(" ")[1].split(":")[0])
        tmpMonkey = Monkey()
        tmpMonkey.number = num

        startItems = file.readline()
        for item in startItems.split(":")[1].split(","):
            tmpMonkey.items.insert(0, int(item.strip()))

        operation = file.readline()
        tmpMonkey.operation = operation.split("=")[1].strip()

        test = file.readline()
        tmpMonkey.testDivisable = int(test.split("by")[1].strip())

        trueAction = file.readline()
        tmpMonkey.trueAction = int(trueAction.split("monkey")[1].strip())

        falseAction = file.readline()
        tmpMonkey.falseAction = int(falseAction.split("monkey")[1].strip())

        monkies.append(tmpMonkey)
        itemInspections.append(0)


for i in range(20):
    round()

itemInspections.sort(reverse=True)
print(itemInspections[0] * itemInspections[1])

file.close()
