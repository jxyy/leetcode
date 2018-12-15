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
        if(head == NULL || head->next == NULL){
            return true;
        }else if(head->next->next == NULL){
            return head->val == head->next->val;
        }

        ListNode *slow = head, *quick = head->next;
        ListNode *p1, *p2;
        ListNode *backup = NULL;
        while(1) {
            if(quick->next == NULL) {
                p1 = slow->next;
                slow->next = backup;
                p2 = slow;
                break;
            }else if(quick->next->next == NULL){
                p1 = slow->next->next;
                slow->next = backup;
                p2 = slow;
                break;
            }
            quick = quick->next->next;
            ListNode* tmp = slow->next;
            slow->next = backup;
            backup = slow;
            slow = tmp;
        }

        while(p1 != NULL){
            if(p1->val != p2->val){
                return false;
            }
            p1 = p1->next;
            p2 = p2->next;
        }
        return true;
    }

    bool isPalindrome0(ListNode* head) {
        int l = 0;
        ListNode* p = head;
        while(p != NULL){
            l++;
            p = p->next;
        }
        if(l < 2){
            return true;
        }

        int half = l / 2;
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
            next = next->next;
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
    ListNode *n2 = n1->next = new ListNode(2);
    ListNode *n3 = n2->next = new ListNode(1);
    cout << s.isPalindrome(head) << endl;

    head = new ListNode(1);
    n1 = head->next = new ListNode(0);
    n2 = n1->next = new ListNode(0);

    cout << s.isPalindrome(head) << endl;
    return 0;
}
