class myDatabase :
    def __init__(self, size):
        # 크기 size의 해쉬 테이블을 만듭니다.
        self.myData = [(-1, -1) for i in range(size)]

    def put(self, key, value) :
        index = self.hashFunction(key)

        for i in range(len(self.myData)):
            if self.myData[index][0] == -1 :
                self.myData[index] = (key, value)
                return
            index = (index + 1) % len(self.myData)
        

    def get(self, key):
        index = self.hashFunction(key)

        for i in range(len(self.myData)) :
            if self.myData[index][0] == key:
                return self.myData[index][1]

            if self.myData[index][0] == -1 :
                return -1
            index = (index + 1) % len(self.myData)
        return -1

    def hashFunction(self, key):
        return key % len(self.myData);


def main() :
    db = myDatabase(100)

    db.put(1, 3)
    db.put(2, 7)
    db.put(3, 8)
    db.put(403, 9)

    print(db.get(3))
    print(db.get(403))


if __name__ == "__main__":
    main()