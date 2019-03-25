 struct ListNode {
     int val;
     ListNode *next;
     ListNode(int x) : val(x), next(NULL) {}
 };
 
class Solution {
public:
    ListNode* swapPairs(ListNode* head) {
        if (head == NULL || head->next == NULL) {
            return head;
        }
                
        ListNode *after = swapPairs(head->next->next);
        ListNode *curr = head;
        ListNode *next = head->next;
        head->next = after;
        next->next = head;
        return next;
    }


    ListNode* swapPairs(ListNode* head) {
        if (head == NULL || head->next == NULL) {
            return head;
        }
                
        ListNode *dummy = new ListNode(0);
        dummy->next = head;
        
        ListNode *curr = head;
        ListNode *prev = dummy;
        
        while(curr != NULL && curr->next != NULL) {
            ListNode *next = curr->next;
            prev->next = next;
            curr->next = next->next;
            next->next = curr;
            
            prev = curr;
            curr = curr->next;
        }
        ListNode *result = dummy->next;
        delete(dummy);
        return result;
    }

};
