#include <iostream>
#include <queue>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

struct Item {
    ListNode *node;
    int index;
    Item(ListNode *n, int i): node(n), index(i) {}
    bool operator < (const Item & a) const
    {
        return node->val > a.node->val;
    }
};

class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        priority_queue<Item> queue;
        ListNode *dummy = new ListNode(0);
        ListNode *curr = dummy;
        
        vector<ListNode *> indexs(lists);
        for(int i = 0; i < indexs.size(); i++) {
            if(indexs[i] != NULL) {
                Item item(indexs[i], i);
                queue.push(item);
            }
        }
        
        while(!queue.empty()) {
            Item min = queue.top();
            queue.pop();
            
            curr->next = min.node;
            curr = curr->next;
            indexs[min.index] = indexs[min.index]->next;
            if (indexs[min.index] != NULL) {
                Item nextItem(indexs[min.index], min.index);
                queue.push(nextItem);
            }
        }
        ListNode *head = dummy->next;
        delete(dummy);
        return head;
    }
};

void printList(ListNode *head) {
    for(ListNode *curr = head; curr != NULL; curr = curr->next) {
        std::cout << curr->val << "->";
    }
    std::cout << std::endl;
}

int main(void) {
    Solution solution;
    ListNode n11(1), n12(4), n13(5);
    ListNode n21(1), n22(3), n23(4);
    ListNode n31(2), n32(6);
    n11.next = &n12;
    n12.next = &n13;
    n13.next = nullptr;
    
    n21.next = &n22;
    n22.next = &n23;
    n23.next = nullptr;
    
    n31.next = &n32;
    n32.next = nullptr;
    
    vector<ListNode *> lists = {&n11, &n21, &n31};
    printList(solution.mergeKLists(lists));
    return 0;
}

