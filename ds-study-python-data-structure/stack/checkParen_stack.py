import msvcrt

class Stack:
    
    def __init__(self) :
        self.myStack = []

    def push(self, n):
        self.myStack.append(n)

    def pop(self):
        if len(self.myStack) == 0:
            return
        else:
            self.myStack.pop()
    def size(self) :
        return len(self.myStack)
        
    def empty(self) :
        if len(self.myStack) == 0 :
            return 1
        else :
            return 0

    def top(self) :
        if len(self.myStack) == 0 :
            return -1
        else :
            return self.myStack[-1]

def checkParen(p) :
    myStack = Stack()

    for x in p :
        if x == "(" :
            myStack.push(x)
        else :
            if myStack.empty == 1:
                return "NO"
            else :
                myStack.pop()
    if myStack.empty() == 1 :
        return "YES"
    else :
        return "NO"
            

def main() :
    x = input()
    print(checkParen(x))


if __name__ == "__main__":
    main()