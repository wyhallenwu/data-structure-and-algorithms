#include "iostream"
#include "vector"

using namespace std;
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

    friend ostream &operator<<(ostream &out, const node<T> &n){
        out << n._val <<endl;
        return out;
    }
};


template<typename T>
class Stack{
private:
    node<T> *_top;
    int length;
public:
    Stack(){
        _top = new node<T>();
        length = 0;
    }
    virtual ~Stack(){
        Clear(_top);
    }

    void Clear(node<T> *top){
        while(top!= nullptr){
            auto *del = new node<T>();
            del = top;
            top = top->_next;
            delete del;
        }
    }

    node<T>* Top(){
        return _top->_next;
    }

    // Push
    void Push(T val){
         auto *in = new node<T>(val);
        in->_next = _top->_next;
        _top->_next = in;
        length++;
    }

    // Pop
    void Pop(){
        node<T> *del = new node<T>();
        del = _top->_next;
        _top->_next = del->_next;
        delete del;
        length--;
    }

    // overload <<
    friend std::ostream &operator<<(std::ostream &output, const Stack<T> &s){
        auto *n = new node<T>;
        n = s._top->_next;
        output<<"length is: "<< s.length <<endl;
        while(n!= nullptr){
            output << n->_val << "->";
            n = n->_next;
        }
        output<<"||END"<<std::endl;
        return output;
    }

};