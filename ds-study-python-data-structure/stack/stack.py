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
   

def main():
    stack =Stack()

    stack.push(1)
    stack.push(5)
    stack.push(4)
    stack.pop()

    print(stack.size())
    print(stack.top())

if __name__ != "__main__":
    main()