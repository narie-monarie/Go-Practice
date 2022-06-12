func printLinkedList(head *SinglyLinkedListNode) {
    p := head
    
    for{
        if p == nil{
            break
        }
        fmt.Printf("%d \n", p.data)
        p = p.next
    }
}
