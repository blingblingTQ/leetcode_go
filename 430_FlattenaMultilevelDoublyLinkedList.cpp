class Node {
public:
    int val;
    Node* prev;
    Node* next;
    Node* child;

    Node() {}

    Node(int _val, Node* _prev, Node* _next, Node* _child) {
        val = _val;
        prev = _prev;
        next = _next;
        child = _child;
    }
};

class Solution {
public:
    Node* flatten(Node* head) {
        Node *curr = head;
        while(curr != NULL) {
            if(curr->child != NULL) {
                Node *next = curr->next;
                Node *child = flatten(curr->child);
                Node *childEnd = child;
                while(childEnd->next != NULL) {
                    childEnd = childEnd->next;
                }
                curr->next = child;
                child->prev = curr;
                curr->child = NULL;
                childEnd->next = next;
                if(next != NULL) {
                    next->prev = childEnd;
                }
            }
            curr = curr->next;
        }
        return head;
    }
};
