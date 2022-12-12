/*
    approach: 
    - To represent snake, we will use a linked list
        - why? because when a snake moves from cell to other
        - what data structure helps us add to end the in o(1) time and remove from the front in o(1) time
        - linkedlist
        - using an array would be more time expensive,
        - LL head will be the tail of snake and LL tail will be head of the snake
        - When we move snake into a new empty cell, then we will add a new node to tail and update tail to point to new node and we will move our head pointer in LL to head.next and remove the old head node -- when a snake moves, if its not growing, its head moves to new cell and its tail follows therefore the above action..
    - Then when we are trying to move to a direction, we will verify a few things
        - is the new cell out of bounds? if yes, return -1, game over
        - if the new cell already occupied by snake body ?
            - if the new cell is occupied by snake tail, thats acceptable beacause when the snake has fully moved, its tail wont be new cell so it wont clash
            - however if its some other middle part of the body, return -1, game over
        - if the new cell has food, then eat the food by
            - add new cell position to LL
        - if the new cell is a blank cell and none of the above rules were met, then
            - add new cell position to snake LL
            - remove the head of the snake LL to fully move the snake
            - Head of LL is the tail of the snake
            - Tail of LL is the head of the snake
    - All of that is fine, what about the scores.... 
        - The size of the snake - 1 is the score we will return after processing a input key stroke
            - Why the -1 ?
            - because the snake LL will ALWAYS start 1 size long, it has a start position of 0,0
            - so any new node we add, will increase the size by 1
            - but when the game ends, we need to remove the starting snake node from the final size of the LL
            - Therefore the -1
*/

type SnakeGame struct {
    foodQ *ll
    snake *ll
    m, n int
}


func Constructor(width int, height int, food [][]int) SnakeGame {
    foodQ := new(ll)
    for _, pos := range food {
        foodQ.addToTail(pos[0], pos[1])
    }
    snake := new(ll)
    snake.addToTail(0,0)
    return SnakeGame{
        foodQ: foodQ,
        snake: snake,
        m : height,
        n : width,
    }
}


func (this *SnakeGame) Move(direction string) int {
    // current snake head position
    snakeRow := this.snake.tail.r
    snakeCol := this.snake.tail.c
    
    // apply the direction change to snakeRow and snakeCol
    if direction == "U" {
        snakeRow--
    } else if direction == "D"{
        snakeRow++
    } else if direction == "L" {
        snakeCol--
    } else if direction == "R" {
        snakeCol++
    }
    
    // now verify new snake head position
    
    // end the game if the new snake is going out of matrix
    if snakeRow == this.m || snakeRow < 0 || snakeCol == this.n || snakeCol < 0 {
        return -1
    }
    
    // end the game if the new snake head hits one of its body nodes, excluding its tail ( or head of LL in this case )
    curr := this.snake.head.next
    for curr != nil {
        if snakeRow == curr.r && snakeCol == curr.c {
            return -1
        }
        curr = curr.next
    }
    
    // otherwise
    // if the new cell has food, make the snake bigger
    if this.foodQ.size != 0 {
        foodRow := this.foodQ.head.r
        foodCol := this.foodQ.head.c
        if snakeRow == foodRow && snakeCol == foodCol {
            this.foodQ.removeHead()
            this.snake.addToTail(snakeRow, snakeCol)
            return this.snake.size-1
        }
    }
    
    // if no food, simply move by adding new position to tail end of the LL
    // and removing head of the LL
    this.snake.addToTail(snakeRow, snakeCol)
    this.snake.removeHead()
    return this.snake.size-1
}


type listNode struct {
    r int
    c int
    next *listNode
}

type ll struct {
    head *listNode
    tail *listNode
    size int
}

func (l *ll) addToTail(r,c int) {
    newNode := &listNode{r: r, c: c}
    if l.head == nil {
        l.head = newNode
        l.tail = newNode
        l.size = 1
        return
    }
    l.tail.next = newNode
    l.tail = newNode
    l.size++
}


func (l *ll) removeHead() {
    if l.head == nil {
        return
    }
    if l.head.next == nil {
        l.head = nil
        l.tail = nil
        l.size = 0
        return
    }
    newHead := l.head.next
    l.head.next = nil 
    l.head = newHead
    l.size--
}
/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */