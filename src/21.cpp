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
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        ListNode *p1 = l1;
        ListNode *p2 = l2;
        ListNode *h = NULL;
        while(p1 != NULL || p2 != NULL){
            ListNode *next;
            if (p1 == NULL || (p2 != NULL && p1->val > p2->val)) {
                next = p2;
                p2 = p2->next;
            } else {
                next = p1;
                p1 = p1->next;
            }
            if (h == NULL) {
                h = next;
            } else {
                h->next = next;
                h = next;
            }
        }
        return l1 == NULL || (l2 != NULL && l1->val > l2->val) ? l2 : l1;
    }
};

void print(ListNode *n) {
    cout << 666 << endl;
    while(n != NULL){
        cout << n->val << " ";
        n = n->next;
    }
    cout << endl;
}

int main(){
    ListNode *l11 = new ListNode(1);
    ListNode *l12 = l11->next = new ListNode(2);
    ListNode *l13 = l12->next = new ListNode(4);

    ListNode *l21 = new ListNode(1);
    ListNode *l22 = l21->next = new ListNode(3);
    ListNode *l23 = l22->next = new ListNode(4);

    Solution s;
    print(s.mergeTwoLists(l11, l21));
    print(s.mergeTwoLists(new ListNode(1), NULL));
    return 0;
}
