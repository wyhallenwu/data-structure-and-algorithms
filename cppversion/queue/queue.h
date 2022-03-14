#include "iostream"

template<typename T>
struct node{
    T _val;
    node *_next;
    node *_pre;

    node(){
        _next = _pre = nullptr;
    }

    node(T val){
        _val = val;
        _next = _pre = nullptr;
    }
};

template<typename T>
class Queue{
private:
    node<T> *head;
    node<T> *tail;
    int length;
public:
    Queue(){
        head = new node<T>;
        tail = new node<T>;
        head->_next = tail;
        head->_pre = nullptr;
        tail->_next = nullptr;
        tail->_pre = head;
        length = 0;
    }

    virtual ~Queue(){
        while(head->_next != tail){
            auto n = head->_next;
            head->_next = n->_next;
            delete n;
        }
        delete head;
        delete tail;
    }

    node<T>* Front(){
        if(head->_next != tail){
            return head->_next;
        } else{
            return nullptr;
        }
    }

    node<T>* Tail(){
        if(head->_next != tail){
            return tail->_pre;
        } else {
            return nullptr;
        }
    }

    void Push(T val) {
        auto n = new node<T>(val);
        n->_pre = tail->_pre->_pre;
        tail->_pre->_next = n;
        n->_next = tail;
        tail->_pre = n;
        length++;
    }

    void Pop(){
        auto n = new node<T>;
        if(length > 0) {
            n = head->_next;
            head->_next = n->_next;
            n->_next->_pre = head;
            delete n;
        } else {
            std::cout<<"Empty."<<std::endl;
        }
    }

    friend std::ostream &operator<<(std::ostream &out, const Queue<T> &q){
        auto n = q.head->_next;
        while(n != q.tail) {
            out<<n->_val<<"->";
            n = n->_next;
        }
        out<<"||END"<<std::endl;
        return out;
    }
};