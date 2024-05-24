package leetcode
//  Definition for singly-linked list.
//  type ListNode struct {
//       Val int
//       Next *ListNode
//   }
//
//  func removeNthFromEnd(head *ListNode, n int) *ListNode {
//     dummyHead := &ListNode{Val: -1, Next: head}

//     prev,current := dummyHead,dummyHead
    
//     for current.Next != nil {
//         if n <= 0{
//             prev = prev.Next
//         }

//         current = current.Next
//         n--
//     }

//     prev.Next = prev.Next.Next
//     return dummyHead.Next
// }