class priorityQueue:

    def __init__(self) :
        self.data = [0]
    
    def push(self, value) :
        '''
        우선순위 큐에 value를 삽입합니다.
        '''

        self.data.append(value)
        inx = len(self.data)-1
        
        while inx != 1 and self.data[inx] < self.data[inx//2] :
            temp = self.data[inx]
            self.data[inx] = self.data[inx//2]
            self.data[inx//2] = temp

            inx = inx // 2

        pass

    def top(self) :
        '''
        우선순위가 가장 높은 원소를 반환합니다.
        '''
        if len(self.data) <=1 :
            return -1

        return self.data[1]

    def pop(self) :
        '''
        우선순위가 가장 높은 원소를 제거합니다.
        '''

        if len(self.data) <= 1:
            return 

        self.data[1] = self.data[-1]
        self.pop()

        inx = 1

        while True :
            min_idx = -1 # 내가 내려갈 정점의 번호

            if inx*2 > len(self.data)-1 and inx*2+1 > len(self.data)-1:
                break
            elif inx*2+1 > len(self.data)-1 :
                min_idx = inx*2

            else :
                if self.data[inx*2] > self.data[inx*2+1] :
                    min_idx = 1nx*2+1
                else :
                    min_idx = inx*2

            if self.data[inx] > self.data[min_idx] :
                temp = self.data[inx]
                self.data[inx] = self.data[min_ix]
                self.data[min_idx] = temp

                inx = min_idx

            else:
                break


