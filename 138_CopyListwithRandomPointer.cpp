#include <iostream>
using namespace std;

 class Node {
 public:
     int val;
     Node* next;
     Node* random;
     
     Node() {}
     
     Node(int _val, Node* _next, Node* _random) {
         val = _val;
         next = _next;
         random = _random;
     }
 };

class Solution {
public:
    
    void doubleList(Node *head) {
        Node *curr = head;
        while(curr != NULL) {
            Node *copy = new Node(curr->val, NULL, NULL);
            Node *next = curr->next;
            copy->next = next;
            curr->next = copy;
            curr = next;
        }
    }
    
    void copyRandomPointer(Node *head) {
        Node *first = head;
        while(first != NULL) {
            if(first->random != NULL) {
                first->next->random = first->random->next;
            } else {
                first->next->random = NULL;
            }
            first = first->next->next;
        }
    }
    
    Node *separateList(Node *head) {
        Node *first = head;
        Node *second = head->next;
        Node *copyHead = second;
        while(second != NULL && second->next != NULL) {
            first->next = first->next->next;
            second->next = second->next->next;
            first = first->next;
            second = second->next;
        }
        first->next = NULL;
        return copyHead;
    }
    
    Node* copyRandomList(Node* head) {
        if (head == NULL) {
            return NULL;
        }
        doubleList(head);
        copyRandomPointer(head);
        return separateList(head);
    }
};

void printList(Node *node) {
    for (auto curr = node; curr != NULL; curr = curr->next) {
        std::cout << curr->val << "->";
    }
    std::cout << std::endl;
}

int main(void) {
    Solution solution;
    Node node1;
    Node node2;
    node1.val = 1;
    node2.val = 2;
    node1.next = &node2;
    node2.next = NULL;
    node1.random = &node2;
    node2.random = &node2;
    Node *result = solution.copyRandomList(&node1);
    printList(result);
    return 0;
}

