def postOrderBin(n, r):
    '''
    n개의 정점을 가지는 완전 이진 트리를 후위순회한 결과를 리스트로 반환하는 함수
    단 현재 루트는 r
    '''

    if r > n :
        return []
    result = []

    result = postOrderBin(n, r * 2)
    result = result + postOrderBin(n, r*2+1)

    result.append(r)

    return result

def completeBinTraverse(n) :

    result = []

    return postOrderBin(n, 1)