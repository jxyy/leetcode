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
    bool isPalindrome(ListNode* head) {
        int l = 0;
        ListNode* p = head;
        while(p != NULL){
            l++;
            p = p->next;
        }
        if(l < 2){
            return true;
        }

        int half = (l - 1) / 2;
        int i = 0;
        ListNode* curr = NULL, *next = head;
        while(i < half) {
            ListNode* backup = next->next;
            next->next = curr;
            curr = next;
            next = backup;
            i ++;
        }
        if(l%2 == 1){
            curr = curr->next;
        }

        ListNode *p1 = curr, *p2 = next;
        while(p1 != NULL){
            if(p1->val != p2->val){
                return false;
            }
            p1 = p1->next;
            p2 = p2->next;
        }
        return true;
    }
};



int main(){
    Solution s;

    ListNode *head = new ListNode(1);
    ListNode *n1 = head->next = new ListNode(2);
    ListNode *n2 = n1->next = new ListNode(1);
    ListNode *n3 = n2->next = new ListNode(1);

    cout << s.isPalindrome(head) << endl;
    return 0;
}
