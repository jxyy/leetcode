#include<iostream>
#include<vector>
#include<queue>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode* reverseKGroup(ListNode* head, int k) {
        if(k < 2 || head == NULL){
            return head;
        }

        bool done = false;
        bool is_first_round = true;
        ListNode* ret = head;
        ListNode* prev_p = NULL;
        ListNode* p = head;
        while(!done){
            ListNode* t = p;
            for(int i = 0; i < k - 1; i++){
                if(t == NULL || t->next == NULL){
                    done = true;
                    break;
                }
                t = t->next;
            }
            if(done){
                break;
            }
            if(is_first_round){
                ret = t;
                is_first_round = false;
            }else{
                prev_p->next = t;
            }

            ListNode* n = t->next;
            ListNode* next_p = n;
            prev_p = p;
            for(int i = 0; i < k; i++){
                ListNode* tmp = p;
                p = p->next;
                tmp->next = n;
                n = tmp;
            }

            p = next_p;
        }
        return ret;
    }
};

int main(){
    ListNode* head = new ListNode(0);
    ListNode* tail = head;
    for(int i = 0; i < 1; i++){
        ListNode* n = new ListNode(i+1);
        tail->next = n;
        tail = n;
    }
    cout << "before:" << endl;
    ListNode* p = head;
    while(p != NULL){
        cout << p->val << ' ';
        p = p->next;
    }
    cout << endl;
    
    Solution s;
    head = s.reverseKGroup(head, 2);

    cout << "after:" << endl;
    p = head;
    while(p != NULL){
        cout << p->val << ' ';
        p = p->next;
    }
    cout << endl;
    return 0;
}
