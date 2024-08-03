package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//  func reorderList(head *ListNode)  {
//     if head == nil {
//         return 
//     }

//     middle := midnode(head)
//     newhead := middle.Next
//     middle.Next = nil

//     newhead = reverseLinkedList(newhead)

//     c1 := head
//     c2 := newhead

//     var f1,f2 *ListNode

//     for c1 != nil && c2 != nil{
//         f1 = c1.Next
//         f2 = c2.Next

//         c1.Next = c2
//         c2.Next = f1

//         c1 = f1
//         c2 = f2
//     } 
// }

// func midnode(head *ListNode) *ListNode {
//     slow := head
//     fast := head

//     for fast.Next != nil && fast.Next.Next != nil {
//         slow = slow.Next
//         fast = fast.Next.Next
//     }

//     return slow
// }

// func reverseLinkedList(head *ListNode) *ListNode {
//     var prev *ListNode
//     current := head
//     for current != nil {
//         nxt := current.Next
//         current.Next = prev
//         prev = current
//         current = nxt
//     }

//     return prev
// }