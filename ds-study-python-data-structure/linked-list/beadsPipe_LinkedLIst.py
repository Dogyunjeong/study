class LinkedListElement :
    def __init__(self) :
        self.value = None
        self.myNext = None

class LinkedLIstPipe :
    def __init__(self) :
        self.start = None
        self.end = None
    
    def addLeft(self, n) :
        if self.start == None :
            element = LinkedListElement()
            element.value = n
            element.myNext = None
            self.start = element
            self.end = element
        else :
            element = LinkedListElement()
            element.value = n
            element.myNext = self.start
            self.start = element

    def addRight(self, n) :
        if self.start == None :
            element = LinkedListElement()
            element.value = n
            element.myNext = None
            self.start = element
            self.end = element
        else :
            element = LinkedListElement()
            element.value = n
            element.myNext = None

            self.end.myNext = element
            self.end = element
        
    def getBeads(self) :
        result = []

        cur = self.start
        while cur != None :
            result.append(cur.value)
            cur = cur.myNext

        return result
        

def processBeads(myInput) :
    myPipe = LinkedLIstPipe()

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