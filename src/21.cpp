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
        ListNode h = NULL;
        while(p1 != NULL || p2 != NULL){
            if (p1 != NULL && p2 != NULL) {
            }
        }
    }
};

int main(){
    return 0;
}
