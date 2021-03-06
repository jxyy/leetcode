#include<iostream>
#include<vector>
#include<queue>
#include<functional>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

struct cmp{
    bool operator()(ListNode* a, ListNode* b){
        return a->val > b->val;
    }
};

class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        ListNode* ret_head = NULL;
        ListNode* ret_tail = NULL;
        int k = lists.size();
        auto cmp = [](ListNode* a, ListNode* b){return a->val > b->val;};
        priority_queue<ListNode*, vector<ListNode*>, decltype(cmp) > q(cmp);
        for(int i = 0; i < k; i++){
            if(lists[i] != NULL)
                q.push(lists[i]);
        }

        while(q.size() > 0){
            ListNode* min = q.top();
            if(ret_head == NULL){
                ret_head = min;
                ret_tail = min;
            }else{
                ret_tail->next = min;
                ret_tail = min;
            }
            q.pop();
            if(min->next != NULL){
                q.push(min->next);
            }
        }    
        return ret_head;
    }
};

int main(){
    ListNode* n1 = new ListNode(1);
    ListNode* n2 = new ListNode(5);
    ListNode* n3 = new ListNode(10);
    n1->next = n2;
    n2->next = n3;

    ListNode* n4 = new ListNode(3);
    ListNode* n5 = new ListNode(7);
    ListNode* n6 = new ListNode(9);
    n4->next = n5;
    n5->next = n6;

    vector<ListNode*> input;
    input.push_back(n1);
    input.push_back(n4);
    Solution s;
    ListNode* ret = s.mergeKLists(input);
    while(ret != NULL){
        cout << ret->val << endl;
        ret = ret->next;
    }
    return 0;
}
