def countPrimes(t):
    result = [ 0 for i in range(t+1)]

    cnt = 0

    for i range(2, t+1):
        if result[i] == 0 :
            cnt = cnt + 1
            for j in range(i+i, t+1, i) :
                result[j] = 1

    return cnt

    # 1,3,4,5,9,10