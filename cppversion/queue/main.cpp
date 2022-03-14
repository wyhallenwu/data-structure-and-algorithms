#include "queue.h"

int main(){
    Queue<int> *s = new Queue<int>;
    s->Push(10);
    s->Push(20);
    s->Push(30);
    std::cout<<*s;
    s->Pop();
    std::cout<<*s;
    std::cout<<s->Front()->_val<<std::endl;
    std::cout<<s->Tail()->_val<<std::endl;
    return 0;
}