#include "stack.h"

int main(){
    Stack<int> *s = new Stack<int>();
    s->Push(10);
    s->Push(20);
    cout<<*s;
    s->Pop();
    cout<<*s;
    auto t = s->Top();
    cout<<*t;
    return 0;
}