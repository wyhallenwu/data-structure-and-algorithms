#include "iostream"
#include "vector"

template<typename T>
struct node{
    T _val;
    node *_next;

    node(T val){
        _val = val;
        _next = nullptr;
    }

    node(){
        _next = nullptr;
    }
};


template<typename T>
class Stack{
private:
    node<T> *_top;
    int length;
public:
    Stack(){
        _top = new(node<T>);
        length = 0;
    }
    virtual ~Stack(){
        Clear(_top);
    }

    void Clear(node<T> *top){
        while(top!= nullptr){
            node<T> *del = new node<T>();
            del = top;
            top = top->_next;
            delete(del);
        }
    }

    // Push
    void Push(T val){
        node<T> *in = new node<T>(val);
        in->_next = _top->_next;
        _top->_next = in;
    }

    // Pop
    void Pop(){
        node<T> del = new node<T>();
        del = _top->_next;
        _top->_next = del._next;
        delete del;
    }

    // overload >>


};