class maxMachine :
    def __init__(self) :
        self.myData = []
    
    def addNumber(self, n) :
        self.myData.append(n)

    def removeNumber(self, n) :
        self.myData.remove(n)

    def getMax(self) :
        return max(self.myData)

def main():
    myMachine = maxMachine()
    myMachine2 = maxMachine()

    myMachine.addNumber(2)
    myMachine.addNumber(5)
    myMachine.addNumber(3)
    myMachine.addNumber(4)

    print(myMachine.getMax())

main()