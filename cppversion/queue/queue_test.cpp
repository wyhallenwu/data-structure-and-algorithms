#include "queue.h"

int main(){
    Queue<int> *q = new Queue<int>;
    q->Push(10);
    q->Push(20);
    q->Push(30);
    std::cout<<*q;
    q->Pop();
    std::cout<<*q;
    std::cout<<q->Front()->_val<<std::endl;
    std::cout<<q->Tail()->_val<<std::endl;
    return 0;
}