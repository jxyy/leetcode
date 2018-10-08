#include<iostream>
#include<vector>
using namespace std;

 struct ListNode {
     int val;
     ListNode *next;
     ListNode(int x) : val(x), next(NULL) {}
 };

class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        if(head == NULL){
            return NULL;
        }
        ListNode *slow, *fast;
        slow = fast = head;
        do{
            slow = slow->next;
            fast = fast->next;
            if(fast != NULL){
                fast = fast->next;
            }
        }while(fast != NULL && slow != fast);
        if(fast == NULL) {
            return NULL;
        }
        slow = head;
        while(slow != fast) {
            slow = slow->next;
            fast = fast->next;
        }
        return slow;
    }
};

int main(){
    Solution s;
    ListNode *n = new ListNode(0);
    ListNode *head = n;
    for(int i = 0; i < 5; i ++) {
        n->next = new ListNode(i+1);
        n = n->next;
    }
    n->next = head->next->next;
    cout << s.detectCycle(head)->val << endl;

    cout << s.detectCycle(NULL) << endl;
    return 0;
}
