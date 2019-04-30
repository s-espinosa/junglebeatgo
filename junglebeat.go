package main

import "fmt"

type node struct {
  data      string
  nextNode *node
}

type linkedList struct {
  head *node
}

func main() {
  list := linkedList{}

  blah := node{data: "one"}
  list  = appendList(list, blah)

  bluh := node{data: "two"}
  list  = appendList(list, bluh)

  bleh := node{data: "three"}
  list  = appendList(list, bleh)

  fmt.Println("\n\nEND")
  fmt.Println(list.head)
  fmt.Println(list.head.nextNode)
  fmt.Println(list.head.nextNode.nextNode)
}

func appendList(list linkedList, n node) linkedList {
  if list.head == nil {
    return setHead(list, n)
  } else {
    return addToEnd(list, n)
  }
  // return list
}

func setHead(list linkedList, n node) linkedList {
  list.head = &n
  return list
}

func addToEnd(list linkedList, n node) linkedList {
  currentNode := list.head
  for currentNode.nextNode != nil {
    currentNode = currentNode.nextNode
  }
  currentNode.nextNode = &n
  return list
}

// func main() {
//   data := "Bleh"
//   list := linkedList{}
//   list = prependList(list, data)
//   fmt.Println(list.head)
//   data = "bluh"
//   list = prependList(list, data)
//   fmt.Println(list.head)
//   fmt.Println(list.head.nextNode)
// }
//
// func prependList(l linkedList, d string) linkedList {
//   if l.head == nil {
//     n := node{data: d}
//     l.head = &n
//   } else {
//     // n := node{data: d}
//     nn := l.head
//     nnAddress := &nn
//     n := node{data: d, nextNode: nnAddress}
//     l.head = &n
//   }
//   return l
// }
