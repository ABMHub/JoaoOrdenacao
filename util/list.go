package util

type node struct{
    nxt *node 
    prv *node
    data T
}

type List struct{
    head *node 
    tail *node
    Size int
}

func NewNode(path T) *node{
    var el node
    el.data = path
    el.nxt = nil
    el.prv = nil    
    return &el
}

func NewList() List{
    return List{nil,nil,0}
}

func (v *List) Empty() bool{
    return v.head == nil
}

func (v *List) Push_back(path T){
    ptr := NewNode(path) 
    if v.head == nil{
        v.head = ptr
        v.tail = ptr
    }else{
        ptr.prv = v.tail
        v.tail.nxt = ptr
        v.tail = ptr
    }
    v.Size++
}

func (v *List) Pop_back() T{
    if v.head == nil{
        return ""
    }
    var schauss T = v.tail.data
    v.tail = v.tail.prv
    if(v.tail == nil){
        v.head = nil
    }else{
        v.tail.nxt = nil
    }
    v.Size--
    return schauss
}

func (v *List) Front() T{
    if(v.head==nil){
        return nil
    }
    return v.head.data
}

func (v *List) Back() T{
    if(v.tail==nil){
        return nil
    }
    return v.tail.data
}