package main

import (
    "fmt"
)


//RedBlackTree adapted from: http://en.literateprograms.org/Red-black_tree_(Java)
type RedBlackTree struct {
        root *RedBlackNode
}

//Represents a node in the RedBlackTree
type RedBlackNode struct {
    parent,left,right *RedBlackNode
    value int
    color int
}

const BLACK int = 1
const RED int = 0

func verifyInvariants(root *RedBlackNode) {
    verifyPropertyOne(root)
    verifyPropertyTwo(root)
}

func verifyPropertyOne(root *RedBlackNode) {
    if root.color != BLACK {
        panic("Invariant broken, root was not BLACK")
    }
}

func verifyPropertyTwo(node *RedBlackNode) {
    if RED == node.color {
        if node.left.color != BLACK {
            panic("Invariant broken:Node color is red, current node left was not BLACK")
        }
        if node.right.color != BLACK {
            panic("Invariant broken:Node color is red, current node right was not BLACK")
        }
        if node.parent.color != BLACK {
            panic("Invariant broken:Node color is red, current node parent was not BLACK")
        }
    }
    if node == nil {
        return
    }
    verifyPropertyTwo(node.left)
    verifyPropertyTwo(node.right)
}

//Remove attempts to remove a RedBlackNode from the RedBlackTree
func (tree *RedBlackTree) Remove(node *RedBlackNode) {
    if nil != node.left && nil != node.right {
        var toReplace = tree.MaximumNode(node)
        node = toReplace
    }
    var child *RedBlackNode
    if nil==node.right {
        child=node.left
    }else{
        child=node.right
    }
    if node.color == BLACK {
        node.color = child.color
        tree.deleteOne(node)
    }
    tree.replaceNode(node,child)
    if RED == tree.root.color {
        tree.root.color = BLACK
    }
}

func (tree *RedBlackTree) deleteOne(node *RedBlackNode) {
    if nil==node.parent {
        return
    }else{
    tree.deleteTwo(node)
    }
}

func (tree *RedBlackTree) deleteTwo(node *RedBlackNode) {
    if RED == tree.sibling(node).color {
        node.parent.color = RED
        tree.sibling(node).color = BLACK
        if node == node.parent.left {
            tree.rotateLeft(node.parent)
        }else{
            tree.rotateRight(node.parent)
        }
    }
    tree.deleteThree(node)
}

func (tree *RedBlackTree) deleteThree(node *RedBlackNode) {
    if BLACK == node.parent.color && tree.sibling(node).color == BLACK && tree.sibling(node).left.color == BLACK && tree.sibling(node).right.color == BLACK{
        tree.sibling(node).color=RED
        tree.deleteOne(node.parent)
    }else{
        tree.deleteFour(node)
    }
}

func (tree *RedBlackTree) deleteFour(node *RedBlackNode) {
    if RED == node.parent.color && tree.sibling(node).color == BLACK && tree.sibling(node).left.color == BLACK && tree.sibling(node).right.color == BLACK{
        tree.sibling(node).color = RED
        node.parent.color = BLACK
    }else{
        tree.deleteFive(node)
    }
}

func (tree *RedBlackTree) deleteFive(node *RedBlackNode) {
    if node == node.parent.left && tree.sibling(node).color == BLACK && tree.sibling(node).left.color == RED && tree.sibling(node).right.color == BLACK {
        tree.sibling(node).color = RED
        tree.sibling(node).left.color = BLACK
        tree.rotateRight(tree.sibling(node))
    }else if node == node.parent.right && tree.sibling(node).color == BLACK && tree.sibling(node).right.color == RED && tree.sibling(node).left.color == BLACK {
            tree.sibling(node).color = RED
            tree.sibling(node).right.color = BLACK
            tree.rotateLeft(tree.sibling(node))
    }
    tree.deleteSix(node)
}

func (tree *RedBlackTree) deleteSix(node *RedBlackNode) {
    tree.sibling(node).color=node.parent.color
    node.parent.color=BLACK
    if node == node.parent.left {
        tree.sibling(node).right.color=BLACK
        tree.rotateLeft(node.parent)
    }else{
        tree.sibling(node).left.color=BLACK
        tree.rotateRight(node.parent)
    }
}

//MaximumNode searches for the largest node in the tree
func (tree RedBlackTree) MaximumNode(node *RedBlackNode) *RedBlackNode {
    var toReturn *RedBlackNode
    for node.right != nil {
        toReturn = node.right
    }
    return toReturn
}

//Insert attempts to Insert a new node into the tree
func (tree *RedBlackTree) Insert(node *RedBlackNode){
    node.color = RED
    fmt.Println("inserting")
    fmt.Println(node)
    if tree.root == nil{
        fmt.Println("New root")
        tree.root = node
    }else{
        var temp *RedBlackNode = tree.root
        for true {
            var compared int = compareTo(node.value,temp.value)
            if compared == 0 {
                fmt.Println("Equal to")
                temp.value = node.value
                break
            }else if compared < 0 {
                fmt.Println("Less than")
                if temp.left == nil {
                    temp.left = node
                    break
                }else{
                    temp = temp.left
                }
            }else{
                fmt.Println("Greater than")
                if temp.right == nil {
                    temp.right = node
                    break
                }else{
                    temp=temp.right
                }
            }
        }
        node.parent = node
    }
}

//Search attempts to find the node that contains the supplied value
func (tree *RedBlackTree) Search(value int) *RedBlackNode{
    var current *RedBlackNode = tree.root
    for current != nil {
        var comparison int=compareTo(value,current.value)
        if 0 == comparison {
            fmt.Println("Equal Search")
            return current
        }else if comparison < 0 {
            fmt.Println("Less than Search")
            current=current.left
        }else{
            fmt.Println("Greater than Search")
            current=current.right
        }
    }
    return current
}

func (tree *RedBlackTree) nodeColor(node *RedBlackNode) int{
    if nil==node {
        return BLACK
    }else{
        return node.color
    }
}

func (tree *RedBlackTree) verifyTreeinsert(node *RedBlackNode) {
   tree.insertCheckOne(node)
}

func (tree *RedBlackTree) insertCheckOne(node *RedBlackNode) {
    if node.parent == nil {
        node.color = BLACK
        return
    }
    tree.insertCheckTwo(node)
}
func (tree *RedBlackTree) insertCheckTwo(node *RedBlackNode) {
    if  node.parent.color == BLACK{
            return
        }else{
            tree.insertCheckThree(node)
        }
}

func (tree *RedBlackTree) insertCheckThree(node *RedBlackNode){
    if tree.uncle(node).color == RED {
        node.parent.color = BLACK
        tree.uncle(node).color = BLACK
        tree.grandParent(node).color = RED
        tree.insertCheckOne(node)
    }else{
        tree.insertCheckFour(node)
    }
}

func (tree *RedBlackTree) insertCheckFour(node *RedBlackNode){
    if node == node.parent.right && node.parent == tree.grandParent(node).left {
        tree.rotateLeft(node.parent)
        node = node.left
    }else if node == node.parent.left && node.parent == tree.grandParent(node).right {
        tree.rotateRight(node.parent)
        node = node.right
    }
    tree.insertCheckFive(node)
}

func (tree *RedBlackTree) insertCheckFive(node *RedBlackNode){
    node.parent.color = BLACK
    tree.grandParent(node).color = RED
    if node == node.parent.left && node.parent == tree.grandParent(node).left {
        tree.rotateRight(tree.grandParent(node))
    }else {
        tree.rotateLeft(tree.grandParent(node))
    }
}

func (tree *RedBlackTree) rotateLeft(node *RedBlackNode) {
    var temp *RedBlackNode = node.right
    tree.replaceNode(node,temp)
    node.right=temp.left
    if temp.left != nil {
        temp.left.parent = node
    }
    temp.left=node
    node.parent = temp
}

func (tree *RedBlackTree) rotateRight(node *RedBlackNode) {
    var temp *RedBlackNode = node.left
    tree.replaceNode(node,temp)
    node.left=temp.right
    if temp.right != nil {
        temp.right.parent = node
    }
    temp.right=node
    node.parent = temp
}

func (tree *RedBlackTree) replaceNode(previousNode,newNode *RedBlackNode) {
    if previousNode.parent == nil {
        tree.root = newNode
    }else{
        if previousNode == previousNode.parent.left {
            previousNode.parent.left=newNode
        }else{
            previousNode.parent.right=newNode
        }
    }
    if newNode != nil {
        newNode.parent=previousNode.parent
    }
}

func (tree *RedBlackTree) grandParent(node *RedBlackNode) *RedBlackNode {
    return node.parent
}

func (tree *RedBlackTree) sibling(node *RedBlackNode) *RedBlackNode {
    if node == node.parent.left {
        return node.parent.right
    }else{
        return node.parent.left
    }
}

func (tree *RedBlackTree) uncle(node *RedBlackNode) *RedBlackNode {
    return tree.sibling(node)
}

func compareTo(a,b int) int {
    if a<b {
        return -1
    }else if a >b {
        return 1
    }else{
        return 0
    }
}


func main() {
    fmt.Println("RBT")
    rbt := new(RedBlackTree)

    var temp *RedBlackNode = new(RedBlackNode)
    temp.value=17
    rbt.Insert(temp)
    var other *RedBlackNode = new(RedBlackNode)
    other.value=12
    rbt.Insert(other)
    var x *RedBlackNode = new(RedBlackNode)
    x.value=15
    rbt.Insert(x)
    fmt.Println(rbt.Search(13))
    fmt.Println(rbt.Search(15))
}
