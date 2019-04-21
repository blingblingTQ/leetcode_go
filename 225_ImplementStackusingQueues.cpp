#include <queue>
using namespace std;

class MyStack {
public:
    /** Initialize your data structure here. */
    MyStack() {
        
    }
    
    /** Push element x onto stack. */
    void push(int x) {
        if(queue1.empty()) {
            queue2.push(x);
        } else {
            queue1.push(x);
        }
    }
    
    /** Removes the element on top of the stack and returns that element. */
    int pop() {
        if(queue1.empty()) {
            int size = queue2.size();
            for(int i = 1; i < size; i++) {
                queue1.push(queue2.front());
                queue2.pop();
            }
            int val = queue2.front();
            queue2.pop();
            return val;
        } else {
             int size = queue1.size();
            for(int i = 1; i < size; i++) {
                queue2.push(queue1.front());
                queue1.pop();
            }
            int val = queue1.front();
            queue1.pop();
            return val;
        }
    }
    
    /** Get the top element. */
    int top() {
        if(queue1.empty()) {
            int size = queue2.size();
            for(int i = 1; i < size; i++) {
                queue1.push(queue2.front());
                queue2.pop();
            }
            int val = queue2.front();
            queue2.pop();
            queue1.push(val);
            return val;
        } else {
             int size = queue1.size();
            for(int i = 1; i < size; i++) {
                queue2.push(queue1.front());
                queue1.pop();
            }
            int val = queue1.front();
            queue1.pop();
            queue2.push(val);
            return val;
        }
    }
    
    /** Returns whether the stack is empty. */
    bool empty() {
        return queue1.empty()&&queue2.empty();
    }
private:
    queue<int> queue1, queue2;
};

/**
 * Your MyStack object will be instantiated and called as such:
 * MyStack* obj = new MyStack();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->top();
 * bool param_4 = obj->empty();
 */
