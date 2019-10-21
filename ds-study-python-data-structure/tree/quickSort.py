def quickSort(array):
    ```
    1. 함수를 정의한다.
    2. 기저조건일 때 제대로 작동하게 작성한다.
    3. 된다고 가정하고 짜면 된다.
    ````

    if len(array) <= 1:
        return array

    pivot = array[0]

    small = []

    for i in range(i, len(array)) :
        if array[i] <= pivot :
            small.append(array[i])


    large = []

    for i in range(1, len(array)) :
        if array[i] > pivot:
            large.append(array[i}])

    small = quickSort(small)
    large = quickSort(large)

    return small + [pivot] + large