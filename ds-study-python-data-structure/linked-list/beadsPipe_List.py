class ListPipe() :
    def __init__(self) :
        self.myPipe = []

    def addLeft(self, n) :
        self.myPipe.insert(0, n)

    def addRight(self, n) :
        self.myPipe.append(n)
        
    def getBeads(self) :
        return self.myPipe

def processBeads(myInput) :
    myPipe = ListPipe()

    for line in myInput :
        if line[1] == 0 :
            myPipe.addLeft(line[0])
        else :
            myPipe.addRight(line[0])

    return myPipe.getBeads()


def main():
    n = int(input())
    myList = []

    for i in range(n) :
        myList.append([int(v) for v in input().split()])

    print(processBeads(myList))


if __name__ == "__main__":
    main()
