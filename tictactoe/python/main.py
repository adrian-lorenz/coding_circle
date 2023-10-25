def checkWin(n, coordinates):
    rowCount = [0] * n
    colCount = [0] * n
    diagonalCount1 = 0
    diagonalCount2 = 0

    for coord in coordinates:
        x, y = coord
        #sum col / row
        rowCount[x] += 1
        colCount[y] += 1
        # check col / row
        if rowCount[x] == n or colCount[y] == n:
            return True
        # sum / check diagonale1
        if x == y:
            diagonalCount1 += 1
            if diagonalCount1 == n:
                return True
        # sum / check diagonale2
        if x + y == n - 1:
            diagonalCount2 += 1
            if diagonalCount2 == n:
                return True

    return False

def drawBoard(n, coordinates):
    # init leeres board mit _
    board = [[" " for _ in range(n)] for _ in range(n)]
    # fülle board mit X an den koordinaten
    for coord in coordinates:
        x, y = coord
        board[x][y] = "X"

    for row in board:
        print(row)
        
    

def main():
    n = 3
    coordinates = [[0, 0], [0, 2], [1, 1], [2, 2]]
    drawBoard(n, coordinates)
    result = checkWin(n, coordinates)
    print(result)

if __name__ == "__main__":
    main()
