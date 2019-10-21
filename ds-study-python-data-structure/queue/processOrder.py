
# 내가 만든 queue
class Queue:
    def __init__(self) :
        self.myQueue = []
        
    def push(self, n):
        self.myQueue.append(n)

    def pop(self):
        return self.myQueue.pop()

    def front(self):
        return self.myQueue[0]

    def empty(self):
        if self.myQueue[0] :
            return 0
        else :
            return 1

def selectQueue(normalQueue, vipQueue, jobEndTime, orders) :
    if normalQueue.empty() == 1 :
        return vipQueue
    
    if vipQueue.empty() == 1 :
        return normalQueue
    
    if orders[normalQueue.front()].time <= jobEndTime or
orders[vipQueue.front()].time <= jobEndTime :
        if orders[vipQueue.front()].time <= jobEndTime :
            return vipQueue
        else :
            return normalQueue
    else :
        if orders[normalQueue.front()].time <= orders[vipQueue.front()].time :
            return normalQueue
        else :
            return vipQueue


def processOrder(orders) :

    result = []
    
    normalQueue = Queue()
    vipQueue = Queue()

    for i in range(len(orders)) :
        if orders[i].vip == 1 :
            vipQueue.push(i)
        else :
            normalQueue.push(i);

    jobEndTime = 0

    while not (normalQueue.empty() == 1 and vipQueue.empty() == 1) :
        # normalQueue에서 하나를 빼야하는지, 혹은 vipQueue에서 하나를 빼야하는지
        # 알려주는 함수. 즉 다음에 무슨일을 해야하는지 알려주는 함수를 작성
        targetQueue = selectQueue(normalQueue, vipQueue, jobEndTime, orders)

        idx = targetQueue.pop()

        jobEndTime = max(jobEndTime, orders[idx].time) + orders[idx].duration
        
        result.append(idx);
    

    return result